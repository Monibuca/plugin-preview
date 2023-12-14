package preview

import (
	"embed"
	"fmt"
	"mime"
	"net/http"
	"path/filepath"
	"strings"

	. "m7s.live/engine/v4"
	"m7s.live/engine/v4/config"
)

//go:embed ui
var f embed.FS

type PreviewConfig struct {
}

func (p *PreviewConfig) OnEvent(event any) {

}

var _ = InstallPlugin(&PreviewConfig{})

func (p *PreviewConfig) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		s := "<h1><h1><h2>Live Streams 引擎中正在发布的流</h2>"
		Streams.Range(func(streamPath string, stream *Stream) {
			s += fmt.Sprintf("<a href='%s'>%s</a> [ %s ]<br>", streamPath, streamPath, stream.GetType())
		})
		s += "<h2>pull stream on subscribe 订阅时才会触发拉流的流</h2>"
		for name, p := range Plugins {
			if pullcfg, ok := p.Config.(config.PullConfig); ok {
				pullconf := pullcfg.GetPullConfig()
				pullconf.PullOnSubLocker.RLock()
				if pullconf.PullOnSub != nil {
					s += fmt.Sprintf("<h3>%s</h3>", name)
					for streamPath, url := range pullconf.PullOnSub {
						s += fmt.Sprintf("<a href='%s'>%s</a> <-- %s<br>", streamPath, streamPath, url)
					}
				}
				pullconf.PullOnSubLocker.RUnlock()
			}
		}
		w.Write([]byte(s))
		return
	}
	ss := strings.Split(r.URL.Path, "/")
	if b, err := f.ReadFile("ui/" + ss[len(ss)-1]); err == nil {
		w.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(ss[len(ss)-1])))
		w.Write(b)
	} else {
		//w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
		//w.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
		b, err = f.ReadFile("ui/demo.html")
		w.Write(b)
	}
}
