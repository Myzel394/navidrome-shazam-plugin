plugin_name  := "navidrome-lyrics-scrape"
data_dir     := "navidrome-instance/data"
plugins_dir  := data_dir / "plugins"
compose_file := "navidrome-instance/docker-compose.yaml"

default:
    @just --list

build:
    tinygo build -o plugin.wasm -target wasip1 -buildmode=c-shared .

pack: build
    zip -j {{plugin_name}}.ndp manifest.json plugin.wasm

install: pack
    mkdir -p {{plugins_dir}}
    cp {{plugin_name}}.ndp {{plugins_dir}}/

lint:
    gofumpt -w .
    gofumpt -d .
    go vet ./...

test:
    go test ./...

clear:
    rm -f plugin.wasm {{plugin_name}}.ndp

