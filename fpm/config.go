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

// Package main implements the AxonASP FPM (FastCGI Process Manager) supervisor.
// This file contains platform-agnostic configuration types and helper functions
// that are safe to compile and test on all platforms, including Windows.
package main

import (
	"fmt"
	"strings"
)

// PoolConfig represents the per-pool configuration parsed from a .conf TOML file.
type PoolConfig struct {
	SiteName      string `toml:"site_name"`
	UID           uint32 `toml:"uid"`
	GID           uint32 `toml:"gid"`
	Socket        string `toml:"socket"`
	ConfigFile    string `toml:"config_file"`
	GlobalAsa     string `toml:"global_asa_path"`
	AppPath       string `toml:"app_path"`
	MemoryLimitMB int    `toml:"memory_limit_mb"`
	MaxRestarts   int    `toml:"max_restarts"`
	TmpDir        string `toml:"tmp_dir"`
}

// normalizePoolSocketEndpoint normalizes pool socket configuration and returns
// the FastCGI listen endpoint plus a filesystem path when using unix sockets.
func normalizePoolSocketEndpoint(raw string) (listenEndpoint string, socketPath string, isUnix bool, err error) {
	value := strings.TrimSpace(raw)
	if value == "" {
		return "", "", false, fmt.Errorf("socket is required in pool config")
	}

	lower := strings.ToLower(value)
	if strings.HasPrefix(lower, "unix:") {
		path := strings.TrimSpace(value[len("unix:"):])
		if path == "" {
			return "", "", false, fmt.Errorf("unix socket path cannot be empty")
		}
		return "unix:" + path, path, true, nil
	}

	if strings.HasPrefix(value, "/") || strings.HasPrefix(value, "./") || strings.HasPrefix(value, "../") {
		return "unix:" + value, value, true, nil
	}

	return value, "", false, nil
}

// buildWorkerArgs returns FastCGI worker startup args from pool configuration.
// It explicitly passes --server.web_root so the FastCGI worker uses the correct
// web root directory instead of falling back to the default ./www relative path.
func buildWorkerArgs(conf PoolConfig, listenEndpoint string) []string {
	args := []string{
		"--fastcgi.server_port", listenEndpoint,
		"--config.config_file", conf.ConfigFile,
		"--global.temp_dir", conf.TmpDir,
	}
	if globalASADir := strings.TrimSpace(conf.GlobalAsa); globalASADir != "" {
		args = append(args, "--config.global_asa", globalASADir)
	}
	if poolName := strings.TrimSpace(conf.SiteName); poolName != "" {
		args = append(args, "--pool.name", poolName)
	}
	// Explicitly pass the web root so the FastCGI worker uses the correct
	// directory. Without this flag, the worker falls back to ./www relative
	// to its CWD, which causes "stat ./www: no such file or directory" when
	// the pool's app_path is already the web root directory.
	if appPath := strings.TrimSpace(conf.AppPath); appPath != "" {
		args = append(args, "--server.web_root", appPath)
	}
	return args
}
