// - Partindo do código abaixo, ordene os []user por idade e sobrenome.
//     - https://play.golang.org/p/BVRZTdlUZ_
// - Os valores no campo Sayings devem ser ordenados tambem, e demonstrados de maneira esteticamente harmoniosa.
// - Solução: https://play.golang.org/p/3wgW4BDasu

package main

import (
	"fmt"
	"sort"
)

type user struct {
	First string
	Last  string
	Age   int
}
type orderByAge []user

func (x orderByAge) Len() int { return len(x) }

func (x orderByAge) Less(i, j int) bool {
	return x[i].Age < x[j].Age
}

func (x orderByAge) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	sort.Sort(orderByAge(users))

	fmt.Println(users)
	// your code goes here

}
