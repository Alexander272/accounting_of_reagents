package http

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"path"
	"runtime/debug"
	"strings"
	"time"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/config"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/models/response"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/services"
	"github.com/Alexander272/accounting_of_reagents/backend/internal/transport/http/middleware"
	httpV1 "github.com/Alexander272/accounting_of_reagents/backend/internal/transport/http/v1"
	"github.com/Alexander272/accounting_of_reagents/backend/pkg/accept_encoding"
	"github.com/Alexander272/accounting_of_reagents/backend/pkg/auth"
	"github.com/Alexander272/accounting_of_reagents/backend/pkg/error_bot"
	"github.com/Alexander272/accounting_of_reagents/backend/pkg/limiter"
	"github.com/Alexander272/accounting_of_reagents/backend/pkg/logger"
	"github.com/Alexander272/accounting_of_reagents/backend/web"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	// permissions casbin.Casbin
	keycloak *auth.KeycloakClient
	services *services.Services
}

func NewHandler(services *services.Services, keycloak *auth.KeycloakClient) *Handler {
	return &Handler{
		services: services,
		keycloak: keycloak,
		// permissions: permissions,
	}
}

func (h *Handler) Init(conf *config.Config) *gin.Engine {
	router := gin.New()

	// Отключаем редиректы для SPA
	router.RedirectTrailingSlash = false
	router.RedirectFixedPath = false

	router.Use(
		gin.LoggerWithConfig(gin.LoggerConfig{
			Skip: func(c *gin.Context) bool {
				path := c.Request.URL.Path
				if strings.HasPrefix(path, "/api") {
					return false
				}
				return c.Writer.Status() < http.StatusBadRequest
			},
		}),
		gin.CustomRecovery(h.ErrorHandler),
		securityHeaders(),
	)

	if err := router.SetTrustedProxies([]string{"192.168.5.0/24", "192.168.4.0/24"}); err != nil {
		logger.Warn("invalid trusted proxies config", logger.ErrAttr(err))
	}

	h.initAPI(router, conf)
	h.initStatic(router, conf)

	return router
}

func securityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		c.Header("Content-Security-Policy",
			"default-src 'self'; script-src 'self'; style-src 'self' 'unsafe-inline'; "+
				"img-src 'self' data: blob:; font-src 'self' data:; connect-src 'self' ws: wss:; "+
				"frame-ancestors 'none'; base-uri 'self'; form-action 'self';")
		c.Header("Permissions-Policy",
			"camera=(), microphone=(), geolocation=(), gyroscope=(), "+
				"accelerometer=(), magnetometer=(), usb=(), payment=(), "+
				"display-capture=()")
		c.Next()
	}
}

func (h *Handler) ErrorHandler(c *gin.Context, origErr any) {
	err := fmt.Errorf("unexpected error: %v", origErr)

	rawStack := string(debug.Stack())                        // 1. Получаем стек в виде байтов
	cleanStack := strings.ReplaceAll(rawStack, "\t", "    ") // 2. Заменяем все табуляции на 4 пробела для красоты
	stackLines := strings.Split(cleanStack, "\n")            // 3. Превращаем в срез строк, разделяя по символу \n

	error_bot.Send(c, err.Error(), gin.H{"Stack trace": stackLines})
	response.NewErrorResponse(c, http.StatusInternalServerError, err.Error(), "Произошла непредвиденная ошибка: "+err.Error())
}

func (h *Handler) initAPI(router *gin.Engine, conf *config.Config) {
	middleware := middleware.NewMiddleware(h.services, &conf.Auth, h.keycloak)
	handlerV1 := httpV1.NewHandler(httpV1.Deps{Services: h.services, Conf: conf, Middleware: middleware})

	api := router.Group("/api")
	{
		api.Use(limiter.Limit(conf.ApiLimiter.RPS, conf.ApiLimiter.Burst, conf.ApiLimiter.TTL))
		api.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
		handlerV1.Init(api)
	}
}

var appStartTime = time.Now()

const (
	frontendRoot = "frontend"
	indexFile    = "index.html"
	assetsPrefix = "assets/"
)

// staticFS — встроенный фронтенд (backend/web/frontend) отдельным FS с корнем в
// этой папке: путь запроса не нужно префиксовать вручную.
var staticFS = mustSub(web.Frontend, frontendRoot)

