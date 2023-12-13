package main

import "log"

func main() {

type User struct {
	FirstName string
	LastName string
	Email string
	Age int
}

var users []User
users = append(users,User{"John","smith","john@smith.com",30})
users = append(users,User{"James","Bond","james@bond.com",40})
users = append(users,User{"Penny","Money","Penny@money.com",50})

for _,l:= range users{
	log.Println(l.FirstName,l.LastName,l.Email,l.Age)
}
// var firstLine = "Once upon a midnight dreary"

// for i,l:= range firstLine{
// 	log.Println(i,":",l)
// }

// animals := make(map[string]string)
// animals["dog"] ="Fido"
// animals["cat"] ="Fluffy"

// for animalType,animal:= range animals{
// 	log.Println(animalType,animal)
// }	
// animals := []string{"one", "fish", "cat"}

// for _,animal:= range animals{
// 	log.Println(animal)
// }

	// myVar := "horse"

	// switch myVar {
	// case "cat":
	// 	log.Println("cat is set to ", myVar)
	// case "dog":
	// 	log.Println("cat is set to ", myVar)
	// case "fish":
	// 	log.Println("cat is set to ", myVar)
	// default:
	// 	log.Println("cat is something else  ")
	// }

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
