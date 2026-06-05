plugin_name := "navidrome-lyrics-scrape"
data_dir    := "navidrome-instance/data"
plugins_dir := data_dir / "plugins"

default:
    @just --list

[working-directory: "plugin"]
tidy:
    go mod tidy

[working-directory: "plugin"]
build:
    tinygo build -o plugin.wasm -target wasip1 -buildmode=c-shared .

[working-directory: "plugin"]
pack: build
    zip -j {{plugin_name}}.ndp manifest.json plugin.wasm

install: pack
    mkdir -p {{plugins_dir}}
    cp plugin/{{plugin_name}}.ndp {{plugins_dir}}/

[working-directory: "plugin"]
lint:
    gofumpt -w .
    go vet ./...

[working-directory: "plugin"]
test:
    go test ./...

[working-directory: "plugin"]
clear:
    rm -f plugin.wasm {{plugin_name}}.ndp

