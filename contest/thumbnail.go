package contest

import (
	"bytes"
	"fmt"
	"github.com/WiiLink24/MiiContestChannel/common"
	"math"
)

type Photo struct {
	Header PhotoHeader
	Data   []byte
}

type PhotoHeader struct {
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

func (t Photo) ToBytes(_ any) []byte {
	buffer := new(bytes.Buffer)
	common.WriteBinary(buffer, t.Header)
	common.WriteBinary(buffer, t.Data)

	return buffer.Bytes()
}

func MakePhoto(photoType common.ListTag, data []byte, contestId uint32) error {
	photo := Photo{
		Header: PhotoHeader{
			Type:         photoType,
			ContestId:    contestId,
			Padding:      [8]byte{math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8},
			ThumbnailTag: photoType,
			TagSize:      uint16(len(data) + 32),
			Unknown:      1,
		},
		Data: data,
	}

	if photoType == common.Thumbnail {
		return common.Write(photo, fmt.Sprintf("contest/%d/thumbnail.ces", contestId))
	} else {
		return common.Write(photo, fmt.Sprintf("contest/%d/photo.ces", contestId))
	}
}
