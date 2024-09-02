package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Story map[string]Chapter

type Chapter struct{
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
		
}

func loadStory() map[string]Chapter{
	data, err := os.ReadFile("./story.json")
	if err != nil{
		fmt.Println("error while reading file")
		os.Exit(0)
	}
	loadStory := Story{}
	errJson := json.Unmarshal(data,&loadStory)
	if errJson != nil{
		fmt.Println("error while unmarshalling the story")
		os.Exit(0)
	}
	//fmt.Println(loadStory["intro"].Options)
	return loadStory

}

func main(){
	story_map := loadStory()
	fmt.Println(story_map)	

}