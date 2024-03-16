package main

import (
	"fmt"
)

func isInt64AndFloat(t any) bool {
	switch t.(type) {
	case int64, float64:
		return true
	default:
		return false
	}
}

func mainSwitch() {
	int64Value := int64(42)
	float64Value := float64(3.14)
	stringValue := "Hello"
	sliceValue := []string{}

	fmt.Printf("isInt64AndFloat(%v): %v\n", int64Value, isInt64AndFloat(int64Value))
	fmt.Printf("isInt64AndFloat(%v): %v\n", float64Value, isInt64AndFloat(float64Value))
	fmt.Printf("isInt64AndFloat(%v): %v\n", stringValue, isInt64AndFloat(stringValue))
	fmt.Printf("isInt64AndFloat(%v): %v\n", sliceValue, isInt64AndFloat(sliceValue))

}
