#!/usr/bin/env bash
# Shared helpers for k3s tag sync. Assumes CWD is the kubernetes repo (fork).

# read vX.Y.Z lines on stdin, print the highest
max_semver() {
  grep -E '^v[0-9]+\.[0-9]+\.[0-9]+$' | sort -V | tail -n1
}

# vX.Y.Z -> vX.Y
minor_of() {
  printf '%s\n' "$1" | sed -E 's/^(v[0-9]+\.[0-9]+)\..*/\1/'
}

# minors that already carry k3s tags in this fork
tracked_minors() {
  git tag -l '*-k3s*' \
    | sed -nE 's/^(v[0-9]+\.[0-9]+)\.[0-9]+-k3s[0-9]+$/\1/p' \
    | sort -uV
}

# vX.Y -> highest k8s patch in the k3s-io/kubernetes fork 
# that has a -k3s tag
latest_fork_k3s_k8s_version() {
  local minor="$1"
  git tag -l "${minor}.*-k3s*" \
    | sed -nE 's/^(v[0-9]+\.[0-9]+\.[0-9]+)-k3s[0-9]+$/\1/p' \
    | sort -uV | max_semver
}

# vX.Y -> highest upstream stable release tag 
# excludes rc, alpha, beta, etc.
latest_upstream_k8s_version() {
  local minor="$1"
  local url="${UPSTREAM_K8S_URL:-https://github.com/kubernetes/kubernetes.git}"
  git ls-remote --tags --refs "${url}" "refs/tags/${minor}.*" \
    | awk '{print $2}' \
    | sed -nE 's|^refs/tags/(v[0-9]+\.[0-9]+\.[0-9]+)$|\1|p' \
    | sort -uV | max_semver
}
