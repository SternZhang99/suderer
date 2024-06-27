package models

type WeaponStatByFaction struct {
	ItemId   string `json:"item_id"`
	StatName string `json:"stat_name"`
	ValueVS  int    `json:"value_vs"`
	ValueNC  int    `json:"value_nc"`
	ValueTr  int    `json:"value_tr"`
}
