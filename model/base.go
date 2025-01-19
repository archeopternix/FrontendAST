package model

import (
	"fmt"
)

// Widget is the basic type that can be used to represent a widget and added to containers
type Widget interface {
	GetName() string
	GetType() string
}

type Widgets map[string]Widget

// Type indicates the type of the widget or container
type Type struct {
	Type string
}

func (t Type) GetType() string {
	return t.Type
}

func NewType(name string) Type {
	return Type{Type: name}
}

// Container is the basic type that can hold multiple widgets
type Container struct {
	Name    string
	Type    `yaml:",inline"`
	Widgets `yaml:",omitempty"`
	Styles  `yaml:"styles,omitempty"`
}

func NewContainer(name string) *Container {
	return &Container{
		Name:    name,
		Widgets: make(Widgets),
		Type:    NewType("container"),
	}
}

func (c Container) GetName() string {
	return c.Name
}

func (c Container) Get(key string) *Widget {
	if ret, ok := c.Widgets[key]; ok {
		return &ret
	}
	return nil
}

func (c Container) Add(child Widget) error {
	if _, ok := c.Widgets[child.GetName()]; ok {
		return fmt.Errorf("Child %s already exists", child.GetName())
	}
	c.Widgets[child.GetName()] = child
	return nil
}

func (c Container) GetAll() Widgets {
	return c.Widgets
}

// Layout is a type of container to arrange widgets
type Layout struct {
	Container `yaml:",inline"`
}

func NewLayout(name string) *Layout {
	c := NewContainer(name)
	c.Type = NewType("layout")
	return &Layout{Container: *c}
}

func (l Layout) GetType() string {
	return l.Type.GetType()
}

// Styles represents the styling elements of a widget
type Styles []string

func MakeStyles(styles ...string) Styles {
	var s []string
	for _, style := range styles {
		s = append(s, style)
	}
	return s
}

var ActionStates = []string{
	"onChange()",
	"onClick()",
	"onSelect()",
}

type Actions []string

func MakeActions(actions ...string) Actions {
	var s []string
	for _, action := range actions {
		s = append(s, action)
	}
	return s
}
