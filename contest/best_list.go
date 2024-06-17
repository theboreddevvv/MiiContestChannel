package contest

import (
	"bytes"
	"context"
	"fmt"
	"github.com/WiiLink24/MiiContestChannel/common"
	"github.com/jackc/pgx/v4/pgxpool"
	"math"
)

type BestList struct {
	Header BestListHeader
	Miis   []BestListMiiPair
}

type BestListHeader struct {
	Tag            common.ListTag
	_              uint16
	ContestId      uint32
	_              uint32
	_              uint32
	_              [12]byte
	Padding        [4]byte
	ContestTag     common.ListTag
	ContestTagSize uint16
	Unknown        uint32
	NumberOfMiis   uint32
}

type BestListMiiPair struct {
	Mii     BestListMii
	Artisan BestListArtisan
}

type BestListMii struct {
	Tag         common.ListTag
	TagSize     uint16
	MiiIndex    uint32
	EntryNumber uint32
	MiiData     [76]byte
}

type BestListArtisan struct {
	Tag                common.ListTag
	TagSize            uint16
	ArtisanIndex       uint32
	ArtisanNumber      uint32
	MiiArtisanData     [76]byte
	_                  uint8
	IsMasterMiiArtisan uint8
	_                  [3]byte
	CountryCode        uint8
	MiiIndex           uint16
}

func (b BestList) ToBytes(_ any) []byte {
	buffer := new(bytes.Buffer)
	common.WriteBinary(buffer, b.Header)

	for _, pair := range b.Miis {
		common.WriteBinary(buffer, pair.Mii)
		common.WriteBinary(buffer, pair.Artisan)
	}

	return buffer.Bytes()
}

func MakeBestList(pool *pgxpool.Pool, ctx context.Context, contestId uint32) error {
	rows, err := pool.Query(ctx, GetBestContestMiis, contestId)
	if err != nil {
		return err
	}

	bestList := BestList{
		Header: BestListHeader{
			Tag:            common.BestList,
			ContestId:      contestId,
			Padding:        [4]byte{math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8},
			ContestTag:     common.ContestTag,
			ContestTagSize: 12,
			Unknown:        1,
			NumberOfMiis:   0,
		},
		Miis: nil,
	}

	var miis []BestListMiiPair
	var index uint32 = 1
	for rows.Next() {
		mii := BestListMii{
			Tag:         common.ContestMii,
			TagSize:     88,
			MiiIndex:    index,
			EntryNumber: 0,
			MiiData:     [76]byte{},
		}

		artisan := BestListArtisan{
			Tag:                common.ContestArtisan,
			TagSize:            96,
			ArtisanIndex:       index,
			ArtisanNumber:      0,
			MiiArtisanData:     [76]byte{},
			IsMasterMiiArtisan: 0,
			CountryCode:        0,
			MiiIndex:           uint16(index),
		}

		var miiData []byte
		var artisanMiiData []byte
		var isMaster bool
		err = rows.Scan(&mii.EntryNumber, &artisan.ArtisanNumber, &miiData, &artisanMiiData, &artisan.CountryCode, &isMaster)
		if err != nil {
			return err
		}

		if isMaster {
			artisan.IsMasterMiiArtisan = 1
		}

		copy(mii.MiiData[:], miiData)
		copy(artisan.MiiArtisanData[:], artisanMiiData)

		miis = append(miis, BestListMiiPair{
			Mii:     mii,
			Artisan: artisan,
		})
		index++
	}

	bestList.Header.NumberOfMiis = uint32(len(miis))
	bestList.Miis = miis

	return common.Write(bestList, fmt.Sprintf("contest/%d/best_list.ces", contestId))
}
