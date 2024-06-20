package plaza

import (
	"context"
	"github.com/WiiLink24/MiiContestChannel/common"
	"github.com/jackc/pgx/v4/pgxpool"
	"math"
)

const (
	GetNumberOfMiis     = `SELECT COUNT(*) FROM miis`
	GetNumberOfArtisans = `SELECT COUNT(*) FROM artisans`
)

type NumberInfo struct {
	Tag                 common.ListTag
	_                   uint16
	CountryRegion       uint32
	ListNumber          uint32
	ErrorCode           uint32
	_                   [12]byte
	Padding             [4]byte
	NumberInfoTag       [2]byte
	TagSize             uint16
	Unknown             uint32
	NumberOfMiis        uint32
	NumberOfMiiArtisans uint32
}

func (f NumberInfo) ToBytes(data any) []byte {
	return common.ToBytes(data)
}

func MakeNumberInfo(pool *pgxpool.Pool, ctx context.Context) error {
	var numberOfMiis uint32
	err := pool.QueryRow(ctx, GetNumberOfMiis).Scan(&numberOfMiis)
	if err != nil {
		return err
	}

	var numberOfArtisans uint32
	err = pool.QueryRow(ctx, GetNumberOfArtisans).Scan(&numberOfArtisans)
	if err != nil {
		return err
	}

	numberInfo := NumberInfo{
		Tag:                 common.NumberInfo,
		CountryRegion:       151,
		ListNumber:          0,
		ErrorCode:           0,
		Padding:             [4]byte{math.MaxUint8, math.MaxUint8, math.MaxUint8, math.MaxUint8},
		NumberInfoTag:       common.NumberInfo,
		TagSize:             16,
		Unknown:             1,
		NumberOfMiis:        numberOfMiis,
		NumberOfMiiArtisans: numberOfArtisans,
	}

	return common.Write(numberInfo, "151/number_info.ces")
}
