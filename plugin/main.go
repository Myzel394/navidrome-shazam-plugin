package main

import (
	"github.com/navidrome/navidrome/plugins/pdk/go/lyrics"
	"github.com/navidrome/navidrome/plugins/pdk/go/pdk"
)

type plugin struct{}

//go:wasmexport nd_on_init
func ndOnInit() int32 {
	pdk.Log(pdk.LogInfo, "navidrome-shazam-plugin: initialized")
	return 0
}

func init() {
	lyrics.Register(&plugin{})
}

func main() {}
