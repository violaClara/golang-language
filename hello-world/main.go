package main

import (
	"fmt"
	"math"


)

func main() {
	// fmt.Println("Hello world")

	// var firstName string = "John"
	// _ = "wick"

	// fmt.Println("halo", firstName)

	// name := new(string)
	// fmt.Println(name)
	// fmt.Println(*name)

	// var decimal = 2.64
	// fmt.Printf("Bilangan desimal: %f", decimal)

	// var point = 6
	// switch {
	// case point == 8:
	// 	fmt.Println("perfect")
	// case (point < 8) && (point > 3):
	// 	fmt.Println("awesome")
	// 	fallthrough
	// case point < 5:
	// 	fmt.Println("you need to learn more")
	// default:
	// 	{
	// 		fmt.Println("not bad")
	// 		fmt.Println("you need to learn more")
	// 	}
	// }

	// var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// fmt.Println("jumlah element \t\t", len(fruits))
	// fmt.Println("isi semua element \t", fruits)

	// var numbers = [...]int{2, 3, 4, 2, 3, 4, 3}
	// fmt.Println("jumlah array number: \t", len(numbers))

	// numbers2 := [2][3]int{{3, 2, 3}, {2, 4, 5}}
	// fmt.Println("numbers matrik multi dimension: \t", numbers2)

	// var fruitsi = make([]int, 3)
	// fruitsi[0] = 9
	// fruitsi[1] = 39

	// fmt.Println(fruitsi)

	// var newFruits = fruits[0:2]
	// fmt.Println(newFruits)

	// var newFruits = []string{"apple", "grape", "banana", "melon"}

	// var afruits = newFruits[0:3]
	// var bfruits = newFruits[1:4]

	// var aafruits = afruits[1:2]
	// var bafruits = bfruits[0:1]

	// fmt.Println(newFruits)
	// fmt.Println(afruits)
	// fmt.Println(bfruits)
	// fmt.Println(aafruits)
	// fmt.Println(bafruits)

	// bafruits[0] = "pinnaple"

	// fmt.Println(newFruits)
	// fmt.Println(afruits)
	// fmt.Println(bfruits)
	// fmt.Println(aafruits)
	// fmt.Println(bafruits)

	var diameter float64 = 15
	var area, circumference = calculated(diameter)

	fmt.Printf("luas linkaran\t\t: %.2f \n", area)
	fmt.Printf("Keliling lingkaran\t: %.2f \n", circumference)

	var avg=calculate(2,3,4,5,3,2,3,4,4,4)
	var msg= fmt.Sprintf("Rata0rata: %.2f",avg)
	fmt.Println(msg)

}

// func printMessage(message string, arr[]string)  {
// 	var nameString=strings.join(arr," ")
// 	fmt.Println(message, nameString)
// }

func calculated(d float64) (float64, float64) {
	var area = math.Pi * math.Pow(d/2, 2)
	var circumference = math.Pi * d

	return area, circumference
}

func calculate(numbers ...int) float64{
	var total int=0
	for_,number:=range numbers{
		total +=number
	}
	var avg= float64(total)/float64(len(numbers))
	return avg
}
