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
	return fmt.Sprintf("SummonerID:%s,AccountID:%s,Name:%s", s.Id, s.AccountId, s.Name)
}

func (s Summoner) ToJSON() string {
	return fmt.Sprintf("{\"accountId\":\"%s\", \"id\":\"%s\", \"level\": \"%d\", \"name\":\"%s\", \"puuid\":\"%s\"}",
		s.AccountId, s.Id, s.Level, s.Name, s.Puuid)
}
