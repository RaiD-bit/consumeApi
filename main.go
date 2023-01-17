package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var client *http.Client

type CatFact struct {
	Fact   string
	Length int
}

func GetCatFact() {
	url := "https://catfact.ninja/fact"
	var catfact CatFact
	err := GetJson(url, &catfact)
	if err != nil {
		fmt.Println("got an error while calling catfact X( ")
		return
	} else {
		fmt.Printf("here is an interesting fact brought to you by catfact: %v\n", catfact.Fact)
	}
}

//  complex api
type RandomUser struct {
	Results []UserResult
}

type UserResult struct {
	Name    Username
	Email   string
	Picture UserPic
}

func (u UserResult) PrintUser() string {
	res := fmt.Sprintf("name : %s %s %s \nEmail: %s\nPicture url {\n\tLarge: %s \n\tMedium: %s \n\tThumbnail: %s\n} ",
		u.Name.Title, u.Name.First, u.Name.Last, u.Email, u.Picture.Large, u.Picture.Medium, u.Picture.Thumbnail)
	return res
}

type Username struct {
	Title string
	First string
	Last  string
}

type UserPic struct {
	Large     string
	Medium    string
	Thumbnail string
}

func GetRandomUser() {
	url := "https://randomuser.me/api/?inc=name,email,picture"

	var user RandomUser

	err := GetJson(url, &user)
	if err != nil {
		fmt.Println("Some error found while calling the api")
	} else {
		fmt.Printf("You very own random user -> \n%v\n", user.Results[0].PrintUser())
	}
}

func GetJson(url string, target interface{}) error {
	res, err := client.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(target)
}

func main() {
	client = &http.Client{Timeout: 10 * time.Second}
	// GetCatFact()
	GetRandomUser()
}
