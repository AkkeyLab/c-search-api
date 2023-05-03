package main

type company struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var companies = []company{
	{ID: "1", Name: "AkkeyLab, inc."},
	{ID: "2", Name: "Apple, inc."},
}
