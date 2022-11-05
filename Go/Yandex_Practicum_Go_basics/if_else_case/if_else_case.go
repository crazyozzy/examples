package main

import "fmt"

func PrintGeneration(user_age int) {
	switch {
	case (user_age >= 1946 && user_age <= 1964):
		fmt.Printf("Hi, Boomer!")
	case (user_age >= 1965 && user_age <= 1980):
		fmt.Printf("Hi, X-generation man!")
	case (user_age >= 1981 && user_age <= 1996):
		fmt.Printf("Hi, Millenial!")
	case (user_age >= 1997 && user_age <= 2012):
		fmt.Printf("Hi, Zuumer'ок!")
	case (user_age >= 2013):
		fmt.Printf("Hi, Alpha generation!")
	default:
		fmt.Printf("Hi, you are not belong to any generation!")
	}
}

func GetUserAge() (user_age int) {
	fmt.Printf("Hey, <user name>! Enter year of your birthday in four digit format (i.e. 1991, 1963 etc.): ")
	fmt.Scanf("%d", &user_age)
	return user_age
}

func main() {
	user_age := GetUserAge()
	PrintGeneration(user_age)
}