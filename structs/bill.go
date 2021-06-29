package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 4.33, "cake": 3.53},
		tip:   0,
	}

	return b
}

// format the bill
func (b bill) format() string {
	fs := "Bill \n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%v ...$%v \n", k+":", v)
		total += v
	}

	// total
	fs += fmt.Sprintf("%v ...%0.2f", "total", total)
	return fs
}

// update the tip
func (b *bill) updateTip(newTip float64) {
	// go automatically dereferences the pointer with structs, maps and slices
	b.tip = newTip
}
