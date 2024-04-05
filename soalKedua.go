package main

import "fmt"

type BMI struct {
	weight int
	height int
}

func calculateBMI(weight float64, height float64) float64 {
	return weight / (height * height)
}

func main() {
	markWeight := 95.0
	markHeight := 1.88

	johnWeight := 85.0
	johnHeight := 1.76

	fmt.Printf("Tinggi badan Mark: %.2f m, Berat Mark: %.2f kg\n", markHeight, markWeight)
	fmt.Printf("Tinggi badan John: %.2f m, Berat John: %.2f kg\n", johnHeight, johnWeight)

	markBMI := calculateBMI(float64(markWeight), markHeight)
	johnBMI := calculateBMI(float64(johnWeight), johnHeight)

	fmt.Printf("BMI John: %.2f\n", johnBMI)
	fmt.Printf("BMI Mark: %.2f\n", markBMI)

	markHigherBMI := markBMI > johnBMI

	fmt.Println("Mark memiliki BMI yang lebih tinggi dari John:", markHigherBMI)
}
