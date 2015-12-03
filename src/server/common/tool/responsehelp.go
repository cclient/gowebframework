package tool

import ()

type ResponseItem struct {
	Item interface{} `json:"item"`
}

type ResponsePage struct {
	Meta Meta        `json:"meta,omitempty" bson:"meta"`
	List interface{} `json:"list,omitempty" bson:"list"`
}

type Meta struct {
	Offset    int `json:"offset" bson:"offset"`
	Limit     int `json:"limit" bson:"limit"`
	Total     int `json:"total" bson:"total"`
	Length    int `json:"length" bson:"length"`
	Remaining int `json:"remaining" bson:"remaining"`
}

func (meta *Meta) SetRemaining() {
	if meta.Total == 0 {
		meta.Remaining = 0
	} else {
		meta.Remaining = meta.Total - meta.Offset - meta.Length
	}
}
