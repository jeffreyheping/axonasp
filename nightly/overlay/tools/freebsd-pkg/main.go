// tools/freebsd-pkg — construct a FreeBSD .pkg archive on any OS.
// A .pkg is an xz-compressed tar archive containing:
//   +MANIFEST        JSON metadata
//   +COMPACT_MANIFEST  compact metadata (optional, generated from +MANIFEST)
//   +POST_INSTALL    shell script
//   +DEINSTALL       shell script
//   files…          installed files with absolute paths
package main

import (
	"archive/tar"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/ulikunitz/xz"
)

type manifest struct {
	Name         string            `json:"name"`
	Version      string            `json:"version"`
	Origin       string            `json:"origin"`
	Comment      string            `json:"comment"`
	Maintainer   string            `json:"maintainer"`
	WWW          string            `json:"www"`
	Prefix       string            `json:"prefix"`
	Arch         string            `json:"arch"`
	Licenses     []string          `json:"licenses"`
	LicenseLogic string            `json:"licenselogic"`
	Files        map[string]string `json:"files"`
	Scripts      map[string]string `json:"scripts,omitempty"`
}

func main() {
	if len(os.Args) < 7 {
		fmt.Fprintf(os.Stderr, "Usage: %s <out.pkg> <name> <version> <arch> <prefix> <stagedir> [post-install-script] [deinstall-script]\n", os.Args[0])
		os.Exit(1)
	}
	outPath := os.Args[1]
	name := os.Args[2]
	version := os.Args[3]
	arch := os.Args[4]
	prefix := os.Args[5]
	stageDir := os.Args[6]

	postInstall := ""
	if len(os.Args) > 7 {
		postInstall = os.Args[7]
	}
	deinstall := ""
	if len(os.Args) > 8 {
		deinstall = os.Args[8]
	}

	m := manifest{
		Name:         name,
		Version:      version,
		Origin:       "www/" + name,
		Comment:      name + " Server",
		Maintainer:   "G3pix <axonasp@g3pix.com.br>",
		WWW:          "https://g3pix.com.br/axonasp",
		Prefix:       prefix,
		Arch:         arch,
		Licenses:     []string{"MPL20"},
		LicenseLogic: "single",
		Files:        make(map[string]string),
		Scripts:      make(map[string]string),
	}

	// Collect files and hashes from stage dir.
	var entries []tarEntry
	err := filepath.Walk(stageDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		rel, err := filepath.Rel(stageDir, path)
		if err != nil {
			return err
		}
		installPath := filepath.Join(prefix, rel)
		installPath = filepath.ToSlash(installPath)

		hash, err := fileSha256(path)
		if err != nil {
			return err
		}
		m.Files[installPath] = hash

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		entries = append(entries, tarEntry{path: installPath, data: data, mode: int64(info.Mode().Perm())})
		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "walk: %v\n", err)
		os.Exit(1)
	}

	if postInstall != "" {
		m.Scripts["post-install"] = postInstall
	}
	if deinstall != "" {
		m.Scripts["deinstall"] = deinstall
	}

	manifestJSON, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "marshal: %v\n", err)
		os.Exit(1)
	}

	// Build tar stream
	pr, pw := io.Pipe()
	go func() {
		tw := tar.NewWriter(pw)
		// +MANIFEST
		writeTarFile(tw, "+MANIFEST", manifestJSON, 0644)
		// +COMPACT_MANIFEST (gzip json, same content)
		var buf strings.Builder
		gzw := gzip.NewWriter(&buf)
		gzw.Write(manifestJSON)
		gzw.Close()
		writeTarFile(tw, "+COMPACT_MANIFEST", []byte(buf.String()), 0644)
		// scripts
		if postInstall != "" {
			writeTarFile(tw, "+POST_INSTALL", []byte(postInstall), 0755)
		}
		if deinstall != "" {
			writeTarFile(tw, "+DEINSTALL", []byte(deinstall), 0755)
		}
		// data files
		for _, e := range entries {
			writeTarFile(tw, e.path, e.data, e.mode)
		}
		tw.Close()
		pw.Close()
	}()

	out, err := os.Create(outPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "create: %v\n", err)
		os.Exit(1)
	}
	defer out.Close()

	xzw, err := xz.NewWriter(out)
	if err != nil {
		fmt.Fprintf(os.Stderr, "xz: %v\n", err)
		os.Exit(1)
	}
	if _, err := io.Copy(xzw, pr); err != nil {
		fmt.Fprintf(os.Stderr, "copy: %v\n", err)
		os.Exit(1)
	}
	xzw.Close()
	fmt.Printf("Created %s\n", outPath)
}

type tarEntry struct {
	path string
	data []byte
	mode int64
}

func writeTarFile(tw *tar.Writer, name string, data []byte, mode int64) {
	hdr := &tar.Header{
		Name: name,
		Mode: mode,
		Size: int64(len(data)),
	}
	tw.WriteHeader(hdr)
	tw.Write(data)
}

func fileSha256(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return "sha256:" + hex.EncodeToString(h.Sum(nil)), nil
}
