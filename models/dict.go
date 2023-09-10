package models

type Dict struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Type   int    `json:"type"`
	DictId string `json:"dict_id"`
}
