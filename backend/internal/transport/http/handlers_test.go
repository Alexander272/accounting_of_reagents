package http

import (
	nethttp "net/http"
	"net/http/httptest"
	"testing"
	"testing/fstest"

	"github.com/gin-gonic/gin"
)

func TestStaticHandlerServe(t *testing.T) {
	gin.SetMode(gin.TestMode)

	h := &staticHandler{fsys: fstest.MapFS{
		"assets/app.js":    {Data: []byte("console.log(1)")},
		"assets/app.js.br": {Data: []byte("br-compressed")},
		"index.html":       {Data: []byte("<html>ok</html>")},
		"index.html.gz":    {Data: []byte("gz-compressed")},
	}}

	cases := []struct {
		name      string
		path      string
		accept    string
		wantCode  int
		wantEnc   string
		wantType  string
		wantBody  string
		wantCache string
	}{
		{
			name:      "сжатый js",
			path:      "/assets/app.js",
			accept:    "br",
			wantCode:  nethttp.StatusOK,
			wantEnc:   "br",
			wantType:  "text/javascript; charset=utf-8",
			wantBody:  "br-compressed",
			wantCache: "public, max-age=31536000, immutable",
		},
		{
			name:      "сжатый index",
			path:      "/",
			accept:    "gzip",
			wantCode:  nethttp.StatusOK,
			wantEnc:   "gzip",
			wantType:  "text/html; charset=utf-8",
			wantBody:  "gz-compressed",
			wantCache: "no-cache",
		},
		{
			name:      "spa fallback",
			path:      "/report",
			wantCode:  nethttp.StatusOK,
			wantType:  "text/html; charset=utf-8",
			wantBody:  "<html>ok</html>",
			wantCache: "no-cache",
		},
		{
			name:     "несуществующий файл",
			path:     "/assets/missing.js",
			wantCode: nethttp.StatusNotFound,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest(nethttp.MethodGet, tc.path, nil)
			if tc.accept != "" {
				c.Request.Header.Set("Accept-Encoding", tc.accept)
			}

			h.serve(c)
			// В реальном запросе gin записывает статус после цепочки middleware.
			c.Writer.WriteHeaderNow()

			if w.Code != tc.wantCode {
				t.Fatalf("статус: хочу %d, получил %d", tc.wantCode, w.Code)
			}
			if got := w.Header().Get("Content-Encoding"); got != tc.wantEnc {
				t.Errorf("Content-Encoding: хочу %q, получил %q", tc.wantEnc, got)
			}
			if got := w.Header().Get("Content-Type"); got != tc.wantType {
				t.Errorf("Content-Type: хочу %q, получил %q", tc.wantType, got)
			}
			if got := w.Body.String(); got != tc.wantBody {
				t.Errorf("тело: хочу %q, получил %q", tc.wantBody, got)
			}
			if got := w.Header().Get("Cache-Control"); got != tc.wantCache {
				t.Errorf("Cache-Control: хочу %q, получил %q", tc.wantCache, got)
			}
			if tc.wantCode == nethttp.StatusOK && w.Header().Get("Vary") != "Accept-Encoding" {
				t.Errorf("ожидался заголовок Vary: Accept-Encoding")
			}
		})
	}
}
