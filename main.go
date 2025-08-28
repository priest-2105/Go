package main

import (
	"fmt"
	// "strconv"
)

// var i float32 = 42

// func main (){

// 	fmt.Printf("%v, %T", i, i)

// }

// func main(){

// var i int
// i = 42
// var j float32 = 27
// k := 99.0

// fmt.Println(i)
// fmt.Println(j)
// fmt.Println(k)
// fmt.Printf("%v, %T", k, k)
// }

// var (
// 	actorName string = "Elizabeth olsen"
// 	companion string = "Frank smith"
// 	doctorNumber int = 3
// 	season int = 11
// )

// var (
// 	 counter int = 0
// )

// var i int = 27

// func main () {
// 	 var i int = 42
// 	//  i = 13
// 	 fmt.Println(i)
// }

// func main () {
// 	 var i int = 42
// 	 j := 13

// 	 fmt.Println(i)
// }

// func main (){

// 	var i int = 42
// 	fmt.Printf("%v, %T/n", i, i)

// 	var j string
// 	j = strconv.Itoa(i)
// 	fmt.Printf("%v, %T/n", j, j)

// }

// func main() {
// 	var n bool = false
// 	fmt.Printf("%v, %T\n", n, n)
// }

// func main() {

// 	n := 1 == 1
// 	m := 1 == 2
// 	fmt.Printf("%v, %T\n", n, n)
// 	fmt.Printf("%v, %T\n", m, m)

// }

// func main() {

// 	var n uint16 = 42
// 	fmt.Printf("%v, %T\n", n, n)

// }

// func main() {
// 	a := 10.2
// 	b := 3.3
// 	fmt.Println(a + b)
// 	fmt.Println(a - b)
// 	fmt.Println(a * b)
// 	fmt.Println(a / b)
// 	// fmt.Println(a % b)
// }

// func main () {
// 	var a  int = 10
// 	var b int8 = 3
// 	fmt.Println(int8(a) + b)
// }

// bit operators

// func main() {
// 	a := 10
// 	b := 3
// 	fmt.Println(a & b)
// 	fmt.Println(a | b)
// 	fmt.Println(a ^ b)
// 	fmt.Println(a &^ b)
// }

// func main() {
// 	a := 8
// 	fmt.Println(a << 3)
// 	fmt.Println(a >> 3)
// }

// func main (){
// 	var n complex64 = 1 + 2i
// 	fmt.Printf("%v, %T\n", real(n), real(n))
// 	fmt.Printf("%v, %T\n", imag(n), imag(n))
// }

// func main (){
// 	var n complex64 = complex(5, 12)
// 	fmt.Printf("%v, %T\n", n, n)
// }

// func main (){
// 	s := "This is a string"
// 	fmt.Printf("%v, %T\n", string(s[2]), s[2])
// }

// func main(){
// 	    const myConst int = 42
// 		fmt.Printf("%v, %T\n", myConst, myConst)
// }

// const a int16 = 27

// func main() {
// 	const a int = 14
// 	const b string = "foo"
// 	const c float32 = 3.14
// 	const d bool = true

//     fmt.Printf("%v\n", a)
//     fmt.Printf("%v\n", b)
//     fmt.Printf("%v\n", c)
//     fmt.Printf("%v\n", d)
// }

// func main(){
// 	const a int = 42
// 	var b int16 = 27
// 	fmt.Printf("%v, %T\n", a + b, a)
// }

// const (
// 	a = iota
// 	b = iota
// 	c = iota
// )

// func main(){
// 	fmt.Printf("%v\n", a)
// 	fmt.Printf("%v\n", b)
// 	fmt.Printf("%v\n", c)
// }

// const (
// 	a = iota
// 	b
// 	c
// )

// const (
// 	 a2 = iota
// )

// func main(){
// 	fmt.Printf("%v\n", a)
// 	fmt.Printf("%v\n", b)
// 	fmt.Printf("%v\n", c)
// 	fmt.Printf("%v\n", a2)
// }

// const (
// 	errorSpecialist = iota
// 	catSpecialist
// 	dogSpecialist
// 	snakeSpecialist
// )

// func main() {
// 	  var specialistType int
// 	  fmt.Printf("%v\n", specialistType == catSpecialist)
// }

// const (
// 	_ = iota
// 	kB = 1 << (10 * iota)
// 	MB
// 	GB
// 	TB
// 	PB
// 	EB
// 	ZB
// 	YB
// )

// func main() {

// 	fileSize := 4000000000.
// 	fmt.Printf("%.2fGB", fileSize/GB)

// }

// func main (){
// 	    // grade1 := 97
// 	    // grade2 := 98
// 	    // grade3 := 45

// 		grades := [3]int{ 97, 85, 93}

// 		fmt.Printf("Grades: %v", grades)
// }

// func main() {
// 	var students [3]string
// 	fmt.Printf("Grades: %v", students)
// 	students[0] = "lisa"
// 	students[1] = "ahmed"
// 	students[2] = "arnold"

// 	fmt.Printf("Students: %v", students)
// 	fmt.Printf("Number of Student: %v\n", len(students))
// }


// func main() {
// 	a := [...]int{1, 2, 3}
// 	b := a
// 	b[1] = 5
// 	fmt.Println(a)
// 	fmt.Println(b)
// }


// func main() {
// 	a := [...]int{1, 2, 3}
// 	b := &a
// 	b[1] = 5
// 	fmt.Println(a)
// 	fmt.Println(b)
// }


// func main() {
// 	a := []int{1, 2, 3}
// 	// b := &a
// 	// b[1] = 5
// 	fmt.Println(a)
// 	// fmt.Println(b)
// 	fmt.Printf("Length: %v\n", len(a))
// 	fmt.Printf("Capacity: %v\n", cap(a))
// }


// func main () {
// 	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	b := a[:]
// 	c := a[3:]
// 	d := a[:6]
// 	e := a[3:6]
// 	a[5] = 42
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// 	fmt.Println(e)
// }
// func main() {
// 	// statePoplations := make(map[string]int, 10)
// 	statePoplations := map[string]int{
// 		"California": 39250017,
// 		"Texas": 39250017,
// 		"Florida": 39250017,
// 		"New York": 39250017,
// 		"Peennsylvania": 39250017,
// 		"Illinois": 39250017,
// 		"Ohio": 39250017,
// 	}
// 	// m := map[[3]int]string{}
// 	// delete(statePoplations, "Georgia")
// 	// fmt.Println(statePoplations)
// 	// fmt.Println(statePoplations["Georgia"])
// 	pop , ok := statePoplations["Ohio"]
// 	fmt.Println(pop, ok)
// }

// func main() {
// 	if true {
// 		fmt.Println("The test is true")
// 	}
// }


func main (){
	 number := 50
	 guess := 990

	 if guess < 1 || guess >  100 {
		fmt.Println("The guess must be between 1 and 100 !")
	 }

	  if guess >= 1 && guess <=  100 {
		fmt.Println("The guess must be between 1 and 100 !")
	 if guess < number {
		fmt.Println("Too Low")
	 }
	 if guess > number {
		fmt.Println("Too High")
	 }
	 if guess == number {
		fmt.Println("Go it !")
	 }
	
	 fmt.Println(number<=guess, number>=guess, number!=guess) 
	}
	fmt.Println(true)
}

