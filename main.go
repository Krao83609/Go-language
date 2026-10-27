package main

import "fmt"

func main() {
	var workshopName = "Golang Workshop"
	const workshopTickets = 120
	var remainingTickets = 40
	var userName string
	var userTickets int

	wishMeLuck := "Nope, you are done!"

	fmt.Println(wishMeLuck)

	fmt.Println("Welcome to", workshopName)
	fmt.Println("We have Total tickets", workshopTickets, "And we have Left out tickets", remainingTickets)
	fmt.Println("This is three day", workshopName)

	userName = "Shae"
	userTickets = 3

	fmt.Printf(" User %v reserved %v tickets for Tyrion and Tywin.\n", userName, userTickets)

}
