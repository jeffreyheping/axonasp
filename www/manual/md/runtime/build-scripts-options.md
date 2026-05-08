# Use AxonASP Build Scripts with Supported Options

## Overview
This page explains how to use the AxonASP build scripts and all supported options in each script.

The build scripts compile all AxonASP executables from one command:
- axonasp-http
- axonasp-fastcgi
- axonasp-cli
- axonasp-testsuite
- axonasp-mcp
- axonasp-service

Use:
- build.ps1 on Windows PowerShell
- build.sh on Linux and macOS Bash

## Syntax
Windows PowerShell:

```powershell
./build.ps1 [-Platform windows|linux|darwin|wasm|all] [-Architecture amd64|arm64|386] [-Clean] [-Test] [-Tags "tag1 tag2"]
```

Linux and macOS Bash:

```bash
./build.sh [--platform|-p windows|linux|darwin|wasm|all] [--arch|-a amd64|arm64|386] [--clean|-c] [--test|-t] [--tags|-g "tag1 tag2"]
```

## Parameters and Arguments
### Windows PowerShell (build.ps1)
- `-Platform`:
  - Type: String
  - Required: No
  - Default: `windows`
  - Allowed values: `windows`, `linux`, `darwin`, `all`
- `-Architecture`:
  - Type: String
  - Required: No
  - Default: `amd64`
  - Allowed values: `amd64`, `arm64`, `386`
- `-Clean`:
  - Type: Switch
  - Required: No
  - Purpose: Removes previous binaries and build directory before compiling.
- `-Test`:
  - Type: Switch
  - Required: No
  - Purpose: Runs `go test ./...` after build steps.
- `-Tags`:
  - Type: String
  - Required: No
  - Purpose: Passes Go build tags to all compilation targets.
  - Tag separators accepted: spaces, commas, semicolons.

### Linux and macOS Bash (build.sh)
- `--platform` or `-p`:
  - Type: String
  - Required: No
  - Default: `linux`
  - Allowed values: `windows`, `linux`, `darwin`, `all`
- `--arch` or `-a`:
  - Type: String
  - Required: No
  - Default: `amd64`
  - Allowed values: `amd64`, `arm64`, `386`
- `--clean` or `-c`:
  - Type: Flag
  - Required: No
  - Purpose: Removes previous binaries and build directory before compiling.
- `--test` or `-t`:
  - Type: Flag
  - Required: No
  - Purpose: Runs `go test ./...` after build steps.
- `--tags` or `-g`:
  - Type: String
  - Required: No
  - Purpose: Passes Go build tags to all compilation targets.
  - Tag separators accepted: spaces, commas, semicolons.

## Return Values
- Exit code `0`: All targets completed successfully.
- Exit code `1`: At least one build or test step failed.

## Remarks
- Both scripts format source files and run `go generate ./...` before building.
- For cross-platform output, use `all` in platform selection.
- On Windows target builds, binaries use `.exe`. On Linux and macOS, binaries are created without extension.
- Tag-based library disabling should be handled through the tags option in each script.

## Code Example
Windows examples:

```powershell
# Default Windows build
./build.ps1

# Build all platforms for amd64 and run tests
./build.ps1 -Platform all -Architecture amd64 -Test

# Clean and compile with disable tags
./build.ps1 -Clean -Tags "lib_adodb_disabled lib_msxml_disabled"
```

Linux and macOS examples:

```bash
# Default Linux build
./build.sh

# Build all platforms with tests
./build.sh --platform all --arch amd64 --test

# Clean and compile with disable tags
./build.sh --clean --tags "lib_adodb_disabled,lib_msxml_disabled"
``````