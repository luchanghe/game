package model

type Hero struct {
	HeroId   int         `json:"heroId"`
	HeroName string      `json:"heroName"`
	HeroAttr []*HeroAttr `json:"heroAttr"`
}

type HeroAttr struct {
	AttrId int `json:"attrId"`
	Value  int `json:"value"`
}
