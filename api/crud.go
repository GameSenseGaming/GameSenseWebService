package api

type LinkedAccount struct {
	id              int
	leagueAccountId int
	discordId       int
}

type LeagueOfLegendsAccount struct {
	id           int
	summonerName string
	puuid        string
	summonerId   string
	accountId    string
}

type DiscordAccount struct {
	id       int
	username string
}

func CreateAccount(linkedAccount *LinkedAccount, lolAccount *LeagueOfLegendsAccount, discordAccount *DiscordAccount) bool {
	return false
}

func CreateLeagueOfLegendsAccount(lolAccount *LeagueOfLegendsAccount) bool {
	return false
}

func UpdateLeagueOfLegendsAccount(lolAccount *LeagueOfLegendsAccount) bool {
	return false
}
