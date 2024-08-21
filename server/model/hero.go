package model

type Hero struct {
	HeroId   int         `bson:"heroId,omitempty" json:"heroId"`
	HeroName string      `bson:"heroName,omitempty" json:"heroName"`
	HeroAttr []*HeroAttr `bson:"heroAttr,omitempty" json:"heroAttr"`
}

type HeroAttr struct {
	AttrId int `bson:"attrId,omitempty" json:"attrId"`
	Value  int `bson:"value,omitempty" json:"value"`
}
