package main

import (
	"fmt"
	"github.com/luisjn/abstract-factory/sports"
)

func main() {
	// smsFactory, _ := getNotificationFactory("SMS")
	// emailFactory, _ := getNotificationFactory("Email")
	//
	// sendNotification(smsFactory)
	// getMethod(smsFactory)
	// sendNotification(emailFactory)
	// getMethod(emailFactory)

	adidasFactory, _ := sports.GetSportsFactory("adidas")
	nikeFactory, _ := sports.GetSportsFactory("nike")

	nikeShoe := nikeFactory.MakeShoe()
	nikeShirt := nikeFactory.MakeShirt()

	adidasShoe := adidasFactory.MakeShoe()
	adidasShirt := adidasFactory.MakeShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}

func printShoeDetails(s sports.IShoe) {
	fmt.Println("Shoe")
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}

func printShirtDetails(s sports.IShirt) {
	fmt.Println("Shirt")
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}
