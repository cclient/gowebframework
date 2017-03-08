package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Account struct {
	//ID入库时很重要  有ID就必须设值才能入库，无ID 则由mongodb数据库自动生成id mgo不作任何设置
	//先在代码里生成id
	Id       bson.ObjectId `json:"id" bson:"_id" schema:"id"`
	Telphone string        `json:"telphone" bson:"telphone" schema:"telphone"`
	Email    string        `json:"email" bson:"email" schema:"email"`
	Truename string        `json:"truename" bson:"truename" schema:"truename"`
	Name     string        `json:"name" bson:"name" schema:"name"`
	//	Passwd        string        `json:"-" bson:"passwd"`
	Roles         []Role `json:"roles" bson:"roles" schema:"roles"`
	RoleCodes     []int  `json:"rolecodes" bson:"rolecodes" schema:"roles"`
	Lastlogintime int    `json:"lastlogintime" bson:"lastlogintimenew"`
	//todo
	Isadmin bool `json:"isadmin" bson:"isadmin"`
	//	IsLogin       string        `json:"isLogin" bson:"isLogin"`
	LoginNum string `json:"-" bson:"loginNum"`
	//todo
	IsLock bool `json:"-" bson:"islock"`
	//todo
	IsSuperAdmin bool   `json:"issuperadmin" bson:"issuperadmin" schema:"issuperadmin"`
	PasswdNew    string `json:"-" bson:"passwdnew" schema:"passwd"`

	Noticemessage   string `json:"noticemessage" bson:"noticemessage" schema:"noticemessage"`
	Isnoticemessage bool   `json:"isnoticemessage" bson:"isnoticemessage" schema:"isnoticemessage"`
	Noticeemail     string `json:"noticeemail" bson:"noticeemail" schema:"noticeemail"`
	Isnoticeemail   bool   `json:"isnoticeemail" bson:"isnoticeemail" schema:"isnoticeemail"`

	//		"noticemessage" : "18888888888",
	//	"isnoticemessage" : true,
	//	"noticeemail" : "name@goyoo.com",
	//	"isnoticeemail" : true,

}
