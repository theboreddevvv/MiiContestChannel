package main

import (
	"context"
	"fmt"
	"github.com/WiiLink24/MiiContestChannel/contest"
	"github.com/WiiLink24/MiiContestChannel/first"
	"github.com/WiiLink24/MiiContestChannel/plaza"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

var (
	ctx  = context.Background()
	pool *pgxpool.Pool
)

func checkError(err error) {
	if err != nil {
		log.Fatalf("Mii Contest Channel server has encountered a fatal error! Reason: %v\n", err)
	}
}

func main() {
	config := GetConfig()

	// Start SQL
	dbString := fmt.Sprintf("postgres://%s:%s@%s/%s", config.Username, config.Password, config.DatabaseAddress, config.DatabaseName)
	dbConf, err := pgxpool.ParseConfig(dbString)
	checkError(err)
	pool, err = pgxpool.ConnectConfig(ctx, dbConf)
	checkError(err)

	err = first.MakeFirst()
	checkError(err)

	first.MakeAddition()

	err = contest.MakeContestInfos(pool, ctx)
	checkError(err)

	err = plaza.MakeSpotList(pool, ctx)
	checkError(err)

	err = plaza.MakeBargainList(pool, ctx)
	checkError(err)

	err = plaza.MakeNewList(pool, ctx)
	checkError(err)

	err = plaza.MakeTop50(pool, ctx)
	checkError(err)

	err = plaza.MakePopCraftsList(pool, ctx)
	checkError(err)
}
