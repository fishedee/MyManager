package user;

type Users struct{
	Count int `json:"count"`
	Data []*User `json:"data"`
}