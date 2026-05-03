//go:build !lib_g3search_disabled

/*
 * AxonASP Server
 * Copyright (C) 2026 G3pix Ltda. All rights reserved.
 *
 * Developed by Lucas Guimarães - G3pix Ltda
 * Contact: https://g3pix.com.br
 * Project URL: https://g3pix.com.br/axonasp
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 *
 * Attribution Notice:
 * If this software is used in other projects, the name "AxonASP Server"
 * must be cited in the documentation or "About" section.
 *
 * Contribution Policy:
 * Modifications to the core source code of AxonASP Server must be
 * made available under this same license terms.
 */

package axonvm

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"g3pix.com.br/axonasp/vbscript"
	"github.com/blugelabs/bluge"
)

var (
	globalBlugeReader     *bluge.Reader
	globalBlugeReaderPath string
	globalReaderMutex     sync.RWMutex
	globalReaderInitOnce  sync.Once
	globalReaderInitErr   error
)

// G3Search provides high-performance document indexing and searching using Bluge.
type G3Search struct {
	indexPath string
	docsPath  string
	extension string
	vm        *VM
}

// NewG3Search creates a new G3Search object with default extension.
func NewG3Search(vm *VM) *G3Search {

	return &G3Search{
		extension: ".md",
		vm:        vm,
	}
}

// closeGlobalReaderLocked closes the shared reader and clears global reader state.
func closeGlobalReaderLocked() error {
	if globalBlugeReader == nil {
		return nil
	}
	err := globalBlugeReader.Close()
	globalBlugeReader = nil
	globalBlugeReaderPath = ""
	return err
}

// openGlobalReaderLocked opens the shared reader for one index path under write lock.
func openGlobalReaderLocked(indexPath string) error {
	config := bluge.DefaultConfig(indexPath)
	reader, err := bluge.OpenReader(config)
	if err != nil {
		return err
	}
	globalBlugeReader = reader
	globalBlugeReaderPath = indexPath
	return nil
}

// tryInitGlobalReaderOnce performs the first lazy reader initialization once.
func (s *G3Search) tryInitGlobalReaderOnce() error {
	globalReaderInitOnce.Do(func() {
		globalReaderMutex.Lock()
		defer globalReaderMutex.Unlock()

		if globalBlugeReader != nil {
			if strings.EqualFold(globalBlugeReaderPath, s.indexPath) {
				globalReaderInitErr = nil
				return
			}
			_ = closeGlobalReaderLocked()
		}

		globalReaderInitErr = openGlobalReaderLocked(s.indexPath)
	})
	return globalReaderInitErr
}

// acquireGlobalReaderForSearch returns a shared reader pinned by RLock for one search call.
func (s *G3Search) acquireGlobalReaderForSearch() (*bluge.Reader, func(), error) {
	_ = s.tryInitGlobalReaderOnce()

	for {
		globalReaderMutex.RLock()
		if globalBlugeReader != nil && strings.EqualFold(globalBlugeReaderPath, s.indexPath) {
			return globalBlugeReader, globalReaderMutex.RUnlock, nil
		}
		globalReaderMutex.RUnlock()

		globalReaderMutex.Lock()
		if globalBlugeReader != nil && !strings.EqualFold(globalBlugeReaderPath, s.indexPath) {
			if err := closeGlobalReaderLocked(); err != nil {
				globalReaderMutex.Unlock()
				return nil, nil, err
			}
		}
		if globalBlugeReader == nil {
			if err := openGlobalReaderLocked(s.indexPath); err != nil {
				globalReaderMutex.Unlock()
				return nil, nil, err
			}
		}
		globalReaderMutex.Unlock()
	}
}

// DispatchMethod executes methods and property Let/Set behavior for G3Search.
func (s *G3Search) DispatchMethod(methodName string, args []Value) Value {

	switch {
	case strings.EqualFold(methodName, "BuildIndex"):
		return s.buildIndex()
	case strings.EqualFold(methodName, "Search"):
		if len(args) < 1 {
			return Value{Type: VTEmpty}
		}
		return s.search(args[0].String())
	case strings.EqualFold(methodName, "IndexPath"):
		if len(args) > 0 {
			s.indexPath = args[0].String()
			return Value{Type: VTEmpty}
		}
		return NewString(s.indexPath)
	case strings.EqualFold(methodName, "DocsPath"):
		if len(args) > 0 {
			s.docsPath = args[0].String()
			return Value{Type: VTEmpty}
		}
		return NewString(s.docsPath)
	case strings.EqualFold(methodName, "Extension"):
		if len(args) > 0 {
			s.extension = args[0].String()
			if s.extension != "" && !strings.HasPrefix(s.extension, ".") {
				s.extension = "." + s.extension
			}
			return Value{Type: VTEmpty}
		}
		return NewString(s.extension)
	default:
		return Value{Type: VTEmpty}
	}
}

