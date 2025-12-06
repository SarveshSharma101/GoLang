package main

import "fmt"

func main() {
	Arrays()
	Slices()
	Maps()
}

func Arrays() {
	fmt.Println("------------- ARRAYS --------------")

	fmt.Println("Declaring array and printing it")
	var a [3]int
	a[0] = 100
	a[1] = 101
	a[2] = 102
	fmt.Println("Array a: ", a)

	fmt.Println("")

	fmt.Println("declaring array using := ")
	b := [3]int{}
	b[1] = 1011
	fmt.Println("array b: ", b)

	fmt.Println("letting compiler decide the length of array using '...' operator")

	c := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println("array c: ", c)

	fmt.Println("length of all array")

	fmt.Println("len of a: ", len(a))
	fmt.Println("len of b: ", len(b))
	fmt.Println("len of c: ", len(c))

	fmt.Println("")

	fmt.Println("Multidimensional arrays")
	m := [4][5]int{}
	m[0][1] = 1
	m[2][2] = 1
	m[3][3] = 1
	fmt.Println("multidimensional array m :", m)
	fmt.Println("len of multi-dimensional array:", len(m))
}

func Slices() {
	fmt.Println("")
	fmt.Println("------------- SLICES --------------")

	fmt.Println("Deriving slice from an existing array")
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:4]
	fmt.Println("array is a and its len is, cap is: ", a, len(a), cap(a))
	fmt.Println("derived slice is s: ", s)
	fmt.Println("len of slice: ", len(s))

	fmt.Println("capacity of slice, usually more than  len as it represents underlying array: ", cap(s))

	fmt.Println("declaring slice using :=")
	p := []int{1, 2, 3, 4}
	fmt.Println("p = ", p)
	fmt.Println("appending to slice")
	p = append(p, 100)
	fmt.Println("new p: ", p)

	fmt.Println("merging slice p and s")
	p = append(p, s...)
	fmt.Println("merged p: ", p)

}

func Maps() {
	fmt.Println("")
	fmt.Println("------------- MAPS --------------")
	var m map[int]int
	fmt.Println("zero value of map=nil:", m == nil)

	fmt.Println("len of zero map: ", len(m))
	fmt.Println("map of array as key and its len as values")

	a := [3]int{1, 2, 3}
	b := [3]int{2, 2, 3}

	var u map[[3]int]int = map[[3]int]int{
		a: len(a),
		b: len(b),
	}
	fmt.Println("map: ", u)

	fmt.Println("declaring map using make keyword")
	n := make(map[int]string)
	n[1] = "one"
	n[2] = "two"
	fmt.Println("map n:", n)
	fmt.Println("checking if 2 exist in n")
	_, ok := n[2]
	fmt.Println("-->", ok)
	fmt.Println("deleting 2 from map n")
	delete(n, 2)

	fmt.Println("checking if 2 exist now or not")
	_, ok = n[2]
	fmt.Println("-->", ok)
	fmt.Println("len of map", len(n))
}
