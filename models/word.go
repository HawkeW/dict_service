package models

type Word struct {
	Id         uint   `json:"id"`
	Word       string `json:"word"`
	LinkWordId string `json:"link_word_id"`
	DictId     string `json:"dict_id"`
	PronUk     string `json:"pron_uk"`
	PronUs     string `json:"pron_us"`
	Captions   string `json:"captions"`
}
