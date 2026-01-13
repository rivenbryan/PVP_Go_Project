package main

import (
	"fmt"
)

type Skill struct {
	Name   string
	Type   string
	Damage int
}
type Fighter interface {
	TakeDamage(damage int)
	IsAlive() bool
	GetHealth() int
}
type SkillUser interface {
	GetSkills() []Skill
	UseSkill(skill Skill, target Fighter)
}

type Player struct {
	Name   string
	Health int
	Skills []Skill
}

func (p *Player) TakeDamage(damage int) {
	p.Health -= damage
}

func (p *Player) IsAlive() bool {
	return p.Health > 0
}

func (p *Player) GetHealth() int {
	return p.Health
}

func (p *Player) GetSkills() []Skill {
	return p.Skills
}

func (p *Player) UseSkill(skill Skill, target Fighter) {
	fmt.Printf("%s uses %s on %d\n", p.Name, skill.Name, target.GetHealth())
	target.TakeDamage(skill.Damage)
}

func (p *Player) Display() {
	fmt.Printf("Player: %s, Health: %d\n", p.Name, p.Health)
	p.DisplaySkills()

}

func (p *Player) DisplaySkills() {
	fmt.Println("Skills:")
	for i, skill := range p.Skills {
		fmt.Printf("%d, %s - %s\n", i+1, skill.Name, skill.Type)
	}
}
