package plaza

import (
	"github.com/WiiLink24/MiiContestChannel/common"
	"math"
)

type Info struct {
	Tag           common.ListTag
	_             uint16
	CountryRegion uint32
	EntryNumber   uint32
	ErrorCode     uint32
	_             [12]byte
	Padding       [4]byte
	MiiInfo
	ArtisanInfo
}

type MiiInfo struct {
	Tag         common.ListTag
	TagSize     uint16
	Unknown     uint32
	EntryNumber uint32
	MiiData     [76]byte
	Unknown2    uint16
	_           uint32
	Initials    [2]byte
}

type ArtisanInfo struct {
	Tag        common.ListTag
	TagSize    uint16
	Unknown    uint32
	Unknown1   uint32
	Unknown2   uint32
	IsMaster   uint16
	Popularity uint8
	_          uint8
	Ranking    uint8
	_          uint16
	Unknown3   uint8
}

func (i *Info) ToBytes(_ any) []byte {
	return common.ToBytes(i)
}

func MakeArtisanInfo(entryNumber uint32, miiInfo *MiiInfo, artisanInfo *ArtisanInfo) []byte {
	info := Info{
		Tag:           common.ArtisanInfo,
		CountryRegion: 0,
		EntryNumber:   entryNumber,
		ErrorCode:     0,
		Padding:       [4]byte{math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8},
		MiiInfo:       *miiInfo,
		ArtisanInfo:   *artisanInfo,
	}

	return info.ToBytes(nil)
}
