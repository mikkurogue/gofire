package ui

import "errors"

type Friend struct {
	Avatar          string `json: "avatar"`
	Status          bool   `json: "online_status"`
	Name            string `json: "name`
	CurrentActivity string `json: "current_activity"`
}

func FriendListItem() error {
	return errors.New("nyi")
}
