// This file is part of the guuid package
//
// (c) Dreamans <dreamans@163.com>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package guuid

import (
    "encoding/hex"
    "encoding/binary"
    "time"
    mrand "math/rand"
)

func uint32ToHexString(n uint32) string {
    return byteToHexString(uint32ToBin(n))
}

func byteToHexString(n []byte) string {
    return hex.EncodeToString(n)
}

func uint32ToBin(n uint32) []byte {
    uintByte := make([]byte, 4, 4)
    binary.BigEndian.PutUint32(uintByte, n)
    return uintByte
}

func timeStamp() uint32 {
    return uint32(time.Now().Unix())
}

func rand() uint32 {
    return uint32(mrand.Int31())
}