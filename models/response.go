package models

import "sort"

type weapon struct {
	weaponName   string
	kills        float64
	accuracy     float64
	headshotRate float64
}

type Response struct {
	Name    string
	Weapons []weapon
}

func SortWeapon(resp Response, key string) {
	if key == "kills" {
		sort.Slice(resp.Weapons, func(i, j int) bool {
			return resp.Weapons[i].kills > resp.Weapons[j].kills
		})
	}
	if key == "accuracy" {
		sort.Slice(resp.Weapons, func(i, j int) bool {
			return resp.Weapons[i].accuracy > resp.Weapons[j].accuracy
		})
	}
	if key == "headshotRate" {
		sort.Slice(resp.Weapons, func(i, j int) bool {
			return resp.Weapons[i].headshotRate > resp.Weapons[j].headshotRate
		})
	}
}
