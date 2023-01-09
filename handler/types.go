package handler

type JSONUser struct {
	Email        string `json:"email"`
	Name         string `json:"name"`
	Organization string `json:"organization"`
	Tag          string `json:"tag"`
}
