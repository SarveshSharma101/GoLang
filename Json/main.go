package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func main() {
	u := user{
		Name:   "Sarvesh",
		Age:    26,
		Gender: "male",
	}

	b, err := json.Marshal(u)
	if err != nil {
		fmt.Println("error ", err)
	} else {
		fmt.Println("json ", string(b))
	}

	var newUser user

	err = json.Unmarshal(b, &newUser)
	if err != nil {
		fmt.Println("error ", err)
	} else {
		fmt.Println("new user ", newUser)
	}

	b, err = json.MarshalIndent(newUser, "", "\t")
	if err != nil {
		fmt.Println("error ", err)
	} else {
		fmt.Println("indented json ", string(b))
	}
}
