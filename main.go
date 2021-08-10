package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/api/resource"
)

func main() {
	q1 := resource.MustParse("100m")
	fmt.Println(q1, "|", q1.ToUnstructured())

	q2 := resource.MustParse("1Gi")
	fmt.Println(q2, "|", q2.ToUnstructured())

	q3 := resource.Quantity{Format: resource.BinarySI}
	fmt.Println(q3, "|", q3.ToUnstructured())

}
