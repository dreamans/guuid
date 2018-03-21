// This file is part of the guuid package
//
// (c) Dreamans <dreamans@163.com>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package guuid

type gUid struct {
    machine []byte
    time 	uint32
    rand 	uint32
    id 		uint32
}

const GUID_LENGTH = 36

const VERSION = "1.0.0"

func createUUID() string {
    gid := make([]byte, 0, GUID_LENGTH)

    id := []byte(byteToHexString(createBaseId()))
	
    gid = append(gid, id[0:8]...)
    gid = append(gid, '-')
    gid = append(gid, id[8:12]...)
    gid = append(gid, '-')
    gid = append(gid, id[12:16]...)
    gid = append(gid, '-')
    gid = append(gid, id[16:20]...)
    gid = append(gid, '-')
    gid = append(gid, id[20:]...)

    return string(gid)
}

func createSimpleUUID() string {
    return byteToHexString(createBaseId())
}

func createBaseId() []byte {
    guid := newGuid()
    return guid.serial()
}

func newGuid() *gUid {
    guid := &gUid{
        machine: machineId,
        time: timeStamp(),
        rand: rand(),
        id: deltaIdCounter(),
    }
    return guid
}

func (u *gUid) serial() []byte {
    id := make([]byte, 0, GUID_LENGTH)

    id = append(id, u.machine[0:4]...)
    id = append(id, uint32ToBin(u.time)...)
    id = append(id, uint32ToBin(u.id)...)
    id = append(id, uint32ToBin(u.rand)...)

    return id
}