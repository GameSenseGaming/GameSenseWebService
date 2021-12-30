package api

import (
	"fmt"
)

const apiKey = "RGAPI-80c01b2a-8d54-4e00-a049-38acbf41167c"
const leagueOfLegendsRegion = "https://na1.api.riotgames.com"

type summonerDto struct {
	accountId     string
	profileIconId int
	revisionDate  uint64
	name          string
	id            string
	puuid         string
	summonerLevel uint64
}

func GetSummonerBySummonerName(summonerName string) {

	url := leagueOfLegendsRegion + "/lol/summoner/v4/summoners/by-name/" + summonerName + "?api_key=" + apiKey
	data := HttpGet(url)
	fmt.Println(string(data))

	// Map data to struct

	// Send to GraphQL Server
}
