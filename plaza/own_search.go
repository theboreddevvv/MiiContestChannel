package plaza

import (
	"bytes"
	_ "embed"
	"github.com/WiiLink24/MiiContestChannel/common"
	"math"
)

// OwnSearch is the data the server requests when searching for a Mii artisan.
// It is essentially MiiList but without the artisan.
type OwnSearch struct {
	Header Header
	Miis   []Mii
}

func (o OwnSearch) ToBytes(_ any) []byte {
	buffer := new(bytes.Buffer)
	common.WriteBinary(buffer, o.Header)

	for _, mii := range o.Miis {
		common.WriteBinary(buffer, mii)
	}

	return buffer.Bytes()
}

func MakeOwnSearch(miis []common.MiiWithArtisan, artisanId uint32) []byte {
	list := OwnSearch{}

	header := Header{
		Tag:           common.OwnSearch,
		CountryRegion: 0,
		ListNumber:    artisanId,
		ErrorCode:     0,
		Padding:       [4]byte{math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8},
		MiiTag:        common.MiiPairNumber,
		MiiTagSize:    12,
		Unknown:       1,
		NumberOfMiis:  uint32(len(miis)),
	}

	var _miis []Mii
	for i, mii := range miis {
		var tempMiiData [76]byte
		var tempInitials [2]byte

		copy(tempMiiData[:], mii.MiiData)
		copy(tempInitials[:], mii.Initials)

		_mii := Mii{
			Tag:         common.PairMii,
			TagSize:     96,
			MiiIndex:    uint32(i + 1),
			EntryNumber: mii.EntryNumber,
			MiiData:     tempMiiData,
			Popularity:  mii.Likes,
			Skill:       mii.Skill,
			Initials:    tempInitials,
		}

		_miis = append(_miis, _mii)
	}

	list.Header = header
	list.Miis = _miis
	return list.ToBytes(nil)
}
