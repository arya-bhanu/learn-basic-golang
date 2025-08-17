package main

import (
	"fmt"
)

func main() {
	var hero1 = Hero{Magic: 10, Anima: Anima{HP: 100, PowerAttack: 20, PowerDefense: 1}}
	var enemy1 = Enemy{Poition: 10, Anima: Anima{HP: 100, PowerAttack: 5, PowerDefense: 10}, IsWalkAway: false}
	var actHero1 Action = &hero1
	var actEnemy2 Action = &enemy1
	Attack(actHero1, actEnemy2)
}

type Anima struct {
	HP           int
	PowerAttack  int
	PowerDefense int
}

type Hero struct {
	Anima
	Magic int
}

type Enemy struct {
	Anima
	Poition    int
	IsWalkAway bool
}

type Action interface {
	Attack(defender *Anima)
	Defend(attacker *Anima)
}

type Move interface {
	WalkAway()
}

func EnemyWalkingAway(m Move, h *Hero) {
	m.WalkAway()
	h.HP = 0
}

func Attack(heroAct Action, enemyAct Action) {
	maxAttack := max(heroAct.(*Hero).PowerAttack, heroAct.(*Hero).Anima.PowerAttack)
	maxDefense := max(enemyAct.(*Enemy).PowerDefense, enemyAct.(*Enemy).Anima.PowerDefense)
	var enemy = enemyAct.(*Enemy)
	if enemy.HP > 0 {
		enemy.HP -= maxAttack - maxDefense
		if enemy.HP < 0 {
			enemy.HP = 0
		}
	}
	heroAct.Attack(&enemy.Anima)
}
func Defend(heroAct Action, enemyAct Action) {

}

func (h *Hero) Attack(defender *Anima) {
	fmt.Println("Hero attacks!")
	fmt.Printf("Defender/Enemy HP: %d\n", defender.HP)
}

func (h *Hero) Defend(attacker *Anima) {
	fmt.Println("Hero defends!")
}

func (e *Enemy) Attack(defender *Anima) {
	fmt.Println("Enemy attacks!")
}

func (e *Enemy) Defend(attacker *Anima) {
	fmt.Println("Enemy defends!")
}

func (e *Enemy) WalkAway() {
	e.IsWalkAway = true
}
