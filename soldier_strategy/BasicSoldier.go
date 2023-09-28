package soldier_strategy

import "fmt"

type BasicSoldier struct {
}

func (b BasicSoldier) Info() {
	fmt.Println("I'm a basic Soldier")
}
