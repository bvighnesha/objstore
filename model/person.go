package model

import (
	"reflect"
	"time"
)

type Person struct {
	Name      string    `json:"name"`
	ID        string    `json:"id"`
	LastName  string    `json:"last_name"`
	Birthday  string    `json:"birthday"`
	BirthDate time.Time `json:"birthday"`
}

func (p *Person) GetKind() string {
	return reflect.TypeOf(p).String()
}
func (p *Person) GetID() string {
	return p.ID
}
func (p *Person) GetName() string {
	return p.Name
}
func (p *Person) SetID(s string) {
	p.ID = s
}
func (p *Person) SetName(s string) {
	p.Name = s
}
