package model


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
