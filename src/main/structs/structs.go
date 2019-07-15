package structs

import "time"

type News struct {
	Author string
	Body   string
}

type Paging struct {
	TotalPage  int
	JumlahData int
}

type NewsReq struct {
	Author string `json:"author"`
	Body   string `json:"body"`
}
type Response struct {
	Message string
}

type NewsModel struct {
	Id      int       `gorm:"Primary_key,AUTO_INCREMENT"`
	Author  string    `gorm:"type:text"`
	Body    string    `gorm:"type:text"`
	Created time.Time `gorm:"DEFAULT:current_timestamp"`
}

type NewsResponse struct {
	Id     string `form:"id" json:"id"`
	Author string `form:"author" json:"author"`
	Body   string `form:"body" json:"body"`
}

type ResponseNews struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	TotalData int    `json:"totalData"`
	TotalPage int    `json:"totalPage"`
	Data      []NewsResponse
}
