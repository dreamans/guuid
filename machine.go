// This file is part of the guuid package
//
// (c) Dreamans <dreamans@163.com>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package guuid

import (
    "os"
    "strings"
    "strconv"
    "crypto/md5"
)

var machineId []byte = machineHash()

func machineHash() (machHash []byte) {
    posixPID := posixPid()
    machineName := machineHostname()
    randId := rand()
    machine := strings.Join([]string{
        machineName,
        strconv.Itoa(posixPID),
        uint32ToHexString(randId),
    }, ",")

    md5Ctx := md5.New()
    md5Ctx.Write([]byte(machine))
    machHash = md5Ctx.Sum(nil)
    return
}

func machineHostname() string {
    host, err := os.Hostname()
    if err != nil {
        host = "localhost"
    }
    return host
}

func posixPid() int {
    return int(os.Getpid())
}
