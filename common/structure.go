package common

import (
	"strconv"
)

// MiiWithArtisan contains Mii data with its artisan.
type MiiWithArtisan struct {
	EntryNumber     uint32
	MiiData         []byte
	Initials        string
	Likes           uint8
	Skill           uint16
	CountryCode     uint8
	ArtisanId       uint32
	IsMasterArtisan uint8
	ArtisanMiiData  []byte
}

func ZFill(value int, size int) string {
	str := strconv.FormatInt(int64(value), 10)
	temp := ""

	for i := 0; i < size-len(str); i++ {
		temp += "0"
	}

	return temp + str
}
