package main

//
//import (
//	"strings"
//)
//
//type Singleton interface {
//	GetInstance() *Singleton
//}
//
//type Person struct {
//	FirstName string
//	LastName  string
//}
//
//func (p *Person) GetName() string {
//	return p.FirstName + " " + p.LastName
//}
//
//func (p *Person) SetName(name string) {
//	names := strings.Split(name, " ")
//	namesLen := len(names)
//	if namesLen == 0 {
//		panic("Invalid name")
//	} else if namesLen == 1 {
//		p.FirstName = names[0]
//	} else {
//		p.FirstName = names[0]
//		p.LastName = names[1]
//	}
//}
//
//func (p *Person) GetInstance() *Singleton {
//	if person == nil {
//		person = &Person{}
//	}
//	return &person
//}
//
//var person Singleton = nil
//
//func main() {
//	var s1 *Singleton = (*Person).GetInstance(&Person{
//		FirstName: "John",
//		LastName:  "Doe",
//	})
//
//	if s1 == nil {
//		panic("Singleton not created")
//	}
//
//	switch s1.(type) {
//	case *Person:
//		println("Person")
//	}
//
//	s1.(*Person).SetName("John Doe")
//
//	s1.SetName("John Doe")
//	println(s1.GetName())
//
//	s2 := GetInstance()
//	println(s2.GetName())
//}
