package main

import (
	"encoding/json"
	"fmt"
)

func chooseFunctionPOST(f string, jString string) {
	n := parseJSONToNote(jString)
	switch f {
	case "addNote":
		err := _addNote(n)
		if err != nil {
			fmt.Println(err)
		}
	case "updateNote":
		err := _updateNote(n)
		if err != nil {
			fmt.Println(err)
		}
	case "deleteNote":
		err := _deleteNote(n.id)
		if err != nil {
			fmt.Println(err)
		}
	}
}
func chooseFunctionGet(f string, id int) {
	switch f {
	case "getNote":
		n, err := _getNote(id)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(n)
	case "getNotes":
		n, err := _getNotes()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(n)
	}
}

func parseJSONToNote(jString string) (n note) {
	json.Unmarshal([]byte(jString), &n)
	return n
}
