package main

import "fmt"

type deck []string

func print(d deck) {
	for _, text := range d {
		fmt.Println(text)
	}

}
