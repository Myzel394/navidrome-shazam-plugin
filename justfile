plugin_name := "navidrome-lyrics-scrape"
data_dir    := "navidrome-instance/data"
plugins_dir := data_dir / "plugins"

username := "admin"
password := "password"

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

# Fetch lyrics via official SubSonic API for a random song
fetch-lyrics:
    #!/usr/bin/env bash
    set -euo pipefail
    auth='u={{username}}&p={{password}}&v=1.16.0&c=test&f=json'
    song_id=$(
        curl -s "http://localhost:4533/rest/getRandomSongs?$auth" \
        | jq -r '.["subsonic-response"].randomSongs.song[0].id'
    )
    lyrics=$(
        curl -s "http://localhost:4533/rest/getLyricsBySongId?id=$song_id&$auth" \
        | jq -r '.["subsonic-response"].lyricsList.structuredLyrics[0].line | map(.value) | join("\n")'
    )

    echo "$lyrics"


