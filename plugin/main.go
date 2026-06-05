package main

import "github.com/extism/go-pdk"

//go:wasmexport nd_on_init
func ndOnInit() int32 {
	pdk.Log(pdk.LogInfo, "hello world from navidrome-lyrics-scrape plugin")
	return 0
}

func main() {}
