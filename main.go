package preview

import (
	"embed"
	"mime"
	"net/http"
	"path/filepath"
	"strings"

	. "m7s.live/engine/v4"
)

//go:embed ui
var f embed.FS

type PreviewConfig struct {
}

func (p *PreviewConfig) OnEvent(event any) {

}

var plugin = InstallPlugin(&PreviewConfig{})

func (p *PreviewConfig) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ss := strings.Split(r.URL.Path, "/")
	if b, err := f.ReadFile("ui/" + ss[len(ss)-1]); err == nil {
		w.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(ss[len(ss)-1])))
		w.Write(b)
	} else {
		w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
		w.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
		b, err = f.ReadFile("ui/index.html")
		w.Write(b)
	}
}
