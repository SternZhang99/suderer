package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type BasicInfo struct {
	CharacterList []Character `json:"character_list"`
}

type Character struct {
	CharacterId string     `json:"character_id"`
	FactionId   int        `json:"faction_id"` //1VS,2NC,3TR,4NSO
	Name        name       `json:"name"`
	Times       times      `json:"times"`
	Certs       certs      `json:"certs"`
	BattleRank  battleRank `json:"battle_rank"`
}

type name struct {
	First string `json:"first"`
}

type times struct {
	CreationDate  string `json:"creation_date"`
	LastSaveDate  string `json:"last_save_date"`
	LastLoginDate string `json:"last_login_date"`
}

type certs struct {
	EarnedPoints    int `json:"earned_points"`
	AvailablePoints int `json:"available_points"`
}

type battleRank struct {
	PercentToNext int `json:"percent_to_next"`
	Value         int `json:"value"`
}

func FetchBasicInfo(url string) BasicInfo {
	resp, err := http.Get(url)
	var userBasic BasicInfo
	if err != nil {
		log.Println(err)
		return userBasic
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return userBasic
	}
	err = json.Unmarshal(body, &userBasic)
	if err != nil {
		log.Println(err)
	}
	return userBasic

}