// mustSub — fs.Sub для константного пути: ошибка означает ошибку сборки.
func mustSub(fsys fs.FS, dir string) fs.FS {
	sub, err := fs.Sub(fsys, dir)
	if err != nil {
		panic("static fs: " + err.Error())
	}
	return sub
}

func (h *Handler) initStatic(router *gin.Engine, conf *config.Config) {
	static := &staticHandler{fsys: staticFS}
	router.NoRoute(limiter.Limit(conf.StaticLimiter.RPS, conf.StaticLimiter.Burst, conf.StaticLimiter.TTL), func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.Status(http.StatusNotFound)
			return
		}
		static.serve(c)
	})
}

// staticHandler раздаёт встроенный фронтенд: файлы сборки и оболочку SPA для
// маршрутов клиента.
type staticHandler struct {
	fsys fs.FS
}

// serve отдаёт файл статики, а для путей, которых в сборке нет, — оболочку SPA.
func (s *staticHandler) serve(c *gin.Context) {
	filePath := staticPath(c.Request.URL.Path)
	encoding := accept_encoding.Negotiate(c.Request.Header.Get("Accept-Encoding"))

	f, enc, err := openStatic(s.fsys, filePath, encoding)
	if err != nil {
		// Файла нет: путь без расширения — маршрут SPA (отдаём оболочку), с
		// расширением — запрос файла (404: HTML вместо картинки не подсовываем).
		if path.Ext(filePath) != "" {
			c.Status(http.StatusNotFound)
			return
		}
		f, enc, err = openStatic(s.fsys, indexFile, encoding)
		if err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		filePath = indexFile
	}
	defer f.Close()

	// Каталог не отдаём: embed отдаёт для него «папку», а читать её нельзя.
	if info, statErr := f.Stat(); statErr == nil && info.IsDir() {
		c.Status(http.StatusNotFound)
		return
	}
	rs, ok := f.(io.ReadSeeker)
	if !ok {
		c.Status(http.StatusInternalServerError)
		return
	}

	// Содержимое зависит от Accept-Encoding (сжатый вариант или обычный) — кэш
	// обязан это учитывать.
	c.Header("Vary", "Accept-Encoding")
	if enc != "" {
		c.Header("Content-Encoding", enc)
	}
	if strings.HasPrefix(filePath, assetsPrefix) {
		// Имена ассетов — с хешем: содержимое не меняется, кэшируем надолго.
		c.Header("Cache-Control", "public, max-age=31536000, immutable")
	} else {
		c.Header("Cache-Control", "no-cache")
	}

	// В ServeContent передаём имя без суффикса сжатия: иначе Content-Type
	// определится как application/gzip (или бинарные данные) и браузер не выполнит скрипт.
	http.ServeContent(c.Writer, c.Request, path.Base(staticName(filePath, enc)), appStartTime, rs)
}

// staticPath превращает путь запроса в относительный путь внутри FS: без
// ведущего слэша, без «..». Дальше путь проверяет сам FS: embed.FS отбрасывает
// невалидные пути (fs.ValidPath), поэтому выйти за пределы сборки нельзя.
func staticPath(urlPath string) string {
	p := strings.TrimPrefix(path.Clean("/"+urlPath), "/")
	if p == "" || p == "." {
		return indexFile
	}
	return p
}

// openStatic открывает файл статики с учётом Accept-Encoding: если рядом есть
// заранее сжатый вариант (.br/.gz) и клиент его принимает — отдаём его с
// соответствующим Content-Encoding. Сжатые варианты делает сборка образа
// (Dockerfile: gzip -9 + brotli -q 11).
func openStatic(fsys fs.FS, filePath, acceptEncoding string) (fs.File, string, error) {
	if suffix := compressedSuffix(acceptEncoding); suffix != "" {
		if f, err := fsys.Open(filePath + suffix); err == nil {
			return f, acceptEncoding, nil
		}
	}
	f, err := fsys.Open(filePath)
	return f, "", err
}

// compressedSuffix — расширение заранее сжатого варианта файла для кодирования
// («gzip» кодируется файлом «.gz»).
func compressedSuffix(encoding string) string {
	switch encoding {
	case "br":
		return ".br"
	case "gzip":
		return ".gz"
	}
	return ""
}

// staticName возвращает имя файла без суффикса сжатия — оно нужно ServeContent
// для корректного определения Content-Type.
func staticName(filePath, encoding string) string {
	if suffix := compressedSuffix(encoding); suffix != "" {
		return strings.TrimSuffix(filePath, suffix)
	}
	return filePath
}
