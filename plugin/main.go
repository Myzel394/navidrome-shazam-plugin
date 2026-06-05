package main

import ("github.com/extism/go-pdk"
	"github.com/navidrome/navidrome/plugins/pdk/go/lyrics"
)

type plugin struct{}


//go:wasmexport nd_on_init
func ndOnInit() int32 {
	pdk.Log(pdk.LogInfo, "Hello from navidrome-lyrics-scrape")
	return 0
}

func init() {
	lyrics.Register(&plugin{})
}
	
func main() {}
