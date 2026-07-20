#!/usr/bin/env bash
set -euo pipefail

# Delete all tags matching a given k8s version from a fork (dry-run by default).
#   ./delete-fork-tags.sh v1.33.13          # dry run, lists matches
#   ./delete-fork-tags.sh v1.33.13 --yes    # actually delete
#
# Refuses any remote other than rafaelbreno/k3s-kubernetes unless --allow-any-remote.

REMOTE_URL="git@github.com:rafaelbreno/k3s-kubernetes.git"
EXPECTED_MATCH="rafaelbreno/k3s-kubernetes"
CONFIRM=false
ALLOW_ANY_REMOTE=false
VERSION=""

for arg in "$@"; do
  case "$arg" in
    --yes|-y)           CONFIRM=true ;;
    --allow-any-remote) ALLOW_ANY_REMOTE=true ;;
    --remote=*)         REMOTE_URL="${arg#*=}" ;;
    -*)                 echo "unknown flag: $arg" >&2; exit 2 ;;
    *)                  VERSION="$arg" ;;
  esac
done

[ -n "$VERSION" ] || { echo "usage: $0 <version> [--yes] [--remote=<url>] [--allow-any-remote]" >&2; exit 2; }

# normalize: ensure leading v
case "$VERSION" in v*) ;; *) VERSION="v${VERSION}" ;; esac

# guard remote
if [ "$ALLOW_ANY_REMOTE" != true ] && [[ "$REMOTE_URL" != *"$EXPECTED_MATCH"* ]]; then
  echo "refusing to operate on '$REMOTE_URL' (expected '$EXPECTED_MATCH'); pass --allow-any-remote to override" >&2
  exit 1
fi

# (start or /) <version> optionally -<suffix>, anchored to end.
# matches: v1.33.13, v1.33.13-k3s1, staging/src/k8s.io/api/v1.33.13-k3s1
# does NOT match: v1.33.130
esc_version="$(printf '%s' "$VERSION" | sed 's/\./\\./g')"
PATTERN="(^|/)${esc_version}(-[^/]*)?$"

echo "==> remote:  $REMOTE_URL"
echo "==> version: $VERSION"
echo "==> pattern: $PATTERN"
echo "==> fetching remote tags..."

mapfile -t TAGS < <(
  git ls-remote --tags "$REMOTE_URL" \
    | awk '{print $2}' \
    | grep -v '\^{}$' \
    | sed 's|^refs/tags/||' \
    | grep -E "$PATTERN" \
    | sort -V
)

if [ "${#TAGS[@]}" -eq 0 ]; then
  echo "no tags match $VERSION on $REMOTE_URL"
  exit 0
fi

echo "==> ${#TAGS[@]} matching tag(s):"
printf '    %s\n' "${TAGS[@]}"

if [ "$CONFIRM" != true ]; then
  echo
  echo "dry run. re-run with --yes to delete."
  exit 0
fi

REFSPECS=()
for t in "${TAGS[@]}"; do
  REFSPECS+=(":refs/tags/${t}")
done

echo "==> deleting ${#TAGS[@]} tag(s) from $REMOTE_URL"
git push "$REMOTE_URL" "${REFSPECS[@]}"
echo "==> done"
