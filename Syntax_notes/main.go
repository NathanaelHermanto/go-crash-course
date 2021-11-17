package main

import "fmt"

func main() {
	//variables must be used otherwise => error
	var x int = 5
	y := 5

	//print with var. println ada jg
	fmt.Printf("Hello World %d\n", sum(x, y)) //call a function

	//arrays
	//var array [5]int //declare array with length 5
	a := []int{5, 4, 3, 2, 1}           // declare and assign
	fmt.Printf("a[1:3] = %d\n", a[1:3]) // 4,3

	//maps (like dict in python, map in java)
	maps := make(map[string]int)
	maps["hi"] = 1
	maps["ho"] = 3
	maps["hu"] = 6

	fmt.Println(maps)
	fmt.Println(maps["hu"])

	//conditionals
	//if
	color := "red"
	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println(color)
	}

	// switch case
	/*switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println(color)
	}*/

	//loops
	//java, c style
	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}
	//using index
	ids := []int{12, 23, 45, 56, 234, 1}
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}
	//no need index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}
	//range with map
	for key, value := range maps {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
	for k := range maps {
		fmt.Println("Key: " + k)
	}

	//pointers
	v := 5
	b := &v // v's address memory
	fmt.Println(v, b)

	//Use * to read value from address
	fmt.Println(*b)
	fmt.Println(*&v)

	// Change val with pointer
	*b = 10
	fmt.Println(v)

	//call closures
	sum := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}

	// Init person using struct
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}
	// Alternative
	//person2 := Person{"Bob", "Johnson", "New York", "m", 30}

	// fmt.Println(person1.firstName)
	// fmt.Println(person1)

	person1.hasBirthday()
	person1.greet()
	fmt.Println(person1.age)

}

//functions
func sum(num1, num2 int) int {
	return num1 + num2
}

//closures (like lamda expressions?)
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

//struct (class is not available in go?)
type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

//value receiver
func (p Person) greet() string {
	return ("Hi, my name is " + p.firstName)
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}
