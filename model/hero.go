package model

type Hero struct {
	HeroId   int         `bson:"heroId,omitempty"`
	HeroName string      `bson:"heroName,omitempty"`
	HeroAttr []*HeroAttr `bson:"heroAttr,omitempty"`
}

type HeroAttr struct {
	AttrId int `bson:"attrId,omitempty"`
	Value  int `bson:"value,omitempty"`
}
