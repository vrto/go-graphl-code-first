package main

type Pub struct {
	Id    int
	Name  string
	Beers []Beer
}

type Beer struct {
	Id    int
	Maker string
	Name  string
	Type  string
}
