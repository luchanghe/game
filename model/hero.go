package model

type Hero struct {
	HeroId   int
	HeroName string
	HeroAttr []HeroAttr
}

type HeroAttr struct {
	attrId int
	value  int
}
