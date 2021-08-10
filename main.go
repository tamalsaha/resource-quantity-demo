package main

import (
	"fmt"
	"io/ioutil"
	"log"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"sigs.k8s.io/yaml"
)

func main() {
	var rr v1.ResourceList
	fmt.Println(rr.Cpu())

	var val interface{} = "abc"
	fmt.Println(val.(map[string]interface{}))
	for k, v := range val.(map[string]interface{}) {
		fmt.Println(k, v)
	}
}

func main_() {
	data, err := ioutil.ReadFile("/home/tamal/go/src/kmodules.xyz/resource-size/testdata/apps/v1/statefulset.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	var obj map[string]interface{}
	err = yaml.Unmarshal(data, &obj)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(obj)

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
