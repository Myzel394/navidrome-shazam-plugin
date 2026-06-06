# Update go module dependencies
[group('plugin'), working-directory: "plugin"]
tidy:
    go mod tidy

# Format and vet all Go source
[group('plugin'), working-directory: "plugin"]
lint-plugin:
    gofumpt -w .

# Run project tests
[group('plugin'), working-directory: "plugin"]
test:
    go test ./...

# Remove all build artifacts: wasm, ndp, zips, and the installed plugin.
[group('plugin'), working-directory: "plugin"]
clear:
    rm -f plugin.wasm {{plugin_name}}.ndp *.zip
    rm -f {{plugins_dir}}/plugin.wasm {{plugins_dir}}/manifest.json {{plugins_dir}}/{{plugin_name}}.ndp

# Optimized production wasm build (small binary).
[group('plugin'), working-directory: "plugin"]
build:
    tinygo build -o plugin.wasm -target wasip1 -buildmode=c-shared .

# Non-minified wasm build: disables optimizations and keeps debug
[group('plugin'), working-directory: "plugin"]
build-dev:
    tinygo build -o plugin.wasm -target wasip1 -buildmode=c-shared -opt=0 .

# Bundle wasm + manifest into a distributable .ndp archive.
[group('plugin'), working-directory: "plugin"]
pack:
    zip -j {{plugin_name}}.ndp manifest.json plugin.wasm
