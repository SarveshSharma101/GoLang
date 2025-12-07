package main

import (
	"fmt"
	"sync"
)

func main() {
	// UserStruct()
	// ShapePolymorphism()
	// Generic()
	UsingSyncPool()
}

type Address struct {
	City string
}

func (u Address) printCity() {
	fmt.Println("-->", u.City)
}

type User struct {
	Name string
	age  int
	Address
}

func (u *User) setAge(age int) {
	u.age = age
}
func (u User) printAge() {
	fmt.Println("-->", u.age)
}

func UserStruct() {
	u := User{
		Name: "abc",
		Address: Address{
			City: "Delhi",
		},
	}
	u.setAge(10)
	u.printAge()
	u.printCity()
}

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius int
}

func (c Circle) Area() float64 {
	return 2 * 3.14 * float64(c.Radius) * float64(c.Radius)
}

type Square struct {
	Side int
}

func (c Square) Area() float64 {
	return float64(c.Side) * float64(c.Side)
}

func PrintAreaOfShape(x Shape) {
	fmt.Println("--> ", x.Area())
}

func ShapePolymorphism() {
	c := Circle{
		Radius: 10,
	}

	s := Square{
		Side: 20,
	}
	PrintAreaOfShape(c)
	PrintAreaOfShape(s)
}

func Add[T int | float32](x T, y T) {
	fmt.Println("-->", x+y)
}

func Generic() {
	Add(1, 2)
	Add(float32(1.1), 2.3)
}

func UsingSyncPool() {
	pool := &sync.Pool{
		New: func() any {
			return &User{}
		},
	}

	o1 := pool.Get().(*User)
	o1.Name = "ABC"
	fmt.Println("-->", o1.Name)

	o2 := pool.Get().(*User)
	o2.Name = "XYZ"
	fmt.Println("-->", o2.Name)

	pool.Put(o1)
	o3 := pool.Get().(*User)
	fmt.Println("-->", o3.Name)

	o4 := pool.New().(*User)
	fmt.Println("-->", o4.Name)
}
