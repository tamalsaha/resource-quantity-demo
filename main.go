package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/api/resource"
)

func main() {
	q1 := resource.MustParse("100m")
	fmt.Println(q1, "|", q1.ToUnstructured(), q1.MilliValue())
	q1_copy := resource.Quantity{Format: resource.DecimalSI}
	q1_copy.SetMilli(3 * q1.MilliValue())
	fmt.Println(q1_copy, "|", q1_copy.ToUnstructured(), q1_copy.MilliValue())
	q1_copy.Format = resource.BinarySI
	fmt.Println(q1_copy, "|", q1_copy.ToUnstructured(), q1_copy.MilliValue())

	q2 := resource.MustParse("1.5Gi")
	fmt.Println(q2, "|", q2.ToUnstructured(), q2.Value(), q2.MilliValue())

	q3 := resource.Quantity{Format: resource.BinarySI}
	fmt.Println(q3, "|", q3.ToUnstructured(), "|", q3.IsZero())
}
