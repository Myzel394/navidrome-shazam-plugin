#!/usr/bin/env just --justfile

# Dev install: copies raw wasm + manifest to the plugins dir
# without zipping. Faster iteration — no repacking needed.
[group('dev')]
dev-install:
    @just build-dev
    @just pack
    mkdir -p {{plugins_dir}}
    cp plugin/{{plugin_name}}.ndp {{plugins_dir}}/

[group('dev')]
format:
    gofumpt -w plugin
    treefmt .

# Vet Go packages.
[group('dev'), working-directory: 'plugin']
lint:
    go vet ./...
