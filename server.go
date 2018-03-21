package guuid

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"

    "net/http"
    "time"
    "os"
)

const (
    _ 		= iota
    GET
    POST
)

type webServer struct {
    addr			string
    readTimeout		time.Duration
    writeTimeout	time.Duration
    echo			*echo.Echo
}

func NewServer(addr string, readTimeout time.Duration, writeTimeout time.Duration) *webServer {
    s := &webServer{
        addr: addr,
        readTimeout: readTimeout,
        writeTimeout: writeTimeout,
        echo: echo.New(),
    }
    return s
}

func (s *webServer) Start() {
    hs := &http.Server{
        Addr: s.addr,
        ReadTimeout: s.readTimeout,
        WriteTimeout: s.writeTimeout,
    }
    s.echo.HideBanner = true

    // echo middleware setting
    s.echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
        Format: `[${time_rfc3339}] {"time":${time_unix},"id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
        `"method":"${method}","uri":"${uri}","status":${status}, "latency":${latency},` +
        `"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
        `"bytes_out":${bytes_out}}` + "\n",
        Output: os.Stdout,
    }))

    s.echo.Use(middleware.GzipWithConfig(middleware.GzipConfig{
        Level: 5,
    }))

    s.echo.Logger.Fatal(s.echo.StartServer(hs))
}

func (s *webServer) Handler(router string, method int, h func(echo.Context) error) {
    switch method {
    case GET:
        s.echo.GET(router, h)
    case POST:
        s.echo.POST(router, h)
    default:
        s.echo.GET(router, h)
    }
}

func (s *webServer) HandlerError(h func(error, echo.Context)) {
    s.echo.HTTPErrorHandler = h
}
