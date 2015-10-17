package main

import (
	"fmt"
	"github.com/gocubes/config"
	"log"
)

type Work struct {
	WorkName string `json:"name"`
	WorkUrl  string `json:"url"`
}

type User struct {
	Useranme string `json:"name"`
	Homepage string `json:"site"`
	Github   string `json:"github"`
	Works    []Work `json:"works"`
}

func main() {
	provier, perr := config.New("get.json", "json")
	if perr != nil {
		log.Panicf("create config provider error: %v\n", perr.Error())
	}

	user := &User{}
	gerr := provier.Get(user)
	if gerr != nil {
		log.Panicf("get config data error: %v\n", gerr.Error())
	}

	fmt.Printf("Name:\t %s\nHomepage:\t %s\nGithub:\t %s\nWorks:\n", user.Useranme, user.Homepage, user.Github)
	for i, work := range user.Works {
		fmt.Printf("\t%d. %s\t %s\n", i+1, work.WorkName, work.WorkUrl)
	}
}
