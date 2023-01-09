package main

import (
	"fmt"
	mon "github.com/kingandrewsolomon/datamon/monster"
)

func main() {
	d := mon.Datamon{Name: "asdf", Age: 0, Is_fainted: false, Times_fainted: 0, Health: 100, Hunger: 100, Strength: 2, Stamina: 0.5, Speed: 3}
	fmt.Println(d.Name)
	fmt.Print("Hello World!")
}
