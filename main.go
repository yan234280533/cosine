package main

import "fmt"

func main() {
	ret := Cosine([]float64{1.0}, []float64{1.0})
	fmt.Printf("%.2f\n", ret)

	ret = Cosine([]float64{1.0, 0.0}, []float64{1.0, 0.0})
	fmt.Printf("%.2f\n", ret)

	ret = Cosine([]float64{0.0, 1.0}, []float64{1.0, 0.0})
	fmt.Printf("%.2f\n", ret)

	ret = Cosine([]float64{0.5, 0.5}, []float64{1.0, 0.0})
	fmt.Printf("%.2f\n", ret)

	ret = Cosine([]float64{0.2, 0.8}, []float64{1.0, 0.0})
	fmt.Printf("%.2f\n", ret)
}
