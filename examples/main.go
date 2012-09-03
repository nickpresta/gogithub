package main

import (
	"flag"
	"fmt"
	"github.com/NickPresta/gogithub"
)

func main() {
	username := flag.String("username", "", "GitHub Username")
	password := flag.String("password", "", "GitHub Password")
	githubUser := flag.String("user", "", "GitHub User")
	flag.Parse()

	if *githubUser == "" {
		panic("You must specify a GitHub user to query.")
	}

	var credentials map[string]string
	if *username != "" || *password != "" {
		credentials = map[string]string{"username": *username, "password": *password}
	}

	client, err := gogithub.Client(credentials)
	if err != nil {
		fmt.Println(err)
		return
	}

	user, err := client.GetUser(*githubUser)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user)

	emails, err := client.GetEmails()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(emails)
}
