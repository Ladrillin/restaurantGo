package model

type Order struct {
	Id         int    `json:"Id"`
	CustomerId int    `json:"CustomerId"`
	Content    string `json:"Content"`
}

type Orders struct {
	Orders []Order `json:"orders"`
}

type Meals struct {
	Salad []Salad
	Soup  []Soup
	Drink []Drink
}

type Drink struct {
	Alcohol []Alcohol
	Cola    string
	Water   string
}

type Alcohol struct {
	Whiskey string
	Vodka   string
	Beer    string
	Wine    string
}

type Soup struct {
	IndianSoup string
	PhoBo      string
	NoodleSoup string
	CreamSoup  string
}

type Salad struct {
	VegetableSalad string
	CaesarSalad    string
	GreekSalad     string
}
