package plaza

import (
	"context"
	"fmt"
	"github.com/WiiLink24/MiiContestChannel/common"
	"github.com/jackc/pgx/v4/pgxpool"
)

func MakeNewList(pool *pgxpool.Pool, ctx context.Context) error {
	var miis []common.MiiWithArtisan
	rows, err := pool.Query(ctx, GetMiis)
	if err != nil {
		return err
	}

	for rows.Next() {
		var isMaster bool
		var likes int
		mii := common.MiiWithArtisan{}
		err = rows.Scan(&mii.EntryNumber, &mii.Initials, &likes, &mii.Skill, &mii.CountryCode, &mii.MiiData,
			&mii.ArtisanMiiData, &mii.ArtisanId, &isMaster)
		if err != nil {
			return err
		}

		mii.Likes = uint8(likes)

		if isMaster {
			mii.IsMasterArtisan = 1
		}

		miis = append(miis, mii)
	}

	for i := 1; i < 11; i++ {
		listNumber := uint32(i)
		err = MakeList(common.NewList, miis, fmt.Sprintf("new_list%s.ces", common.ZFill(i, 2)), &listNumber)
		if err != nil {
			return err
		}
	}

	return nil
}
