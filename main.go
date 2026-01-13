package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var SkillPool = []Skill{
	{Name: "Slash", Type: "Physical", Damage: 10},
	{Name: "Fireball", Type: "Magic", Damage: 20},
	{Name: "Ice Spike", Type: "Magic", Damage: 15},
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Key in the Name of your player> ") // prompt

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	input = strings.TrimSpace(input)
	if input == "exit" {
		fmt.Println("Goodbye!")
		return
	}

	player1 := generatePlayer(input)
	player2 := generatePlayer("Player 2")
	player1.Display()
	fmt.Println("----------------------")
	player2.Display()

	reader = bufio.NewReader(os.Stdin)

	for {
		displayBoard()
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		switch input {
		case "attack":
			player2.TakeDamage(10)
		case "skills":
			player1.DisplaySkills()
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			index, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Invalid input")
				continue
			}
			player1.UseSkill(SkillPool[index-1], &player2)
		case "exit":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Unknown command")
		}

		player1.TakeDamage(10)
		summary(player1, player2)
		checkResult(player1, player2)
	}

}

func generatePlayer(input string) Player {

	return Player{
		Name:   input,
		Health: 100,
		Skills: generateSkills(),
	}
}

func generateSkills() []Skill {
	i := rand.Intn(len(SkillPool))
	j := rand.Intn(len(SkillPool))

	return []Skill{
		SkillPool[i],
		SkillPool[j],
	}
}

func displayBoard() {
	fmt.Println()
	fmt.Println("Choose an action:")
	fmt.Println("  attack  - Perform a basic attack")
	fmt.Println("  skills  - View your skills")
	fmt.Println("  exit    - Quit the game")
	fmt.Print("> ")
}

func summary(player1 Player, player2 Player) {
	fmt.Println("----------------------")
	fmt.Printf("%s's health: %d\n", player1.Name, player1.Health)
	fmt.Printf("%s's health: %d\n", player2.Name, player2.Health)
}

func checkResult(player1 Player, player2 Player) {
	if player1.Health <= 0 {
		fmt.Printf("%s has been defeated!\n", player2.Name)
	} else if player2.Health <= 0 {
		fmt.Printf("%s has won!\n", player1.Name)
	}
}
