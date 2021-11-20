package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"example.com/parserjson/structure"
)

func main() {

	var txt = "txt: "
	// defining a struct instance
	var data structure.AutoGenerated

	file, err := ioutil.ReadFile("text.json")

	if err != nil {
		fmt.Println("FileReading err: ")
		fmt.Println(err)
	}

	err = json.Unmarshal([]byte(file), &data)

	if err != nil {
		fmt.Println("Converting err: ")
		fmt.Println(err)
	}

	for i := 0; i < len(data.Response.Chunks); i++ {
		if data.Response.Chunks[i].ChannelTag == "1" {
			for j := 0; j < len(data.Response.Chunks[i].Alternatives); j++ {
				txt += "\n" + data.Response.Chunks[i].Alternatives[j].Text
			}
		}
	}

	fmt.Println(txt)

	ioutil.WriteFile("result.txt", []byte(txt), 0644)
}
