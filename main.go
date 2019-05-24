package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Hotel Automation. Please enter details:")

	h := Handler{}

	for {
		fmt.Println("Enter No. of Floors:")
		fmt.Scanf("%d", &h.InputHandler.Floor)
		fmt.Println("Enter No. of Main Corridor per Floor:")
		fmt.Scanf("%d", &h.InputHandler.MainCorridor)
		fmt.Println("Enter No. of Sub Corridor per Floor:")
		fmt.Scanf("%d", &h.InputHandler.SubCorridor)

		if h.InputHandler.Floor > 0 && h.InputHandler.MainCorridor > 0 && h.InputHandler.SubCorridor > 1 {
			break
		}

		fmt.Println("Please Enter correct number. No. of Floor and Main Corridor > 0 and Subcorridor > 1")
	}

	h.BuildHotelStructure()
	h.PowerHotelStructure()

	time.Sleep(time.Second * 1)
	h.Movements()

}
