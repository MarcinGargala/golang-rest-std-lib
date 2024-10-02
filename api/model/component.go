package model

type Component struct {
	Name       string
	Type       string
	Metadata   map[string]string
	Properties map[string]interface{}
	Constrains string
}
