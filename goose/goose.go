package goose

import (
	"fmt"
)

type Goose struct {
	Eyes         string
	Hair         string
	Age          uint
	Intelligence uint
	Height       float32
	Weight       float32
	Name         string
}

func (goose Goose) Describe() string {
	description := fmt.Sprintf("%v el Ganso te saluda\n%v pesa %f y mide %f\nTiene unos ojazos %s, y un sedoso copetazo %s\n",
		goose.Name,
		goose.Name,
		goose.Weight,
		goose.Height,
		goose.Eyes,
		goose.Hair)
	if goose.Intelligence <= 5 {
		description += "Eso si, es un poco tontorron\n"
	} else if goose.Intelligence <= 10 {
		description += "Es un ganso brillante!\n"
	} else {
		description += "...Tiene ideas interesantes\n"
	}
	description += fmt.Sprintf("A %v le gusta mucho fofear", goose.Name)
	return description
}
