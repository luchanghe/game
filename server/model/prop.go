package model

type Prop struct {
	PropId  int `bson:"propId,omitempty" json:"propId"`
	PropNum int `bson:"propNum,omitempty" json:"propNum"`
}
