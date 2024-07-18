package plaza

import (
	_ "embed"
	"github.com/WiiLink24/MiiContestChannel/common"
	"math"
)

func MakeNameSearch(miis []common.MiiWithArtisan, entryNumber uint32) []byte {
	list := MiiList{}

	header := Header{
		Tag:           common.NameSearch,
		CountryRegion: 0,
		ListNumber:    entryNumber,
		ErrorCode:     0,
		Padding:       [4]byte{math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8},
		MiiTag:        common.MiiPairNumber,
		MiiTagSize:    12,
		Unknown:       1,
		NumberOfMiis:  uint32(len(miis)),
	}

	var miiPair []MiiPair
	for i, mii := range miis {
		var tempMiiData [76]byte
		var tempArtisanMiiData [76]byte
		var tempInitials [2]byte

		copy(tempMiiData[:], mii.MiiData)
		copy(tempArtisanMiiData[:], mii.ArtisanMiiData)
		copy(tempInitials[:], mii.Initials)

		pair := MiiPair{
			Mii: Mii{
				Tag:         common.PairMii,
				TagSize:     96,
				MiiIndex:    uint32(i + 1),
				EntryNumber: mii.EntryNumber,
				MiiData:     tempMiiData,
				Popularity:  mii.Likes,
				Skill:       mii.Skill,
				Initials:    tempInitials,
			},
			Artisan: Artisan{
				Tag:                common.PairArtisan,
				TagSize:            96,
				ArtisanIndex:       uint32(i + 1),
				ArtisanNumber:      mii.ArtisanId,
				MiiArtisanData:     tempArtisanMiiData,
				IsMasterMiiArtisan: mii.IsMasterArtisan,
				CountryCode:        mii.CountryCode,
			},
		}

		miiPair = append(miiPair, pair)
	}

	list.Header = header
	list.Miis = miiPair
	return list.ToBytes(nil)
}
