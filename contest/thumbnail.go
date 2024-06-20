package contest

import (
	"bytes"
	"fmt"
	"github.com/WiiLink24/MiiContestChannel/common"
	"math"
)

type ThumbnailS struct {
	Header ThumbnailHeader
	Data   []byte
}

type ThumbnailHeader struct {
	Type         common.ListTag
	_            uint16
	ContestId    uint32
	_            uint32
	_            uint32
	_            [8]byte
	Padding      [8]byte
	ThumbnailTag common.ListTag
	TagSize      uint16
	Unknown      uint32
	_            [24]byte
}

func (t ThumbnailS) ToBytes(_ any) []byte {
	buffer := new(bytes.Buffer)
	common.WriteBinary(buffer, t.Header)
	common.WriteBinary(buffer, t.Data)

	return buffer.Bytes()
}

func MakeThumbnail(data []byte, contestId uint32) error {
	thumbnail := ThumbnailS{
		Header: ThumbnailHeader{
			Type:         common.Thumbnail,
			ContestId:    contestId,
			Padding:      [8]byte{math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8},
			ThumbnailTag: common.Thumbnail,
			TagSize:      uint16(len(data) + 32),
			Unknown:      1,
		},
		Data: data,
	}

	return common.Write(thumbnail, fmt.Sprintf("contest/%d/thumbnail.ces", contestId))
}
