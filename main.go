package main

import "log"

func main() {

	myVar := "horse"

	switch myVar {
	case "cat":
		log.Println("cat is set to ", myVar)
	case "dog":
		log.Println("cat is set to ", myVar)
	case "fish":
		log.Println("cat is set to ", myVar)
	default:
		log.Println("cat is something else  ")
	}

	// myNum :=100
	// isTrue := false

	// if myNum > 99 && !isTrue {
	// 	log.Println("myNum is greater than 99 and isTrue is set to true")

	// }
	// cat := "cat2"
	// if cat == "cat" {
	// 	log.Println("Cat is cat")

	// } else {
	// 	log.Println("Cat is not cat")
	// }
	// var isTrue bool

	// isTrue = false

	// if isTrue {
	// 	log.Println("isTrue is", isTrue)

	// } else {
	// 	log.Println("isTrue is", isTrue)
	// }
	// names := []string{"one", "fish", "cat"}
	// log.Println(names)

	// var mySlice []string
	// mySlice = append(mySlice, "James")
	// mySlice = append(mySlice, "Bond")
	// mySlice = append(mySlice, "Penny")

	// log.Println(mySlice)

	// numbers := []int{1, 2, 3, 4, 5}

	// log.Println(numbers)
	// log.Println(numbers[0:2])
	// myMap := make(map[string]string)

	// myMap["dog"]="Simpson"
	// myMap["other-dog"]="Cassie"

	// log.Println(myMap["dog"])
	// log.Println(myMap["other-dog"])
	// var myString string
	// var myInt int

	// myString = "Hi"
	// myInt = 11

	// mySecondString := "another string"

	// log.Println(myString, mySecondString, myInt)
}
