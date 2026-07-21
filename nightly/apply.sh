#!/usr/bin/env bash
#
# Nightly build: replay our fork-specific patches on top of fresh upstream code.
#
# Run this from the repository root, AFTER the working tree has been switched
# to upstream/main (so the tree is pure upstream + this nightly/ directory left
# behind as untracked). It:
#   1. Copies overlay files (new files we add + files we fully own) over the tree.
#   2. Patches go.mod idempotently (add ulikunitz/xz dep for freebsd-pkg tool).
#   3. Runs go mod tidy.
#
# The script is idempotent: re-running it on an already-patched tree is a no-op.
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
OVERLAY="$ROOT/nightly/overlay"

echo "==> 1/3 Copying overlay files onto the upstream tree"
cp -a "$OVERLAY/." "$ROOT/"
echo "    overlay applied"

echo "==> 2/3 Patching go.mod (ulikunitz/xz dependency for freebsd-pkg tool)"
if grep -q 'github.com/ulikunitz/xz' go.mod; then
    echo "    ulikunitz/xz already in go.mod"
else
    go get github.com/ulikunitz/xz@v0.5.12
    echo "    ulikunitz/xz added"
fi

echo "==> 3/3 go mod tidy"
go mod tidy

echo "==> Patch replay complete"