// DispatchPropertyGet resolves property reads for G3Search.
func (s *G3Search) DispatchPropertyGet(propertyName string) Value {

	switch {
	case strings.EqualFold(propertyName, "IndexPath"):
		return NewString(s.indexPath)
	case strings.EqualFold(propertyName, "DocsPath"):
		return NewString(s.docsPath)
	case strings.EqualFold(propertyName, "Extension"):
		return NewString(s.extension)
	default:
		return Value{Type: VTEmpty}
	}
}

// DispatchPropertySet handles Let assignments on G3Search properties.
func (s *G3Search) DispatchPropertySet(propertyName string, val Value) {

	switch {
	case strings.EqualFold(propertyName, "IndexPath"):
		s.indexPath = val.String()
	case strings.EqualFold(propertyName, "DocsPath"):
		s.docsPath = val.String()
	case strings.EqualFold(propertyName, "Extension"):
		s.extension = val.String()
		if s.extension != "" && !strings.HasPrefix(s.extension, ".") {
			s.extension = "." + s.extension
		}
	}
}

// buildIndex iterates through DocsPath and indexes files matching Extension.
func (s *G3Search) buildIndex() Value {

	if s.docsPath == "" {
		s.vm.raise(vbscript.InternalError, ErrG3SearchDocsPathMissing.String())
		return Value{Type: VTEmpty}
	}
	if s.indexPath == "" {
		s.vm.raise(vbscript.InternalError, ErrG3SearchIndexPathMissing.String())
		return Value{Type: VTEmpty}
	}

	globalReaderMutex.Lock()
	defer globalReaderMutex.Unlock()

	if err := closeGlobalReaderLocked(); err != nil {
		s.vm.raise(vbscript.InternalError, ErrG3SearchIndexOpenFailed.String()+": "+err.Error())
		return Value{Type: VTEmpty}
	}

	config := bluge.DefaultConfig(s.indexPath)
	writer, err := bluge.OpenWriter(config)
	if err != nil {
		s.vm.raise(vbscript.InternalError, ErrG3SearchIndexOpenFailed.String()+": "+err.Error())
		return Value{Type: VTEmpty}
	}

	err = filepath.Walk(s.docsPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && (s.extension == "" || strings.EqualFold(filepath.Ext(path), s.extension)) {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			doc := bluge.NewDocument(path)
			doc.AddField(bluge.NewTextField("content", string(content)))
			doc.AddField(bluge.NewStoredOnlyField("filename", []byte(path)))

			err = writer.Update(doc.ID(), doc)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		_ = writer.Close()
		s.vm.raise(vbscript.InternalError, ErrG3SearchIndexWriteFailed.String()+": "+err.Error())
		return Value{Type: VTEmpty}
	}

	if err := writer.Close(); err != nil {
		s.vm.raise(vbscript.InternalError, ErrG3SearchIndexWriteFailed.String()+": "+err.Error())
		return Value{Type: VTEmpty}
	}

	if err := openGlobalReaderLocked(s.indexPath); err != nil {
		s.vm.raise(vbscript.InternalError, ErrG3SearchIndexOpenFailed.String()+": "+err.Error())
		return Value{Type: VTEmpty}
	}

	return Value{Type: VTEmpty}
}

// search performs a search on the index and returns [filename, score] tuples.
func (s *G3Search) search(term string) Value {

	if s.indexPath == "" {
		s.vm.raise(vbscript.InternalError, ErrG3SearchIndexPathMissing.String())
		return Value{Type: VTEmpty}
	}

	reader, release, err := s.acquireGlobalReaderForSearch()
	if err != nil {
		// If index doesn't exist yet, return empty array.
		return Value{Type: VTArray, Arr: NewVBArrayFromValues(0, nil)}
	}
	defer release()

	query := bluge.NewMatchQuery(term).SetField("content")
	searchRequest := bluge.NewAllMatches(query)

	iter, err := reader.Search(context.Background(), searchRequest)
	if err != nil {
		s.vm.raise(vbscript.InternalError, ErrG3SearchSearchFailed.String()+": "+err.Error())
		return Value{Type: VTEmpty}
	}

	results := make([]Value, 0, 16)
	match, iterErr := iter.Next()
	for iterErr == nil && match != nil {
		filename := ""
		match.VisitStoredFields(func(field string, value []byte) bool {
			if field == "filename" {
				filename = string(value)
			}
			return true
		})

		if filename != "" {
			tuple := [2]Value{NewString(filename), NewDouble(match.Score)}
			results = append(results, Value{Type: VTArray, Arr: NewVBArrayFromValues(0, tuple[:])})
		}

		match, iterErr = iter.Next()
	}

	if iterErr != nil {
		s.vm.raise(vbscript.InternalError, ErrG3SearchSearchFailed.String()+": "+iterErr.Error())
		return Value{Type: VTEmpty}
	}

	return Value{Type: VTArray, Arr: NewVBArrayFromValues(0, results)}
}
