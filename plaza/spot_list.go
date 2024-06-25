package plaza

import (
	"context"
	"github.com/WiiLink24/MiiContestChannel/common"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	GetMiis = `SELECT miis.entry_id, miis.initials, miis.perm_likes, miis.skill, miis.country_id, miis.mii_data, 
       			artisans.mii_data, artisans.artisan_id, artisans.is_master 
				FROM miis, artisans WHERE miis.artisan_id = artisans.artisan_id 
				ORDER BY miis.likes DESC LIMIT 100`
)

func MakeSpotList(pool *pgxpool.Pool, ctx context.Context) error {
	// TODO: This is generally incorrect, it should be a mix of popular and random miis.
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

		// Downcast to u8 as database contains numbers larger.
		mii.Likes = uint8(likes)

		if isMaster {
			mii.IsMasterArtisan = 1
		}

		miis = append(miis, mii)
	}

	err = MakeList(common.SpotList, miis, "spot_list.ces", nil)
	if err != nil {
		return err
	}

	return nil
}
