package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)
	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["MN"] = "Man"
	fmt.Println(states)

	states["CA"] = "California"
	california := states["CA"]
	fmt.Println(california)

	delete(states, "MN")
	fmt.Println(states)

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println("Sorted")

	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
