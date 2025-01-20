package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	. "frontendast/model"
)


func main() {
	w := NewContainer("APP")
	w.Styles = MakeStyles("bg-blue-500", "text-white")
	w.Add(NewLabel("Hello World"))
	w.Add(NewLabel("Master of the Universe"))

	bg := NewButtonGroup( 
		"buttons",
		NewButton("Submit"),
		NewButton("Cancel"),
		NewButton("...more"), 
	)
	w.Add(bg)
	l := NewLayout("HeroTiles")
	w.Add(l)
	l.Add(NewLabel("Google"))
	l.Styles = MakeStyles("is-header")
	w.Add(NewIcon("fas fa-home"))

	// Convert the struct to JSON
	jsonData, err := yaml.Marshal(w)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
	}
	// Print the JSON string
	fmt.Println(string(jsonData))

}
