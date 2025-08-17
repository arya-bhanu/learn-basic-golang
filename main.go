package main

import (
	"fmt"
)

func main() {
	chicken := map[string]int{
		"egg": 2,
		"sex": 1,
	}
	fmt.Printf("Chicken: %+v\n", chicken)

	chickens := []map[string]int{}
	chickens = append(chickens, chicken)
	chickens = append(chickens, map[string]int{
		"egg": 10,
		"sex": 0,
	})

	fmt.Printf("Chickens: %+v\n", chickens)
	calc := calculate(1, 2, 3, 4, 10)
	fmt.Printf("Calc: %+v\n", calc)

	numbers := []int{2, 3, 10, 12, 45, 50}
	filteredNumber := filterNumber(numbers, func(i int) bool {
		return i > 50
	})
	fmt.Println(filteredNumber)

	var val *int = new(int)
	*val = 42
	fmt.Println(*val)

	myNum := 50
	fmt.Println(myNum)
	changeMyNumber(&myNum, 100)
	fmt.Println(myNum)

	var anima1 Anima
	fmt.Printf("Anima 1: %+v\n", anima1)

	var anima2 = Anima{
		Name: "Jordan",
		HP:   100,
	}
	fmt.Printf("Anima 2: %+v\n", anima2)

	var anima3 = Anima{
		Name: "Anima 3",
		HP:   100,
	}

	var anima4 *Anima = &anima3

	anima4.Name = "Changed Anima 3"

	fmt.Printf("Anima 3: %+v\n", anima3)

	var hero1 = Hero{
		Anima: Anima{
			Name:         "Hero 1",
			HP:           100,
			DefensePower: 10,
		},
		AttackPower: 200,
	}

	var defense1 = WallDefense{
		Anima: Anima{
			Name:        "Wall Defense 1",
			HP:          300,
			AttackPower: 10,
		},
		DefensePower: 100,
	}

	fmt.Println("========= GAME START =======")
	fmt.Printf("Hero: %+v\n", hero1)
	fmt.Printf("Defense: %+v\n", defense1)
	hero1.HeroAttack(&defense1.Anima, defense1.DefensePower)
	fmt.Printf("Defense: %+v", defense1)
}

type Anima struct {
	Name         string
	HP           int
	AttackPower  int
	DefensePower int
}

type Hero struct {
	Anima
	AttackPower int
}

type WallDefense struct {
	Anima
	DefensePower int
}

func calculate(numbers ...int) int {
	counter := 0
	for _, val := range numbers {
		counter += val
	}
	return counter
}

func filterNumber(numbers []int, callback func(int) bool) []int {
	filtered := []int{}
	for _, i := range numbers {
		if callback(i) {
			filtered = append(filtered, i)
		}
	}
	return filtered
}

func changeMyNumber(oldNum *int, newNum int) {
	*oldNum = newNum
}

func (h *Hero) HeroAttack(enemy *Anima, parentDefense int) {
	maxPowerAttack := max(h.AttackPower, h.Anima.AttackPower)
	maxDefense := max(enemy.DefensePower, parentDefense)
	fmt.Println(maxDefense)
	if enemy.HP > 0 {
		enemy.HP -= maxPowerAttack - maxDefense
		if enemy.HP < 0 {
			enemy.HP = 0
		}
	}
}

func (wd *WallDefense) DefensingFromWall(enemy *Anima, parentAttack int) {

}
