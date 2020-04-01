package model

import "fmt"

type Match struct {
	PlatformId string `json:"platformId"`
	Id         uint64 `json:"gameId"`
	ChampionId uint16 `json:"champion"`
	Queue      uint16 `json:"queue"`
	Season     uint8 `json:"season"`
	Timestamp  uint64 `json:"timestamp"`
	Role       string `json:"role"`
	Lane       string `json:"lane"`
}

func (m Match) ToString() string {
	return fmt.Sprintf("GameID:%d,ChampionID:%d",
		m.Id, m.ChampionId)
}

func (m Match) ToJSON() string {
	return fmt.Sprintf("{\"gameId\":\"%d\",\"championId\":\"%d\"}",
	m.Id, m.ChampionId)
}
