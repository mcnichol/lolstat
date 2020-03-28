package model

import "fmt"

type Summoner struct {
	Id            string `json:"id"`
	AccountId     string `json:"accountId"`
	Puuid         string `json:"puuid"`
	Name          string `json:"name"`
	ProfileIconId uint16 `json:"profileIconId"`
	RevisionDate  uint64 `json:"revisionDate"`
	Level         uint8  `json:"summonerLevel"`
}

func (s Summoner) ToString() string {
	return fmt.Sprintf(
		"Summoner Info:\n\n"+
			"%15v:%50s\n"+
			"%15v:%50d\n"+
			"%15v:%50v\n"+
			"%15v:%50v\n",

		"Name", s.Name, "Level", s.Level, "Account ID", s.AccountId, "Summoner ID", s.Id)
}
