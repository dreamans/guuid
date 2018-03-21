// This file is part of the guuid package
//
// (c) Dreamans <dreamans@163.com>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package guuid

import (
    "github.com/labstack/echo"

    "net/http"
    "strconv"
)

type Response struct {
    Code	int		`json:"code"`
    Message	string	`json:"message"`
}

type ResponseItem struct {
    Response
    Uuid 		string	`json:"uuid"`
}

type ResponseData struct {
    Response
    Data 		[]string	`json:"data"`
}

func VersionHandler(c echo.Context) error {
    return c.JSON(http.StatusOK, &Response{
        Code: http.StatusOK,
        Message: "Version " + VERSION,
    })
}

func ErrorHandler(err error, c echo.Context) {
    var code int
    if he, ok := err.(*echo.HTTPError); ok {
        code = he.Code
    }
    msg := http.StatusText(code)
    if msg == "" {
        code = http.StatusInternalServerError
        msg = http.StatusText(code)
    }

    c.JSON(code, &Response{
        Code: code,
        Message: msg,
    })

    c.Logger().Error(err)
}

func CreateHandler(c echo.Context) error {
    resp := createSingleUUID(false)
    return c.JSON(http.StatusOK, resp)
}

func CreateSimpleHandler(c echo.Context) error {
    resp := createSingleUUID(true)
    return c.JSON(http.StatusOK, resp)
}

func createSingleUUID(isSimple bool) *ResponseItem {
    resp := &ResponseItem{}
    resp.Code = http.StatusOK
    resp.Message = http.StatusText(http.StatusOK)
    if isSimple {
        resp.Uuid = createSimpleUUID()
    } else {
        resp.Uuid = createUUID()
    }
    return resp
}

func CreateMultiHandler(c echo.Context) error {
    num, _ := strconv.Atoi(c.Param("num"))
    resp := createMultiUUID(num, false)
    return c.JSON(http.StatusOK, resp)
}

func CreateMultiSimpleHandler(c echo.Context) error {
    num, _ := strconv.Atoi(c.Param("num"))
    resp := createMultiUUID(num, true)
    return c.JSON(http.StatusOK, resp)
}

func createMultiUUID(num int, isSimple bool) *ResponseData {
    if num < 1 {
        num = 1
    }
    if num > 1000 {
        num = 1000
    }
    var (
        index = 0
        resp = &ResponseData{}
    )
    resp.Code = http.StatusOK
    resp.Message = http.StatusText(http.StatusOK)

    for {
        if index >= num {
            break
        }
        index++

        var uuid string
        if isSimple {
            uuid = createSimpleUUID()
        } else {
            uuid = createUUID()
        }
        resp.Data = append(resp.Data, uuid)
    }

    return resp
}
