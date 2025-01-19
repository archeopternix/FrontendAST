package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"github.com/archeopternix/FrontendAST/model"
)


// ----------------------
type Label struct {
	Text string `yaml:"text"`
	Type `yaml:",inline"`
}

func NewLabel(text string) *Label {
	return &Label{
		Text: text,
		Type: NewType("label"),
	}
}

func (l Label) GetName() string {
	return l.Text
}

func (l Label) GetType() string {
	return "label"
}

type Button struct {
	Type    `yaml:",inline"`
	Text    string `yaml:"text"`
	Active  bool   `yaml:"active"`
	OnKlick string `yaml:"onKlick"`
	Actions
}

func NewButton(text string) *Button {
	return &Button{
		Text:   text,
		Active: true,
		Type:   NewType("button"),
	}
}

func (b Button) GetName() string {
	return b.Text
}

func (b Button) GetType() string {
	return "button"
}

func main() {
	w := NewContainer("APP")
	w.Styles = MakeStyles("bg-blue-500", "text-white")
	w.Add(NewLabel("Hello World"))
	w.Add(NewLabel("Master of the Universe"))
	b := NewButton("Click Me")
	b.Actions = MakeActions(ActionStates[0], ActionStates[1])
	w.Add(b)
	l := NewLayout("HeroTiles")
	w.Add(l)
	l.Add(NewLabel("Google"))
	l.Styles = MakeStyles("is-header")

	// Convert the struct to JSON
	jsonData, err := yaml.Marshal(w)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
	}
	// Print the JSON string
	fmt.Println(string(jsonData))

}
