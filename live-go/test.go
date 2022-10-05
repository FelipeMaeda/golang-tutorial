package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"live-go/exemplo"
	"net/http"
)

/*----------------------- STEP 8 -----------------------*/

// You can work with of the Types
// type CarName string
// type CarYear int

type Car struct {
	CarName string `json:"car"` // This is the tags of the Variables
	CarYear int    `json:"-"`   // This tag will be ignored in the Json
}

// This function has been hitched a Struct above
func (c *Car) drive() { // With a Pointer in the Parameter Variable, the manipulation of this Function ocurre in the memory level.
	c.CarName = "XPTO"
	fmt.Println(c.CarName, "Andou")
}

/*----------------------- STEP 10 -----------------------*/

type vehicle interface {
	start() string
}

func main() {
	/*----------------------- STEP 1 -----------------------*/

	// Printing variables
	println("Hello")

	nome := "Felipe"
	println(nome)
	nome = "Iris"
	println(nome)

	// Automatic type variables
	a := 10
	b := "Hello"
	c := 10.44
	d := false
	e := 'w'
	f := `Wesley
	Felipe
	Teste`

	// Printing type of the Variables
	fmt.Printf("%T,\n %T,\n %T,\n %T,\n %T,\n %T\n", a, b, c, d, e, f)

	/*----------------------- STEP 2 -----------------------*/

	// Get the 2 values of return of the function http.Get()
	res, err := http.Get("https://google.com.br")

	// If the return of the erro its diferent of the null (nil)
	if err != nil {
		panic("ERRROOO")
	}

	// Variable "res" it's a propertie of the function Get of the Type Response and have any properties (like Body)
	fmt.Println(res.Body)

	/*----------------------- STEP 3 -----------------------*/

	// Print the address of the variable in the Memory
	x := 10
	fmt.Println(&x)

	// Print the value of the address used in the Memory
	y := &x
	fmt.Println(*y)

	/*----------------------- STEP 4 -----------------------*/

	// This variable its a Pointer (like the variable y, but explicitly declared)
	var z *int = &x
	fmt.Println(*z)
	response := abc(&x)
	fmt.Println(response)

	/*----------------------- STEP 5 -----------------------*/

	// You can realize negotiations of the error and variables
	g, h, err := Nome("wesley", 5) // Nome("wesley", 11)
	if err != nil {
		panic("Panico")
	}
	fmt.Println(g, h)

	/*----------------------- STEP 6 -----------------------*/

	// Array in the Go
	arrayVar := [10]int{4, 6, 8, 10, 12} // var arrayVar [10]int
	fmt.Println(arrayVar)
	fmt.Println(arrayVar[0])

	// Slice its a Array without limits, with infinite slot for values
	slice := make([]int, 5)
	slice[0] = 1
	slice[1] = 1
	slice[2] = 1
	slice[3] = 1
	slice[4] = 1
	fmt.Println(slice)

	// You add new slots in the Slice Type variable, different of the Array
	slice = append(slice, 1, 2, 3, 4)
	fmt.Println(slice)

	// This is a String Slice variable
	sliceString := []string{
		"Hello",
		"World",
	}
	fmt.Println(sliceString[0])

	/*----------------------- STEP 7 -----------------------*/

	// Map is a Array with Key and Values (like Dictionary)
	mapVar := make(map[string]int)
	mapVar["Wesley"] = 10
	mapVar["Maeda"] = 22
	fmt.Println(mapVar["Maeda"])

	/*----------------------- STEP 8 -----------------------*/

	// You can work with of the Types
	car1 := Car{
		CarName: "Fusion",
		CarYear: 2020,
	}
	fmt.Println(car1.CarName)

	car1.drive()

	/*----------------------- STEP 9 -----------------------*/

	// Case you want to send a Json to you API
	car2 := Car{
		CarName: "Gol",
		CarYear: 2000,
	}

	// In GO we have Tags (are marcations of the Objects)
	result, _ := json.Marshal(car2)
	fmt.Println(string(result))

	// We can use a Json struct to format and print your variables (separately)
	jason := []byte(`{"car":"BMW","year":"2020"}`)
	var car Car
	json.Unmarshal(jason, &car)
	fmt.Println(car.CarName)

	/*----------------------- STEP 10 -----------------------*/
	car3 := Car{
		CarName: "BMW",
		CarYear: 2010,
	}
	exemplo1(car3)

	/*----------------------- STEP 11 -----------------------*/
	exemplo.PrintExemplo()
	casa := exemplo.Casa{
		Cor:    "branco",
		Numero: 532,
	}
	fmt.Println(casa)
}

/*----------------------- STEP 10 -----------------------*/
// Function of the
func exemplo1(car vehicle) {
}

// Verify if the Function has the method start to make a class "vehicle"
func (c Car) start() string {
	return "Startou"
}

/*----------------------- STEP 4 -----------------------*/

// Function abc receives a Pointer and return a Int value
// fun Name of the Funcion(type of the parameter) return type { code }
func abc(a *int) int {
	return *a
}

/*----------------------- STEP 5 -----------------------*/

// In this Function we receive 2 variables and return 3 types of variables
func Nome(a string, b int) (string, int, error) {

	if b > 10 {
		return "", 0, errors.New("Deu ruim")
	}

	return a, b, nil
}

// func PrintExemplo() {
// 	exemplo.PrintExemplo()
// }
