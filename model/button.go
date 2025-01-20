package model

// Icon represents an Icon fase on font awesome icons
type Icon struct {
	Name string `yaml:"name"`
	Type `yaml:",inline"`
}

func NewIcon(name string) *Icon {
	return &Icon{
		Name: name,
		Type: NewType("icon"),
	}
}

func (i Icon) GetName() string {
	return i.Name
}

func (i Icon) GetHTML(level int) string {
	return "<i class=\"" + i.Name + "\"></i>"
}

// Button is a type of widget that represents a button
type Button struct {
	Type    `yaml:",inline"`
	Text    string `yaml:"text"`
	Active  bool   `yaml:"active"`
	Actions `yaml:",omitempty"`
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

// ButtonGroup is a type of container that holds multiple buttons
type ButtonGroup struct {
	Container `yaml:",inline"`
}

func NewButtonGroup(name string, buttons ...*Button) *ButtonGroup {
	var bg ButtonGroup

	bg.Container = *NewContainer(name)
	bg.Type = NewType("buttongroup")

	for _, b := range buttons {
		bg.Add(b)
	}
	return &bg
}
