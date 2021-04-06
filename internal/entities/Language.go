package entities

import "fmt"

// Constructor
func NewLanguage(code string, name string) *Language {
	r := new(Language)
	r.Code = code
	r.Name = name
	return r
}

// Language type
type Language struct {
	Code string
	Name string
}

// Stringer interface implementation
func (this Language) String() string {
	return fmt.Sprintf(
		"[%s : %s]",
		this.Code,
		this.Name)
}
