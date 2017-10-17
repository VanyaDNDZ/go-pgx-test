package forms

import "time"

type VFSForm struct {
	Filename string `form:"filename" json:"filename" binding:"required"`
	Attributes interface{} `form:"Attributes" json:"Attributes" binding:"required"`
}

type VFSResponse struct {
	Userid string
	Fileid string
	Filename string
	Creation_date time.Time
	Attributes interface{}
}