package main

import (
	"bulletin_board/src/dbmanager"
	"bulletin_board/src/reqmanager"
	"fmt"
	"strconv"
)


func main() {
	// Test
	for i := 0; i < 50; i++ {
		_, err := dbmanager.AddElement(dbmanager.Element{ Title: "Title " + strconv.Itoa(i + 1), Text: "Text: text  text  text  text  text  text  text  text  text  text  text  text  text  text  text  text  text  text  text  text  text  text  text  text  text  text  text  text  text  text" })
		if (err != nil) {
			fmt.Println("Error" + err.Error())
		}
	}
	fmt.Println("Listening...")
	reqmanager.HandleHTTP(":8080")
}
