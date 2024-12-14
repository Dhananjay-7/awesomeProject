package main

import "fmt"

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	//averageOpenRate, average := 23, 45
	//fmt.Println(averageOpenRate, "deifning domethidfjsdjf", average)
	x := 5
	x = increment(x)
	fmt.Println(x)
	fmt.Println(yearutilreturn(45))

}
func increment(x int) (y int) {

	x++
	return
}

func yearutilreturn(age int) (yearsUntilAdult int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 4
	}
	return

}
