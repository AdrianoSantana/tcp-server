package dto

type Request struct {
	Id int `json:"id"`
	Action string `json:"action"`
	Body any `json:"body"`
}