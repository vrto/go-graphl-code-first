package main

type Pub struct {
	Id    int
	Name  string
	Beers []Beer
}

type Beer struct {
	Id     int
	PubIds []int
	Maker  string
	Name   string
	Type   string
}
