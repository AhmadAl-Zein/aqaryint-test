package main

import (
	"fmt"
)

type Seat struct {
	id      int
	student string
}

func swap(seats []Seat) []Seat {
	for i := 0; i < len(seats)-1; i += 2 {
		seats[i].id, seats[i+1].id = seats[i+1].id, seats[i].id
	}

	return seats
}

func main() {
	seats := []Seat{
		{1, "Ahmad"},
		{2, "Mohammed"},
		{3, "Samer"},
		{4, "Abd"},
		{5, "Amr"},
	}

	fmt.Println("Before:")
	for _, seat := range seats {
		fmt.Printf("%2d - %-8s\n", seat.id, seat.student)
	}

	swappedSeats := swap(seats)

	fmt.Println("\nAfter:")
	for _, seat := range swappedSeats {
		fmt.Printf("%2d - %-8s\n", seat.id, seat.student)
	}
}
