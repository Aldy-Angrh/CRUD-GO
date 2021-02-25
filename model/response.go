package model

//Response is a struct
type Response struct {
	ResponseCode string      `json:"RC"`
	Deskripsi    string      `json:"Deskripsi"`
	Data         interface{} `json:"Data"`
}

//DataResponse is a struct
type DataResponse struct {
	DeskripsiError string `json:"error"`
	DateFromat     string `json:"issue_at"`
}
