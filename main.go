package main

import (
	"fmt"

	"github.com/sagarciaescobar/mb-integrador-I/internal/tickets"
)

func main() {
	//	total, err := tickets.GetTotalTickets("Brazil")
	// data, _ := tickets.GetBeTimeTickes("manana")
	data, _ := tickets.AverageDestination("Brazil")
	fmt.Println(data)
}
