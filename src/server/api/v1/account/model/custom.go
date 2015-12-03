package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Role struct {
	Province string `json:"province" bson:"province,omitempty"`
	City     string `json:"city" bson:"city,omitempty"`
	Area     string `json:"area" bson:"area,omitempty"`
	//	Area4    string `json:"area4" bson:"area4,omitempty"`
}

type AccountJustRoles struct {
	Id    bson.ObjectId `json:"id" bson:"_id" schema:"id"`
	Roles []Role        `json:"roles" bson:"roles" schema:"roles"`
}

func (r Role) ToPCAString() string {
	str := ""
	if r.Province != "" {
		str += r.Province
		if r.City != "" {
			str += "," + r.City
			if r.Area != "" {
				str += "," + r.Area
			}
		}
	}
	return str
}
