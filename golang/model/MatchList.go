package model

import "fmt"

type MatchList struct {
	Matches    []Match `json:"matches"`
	StartIndex uint `json:"startIndex"`
	EndIndex   uint `json:"endIndex"`
	TotalGames uint `json:"totalGames"`
}

func (ml MatchList) ToString() string {
	return fmt.Sprintf("Matches:%v", ml.Matches)
}

func (ml MatchList) ToStringSlice(slice uint) string {
	return fmt.Sprintf("Matches:%v", ml.Matches[0:slice])
}