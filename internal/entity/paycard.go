package entity

type PayCard struct {
	Id         int64
	Secret     Secret
	CardNumber string
	Holder     string
	Dates      string
}
