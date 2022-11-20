package handlers

import "github.com/go-playground/validator/v10"

var flasher Flasher
var validate *validator.Validate

type Flasher struct {
	Type    string
	Message string
}

func (f *Flasher) Set(_type, msg string) {
	f.Message = msg
	f.Type = _type
}

func (f *Flasher) Del() {
	f.Message = ""
	f.Type = ""
}
