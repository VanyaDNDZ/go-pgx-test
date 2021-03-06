package forms

import "time"

type VFSForm struct {
	Filename string `form:"filename" json:"filename" binding:"required"`
	Attributes interface{} `form: - json:"Attributes" binding:"required"`
}

type VFSResizeForm struct {
	Width uint `form:"width"`
	Height uint `form:"height"`
}

type VFSResponse struct {
	Userid string
	Fileid string
	Filename string
	Creation_date time.Time
	Attributes interface{}
}