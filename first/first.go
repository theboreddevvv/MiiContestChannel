package first

import (
	"github.com/WiiLink24/MiiContestChannel/common"
	"math"
)

type First struct {
	Type             common.ListTag
	DiscontinuedFlag uint8
	_                [4]byte
	CountryCode      uint8
	_                [16]byte
	Padding          [8]byte
	Tag              common.ListTag
	TagSize          uint16
	Unk              uint32
	CountryGroup     uint8
	/*
		3rd lowest bits are unknown
		4th bit: Marquee Text
		5th bit: Re-enable initials
		6th bit: Disable initials
		7th bit: Show Mii Artisan
		8th bit: Show number of Miis in Posting Plaza
	*/
	FirstBitField uint8
	_             uint8
	/*
		The below toggle the languages available for the message board service.

		1st bit: Unknown
		2nd bit: Dutch
		3rd bit: Italian
		4th bit: Spanish
		5th bit: French
		6th bit: German
		7th bit: English
		8th bit: Japanese
	*/
	SecondBitField uint8
}

func (f First) ToBytes(data any) []byte {
	return common.ToBytes(data)
}

func MakeFirst() error {
	first := First{
		Type:             common.FirstList,
		DiscontinuedFlag: 0,
		CountryCode:      110,
		Padding:          [8]byte{math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8},
		Tag:              common.FirstList,
		TagSize:          12,
		Unk:              1,
		CountryGroup:     151,
		FirstBitField:    0x10,
		SecondBitField:   154,
	}

	return common.Write(first, "first/110.ces")
}
