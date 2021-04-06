package entities

import (
	"encoding/json"
	"fmt"
)

// Constructor
func NewStudent(id int, firstName string, lastName string, age int, languageCode string) *Student {
	r := new(Student)
	r.Id = id
	r.FirstName = firstName
	r.LastName = lastName
	r.Age = age
	r.LanguageCode = languageCode
	return r
}

// Student type
type Student struct {
	Id           int
	FirstName    string
	LastName     string
	Age          int
	LanguageCode string
}

// Stringer interface implementation
func (this *Student) String() string {
	return fmt.Sprintf(
		"[%d : %s, %s, %d, %s]",
		this.Id,
		this.FirstName,
		this.LastName,
		this.Age,
		this.LanguageCode)
}

func (this *Student) ToJSON() []byte {
	jsonBytes, err := json.Marshal(this)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

func (this *Student) FromJSON(jsonBytes []byte) {
	err := json.Unmarshal(jsonBytes, this)
	if err != nil {
		print(err)
		return
	}
}
