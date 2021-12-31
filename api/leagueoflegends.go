package api

import (
	"encoding/json"
	"fmt"
)

const apiKey = "API_KEY"

type summonerDto struct {
	accountId     string
	profileIconId int
	revisionDate  uint64
	name          string
	id            string
	puuid         string
	summonerLevel uint64
}

type leagueEntryDto struct {
	leagueId     string
	summonerId   string
	summonerName string
	queueType    string
	tier         string
	rank         string
	leaguePoints int
	wins         int
	losses       int
	hotStreak    bool
	veteran      bool
	freshBlood   bool
	inactive     bool
	miniSeries   *miniSeriesDto
}

type miniSeriesDto struct {
	losses   int
	progress string
	target   int
	wins     int
}

type accountDto struct {
	puuid    string
	gameName string // summoner name
	tagLine  string
}

func GetAccountInfoBySummonerName(summonerName string) {
	url := "https://americas.api.riotgames.com/riot/account/v1/accounts/by-riot-id/" + summonerName + "/NA1" + "?api_key=" + apiKey
	data := HttpGet(url)

	var accountObj map[string]json.RawMessage
	accountObjError := json.Unmarshal(data, &accountObj)
	ErrorHandler(accountObjError)

	var account accountDto
	puuidError := json.Unmarshal(accountObj["puuid"], &account.puuid)
	ErrorHandler(puuidError)
	gameNameError := json.Unmarshal(accountObj["gameName"], &account.gameName)
	ErrorHandler(gameNameError)
	tagLineError := json.Unmarshal(accountObj["tagLine"], &account.tagLine)
	ErrorHandler(tagLineError)

	// TO-DO: Send info to GraphQL Server (response)
}

// Entry point when a user register/link account to the web site.
func PostAccountInfoBySummonerName() {
	// TO-DO: Register info in database
}

func GetSummonerByPuuid(puuid string) {
	url := "https://na1.api.riotgames.com/lol/summoner/v4/summoners/by-puuid/" + puuid + "?api_key=" + apiKey
	data := HttpGet(url)

	var summonerObj map[string]json.RawMessage
	summonerObjError := json.Unmarshal(data, &summonerObj)
	ErrorHandler(summonerObjError)

	var summoner summonerDto
	accountIdError := json.Unmarshal(summonerObj["accountId"], &summoner.accountId)
	ErrorHandler(accountIdError)
	profileIconIdError := json.Unmarshal(summonerObj["profileIconId"], &summoner.profileIconId)
	ErrorHandler(profileIconIdError)
	revisionDateError := json.Unmarshal(summonerObj["revisionDate"], &summoner.revisionDate)
	ErrorHandler(revisionDateError)
	nameError := json.Unmarshal(summonerObj["name"], &summoner.name)
	ErrorHandler(nameError)
	idError := json.Unmarshal(summonerObj["id"], &summoner.id)
	ErrorHandler(idError)
	puuidError := json.Unmarshal(summonerObj["puuid"], &summoner.puuid)
	ErrorHandler(puuidError)
	summonerLevelError := json.Unmarshal(summonerObj["summonerLevel"], &summoner.summonerLevel)
	ErrorHandler(summonerLevelError)

	// TO-DO: Send info to GraphQL Server (response)
}

func GetRankByEncryptedSummonerId(summonerId string) {
	url := "https://na1.api.riotgames.com/lol/league/v4/entries/by-summoner/" + summonerId + "?api_key=" + apiKey
	data := HttpGet(url)
	fmt.Println(string(data))

	// TO-DO: Convert response to dto
	// TO-DO: Send info to GraphQL Server (response)
}
