package main

import "fmt"

func main() {
	var name = "Go confrance"
	const confranceTickets = 50
	var remainingTicket = 50

	fmt.Printf("Welcome to %v booking application\n", name)
	fmt.Printf("We have total of %v tickects and %v are still avialable\n", confranceTickets, remainingTicket)
	fmt.Println("Get your tickets here to attend")
}
