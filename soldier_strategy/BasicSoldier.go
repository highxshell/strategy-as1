package soldier_strategy

import (
	"fmt"
	"sync"
)

var once sync.Once

type basicSoldier struct {
}

func (b basicSoldier) Info() {
	fmt.Println("I'm a basic Soldier")
}

var singleInstance *basicSoldier

func GetBasicSoldier() *basicSoldier {
	if singleInstance == nil {
		once.Do(
			func() {
				fmt.Println("No instances created, creating one")
				singleInstance = &basicSoldier{}
			})
	} else {
		fmt.Println("Instance was already created, don't ruin singleton")
	}
	return singleInstance
}
