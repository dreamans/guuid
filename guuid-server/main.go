// This file is part of the guuid package
//
// (c) Dreamans <dreamans@163.com>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package main

import (
    "github.com/dreamans/guuid"
    "github.com/labstack/echo"

    "time"
    "flag"
    "strconv"
)

type routerItem struct {
    url		string
    method 	int
    handler	func(echo.Context) error
}

func initRouter() []*routerItem {
    return []*routerItem{
        &routerItem{
            url: "/version",
            handler: guuid.VersionHandler,
        },
        &routerItem{
            url: "/get",
            handler: guuid.CreateHandler,
        },
        &routerItem{
            url: "/get/simple",
            handler: guuid.CreateSimpleHandler,
        },
        &routerItem{
            url: "/mget/:num",
            handler: guuid.CreateMultiHandler,
        },
        &routerItem{
            url: "/mget/:num/simple",
            handler: guuid.CreateMultiSimpleHandler,
        },
    }
}

func main() {

    port := flag.Int("port", 12345, "guuid http server port")
    timeout := flag.Int("timeout", 5, "connect timeout")
    flag.Parse()

    s := guuid.NewServer(":" + strconv.Itoa(*port), time.Second * time.Duration(*timeout), time.Second * time.Duration(*timeout))

    s.HandlerError(guuid.ErrorHandler)

    for _, r := range initRouter() {
        s.Handler(r.url, r.method, r.handler)
    }

    s.Start()
}
