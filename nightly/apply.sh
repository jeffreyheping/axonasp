#!/usr/bin/env bash
#
# Nightly build: replay our fork-specific patches on top of fresh upstream code.
#
# Run this from the repository root, AFTER the working tree has been switched
# to upstream/main (so the tree is pure upstream + this nightly/ directory left
# behind as untracked). It:
#   1. Copies overlay files (new files we add + files we fully own) over the tree.
#   2. Patches go.mod idempotently (replace directive + ulikunitz/xz dep).
#   3. Patches Dockerfile idempotently (COPY third_party before go mod download).
#
# The script is idempotent: re-running it on an already-patched tree is a no-op.
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
OVERLAY="$ROOT/nightly/overlay"

echo "==> 1/4 Copying overlay files onto the upstream tree"
cp -a "$OVERLAY/." "$ROOT/"
echo "    overlay applied"

echo "==> 2/4 Patching go.mod (replace directive)"
if grep -q 'replace github.com/ricochet2200/go-disk-usage/du => ./third_party/go-disk-usage' go.mod; then
    echo "    replace directive already present"
else
    printf '\nreplace github.com/ricochet2200/go-disk-usage/du => ./third_party/go-disk-usage\n' >> go.mod
    echo "    replace directive added"
fi

echo "==> 3/4 Patching go.mod (ulikunitz/xz dependency) + go mod tidy"
if grep -q 'github.com/ulikunitz/xz' go.mod; then
    echo "    ulikunitz/xz already in go.mod"
else
    go get github.com/ulikunitz/xz@v0.5.12
    echo "    ulikunitz/xz added"
fi
go mod tidy

echo "==> 4/4 Patching Dockerfile (COPY third_party before go mod download)"
if grep -q 'COPY third_party/go-disk-usage' Dockerfile; then
    echo "    COPY third_party already present"
else
    # Insert the COPY line right after "COPY go.mod go.sum ./"
    if grep -q '^COPY go.mod go.sum \./' Dockerfile; then
        sed -i '/^COPY go.mod go.sum \.\//a COPY third_party/go-disk-usage ./third_party/go-disk-usage/' Dockerfile
        echo "    COPY third_party inserted after go.mod line"
    else
        echo "ERROR: could not find 'COPY go.mod go.sum ./' in Dockerfile; upstream layout changed, manual fix needed" >&2
        exit 1
    fi
fi

echo "==> Patch replay complete"
