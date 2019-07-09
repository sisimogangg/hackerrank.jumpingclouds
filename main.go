package main

import "fmt"

//

func main() {
	//clouds := [...]int{0, 1, 0, 0, 0, 1, 0}
	clouds := [...]int{0, 0, 1, 0, 0, 1, 0}
	size := int32(len(clouds))
	var step int32
	stepCount := 0

	//steps 2 1 1 2
	//
	for step = 0; step < size; step++ {
		isNextValueValid := step+1 < size
		isNextTwoValuesValid := step+2 < size

		if isNextValueValid && clouds[step+1] != 0 {
			if isNextTwoValuesValid && clouds[step+2] == 0 {
				if isNextTwoValuesValid {
					step = step + 2
					fmt.Println("step : ", step)
					stepCount++
				}
			}
		} else {
			if isNextTwoValuesValid && clouds[step+2] == 0 {
				step = step + 2
				fmt.Println("step : ", step)
				stepCount++
			} else if isNextValueValid {
				step = step + 1
				fmt.Println("step : ", step)
				stepCount++
			}
		}
	}

	fmt.Println("Steps count: ", stepCount)
}
