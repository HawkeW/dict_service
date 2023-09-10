package models

type WordBookData struct {
	Id         uint
	UserId     int
	WordId     uint
	WordBookId int
}

func (WordBookData) TableName() string {
	return "word_book_data"
}
