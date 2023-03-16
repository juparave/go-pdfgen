package model

import "time"

type Invoice struct {
	ID       string
	Customer Customer
	Date     *time.Time
	Articles []Article
	Total    float64
	Tax      float64
}

type Customer struct {
	Name string
}

type Article struct {
	Name  string
	Price float64
}
