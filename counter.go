// This file is part of the guuid package
//
// (c) Dreamans <dreamans@163.com>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package guuid

import (
    "sync/atomic"
)

var idCounter uint32 = 0

func deltaIdCounter() uint32 {
    return atomic.AddUint32(&idCounter, 1)
}
