package main

import "fmt"

type Team struct {
	Name  string
	Score []int
}

func (team *Team) calculateAverage() float64 {
	var total int
	for _, score := range team.Score {
		total += score
	}
	return float64(total) / float64(len(team.Score))
}

func main() {
	teams := []Team{
		{"Lumba-lumba", []int{97, 112, 101}},
		{"Koala", []int{109, 95, 106}},
	}

	var highestAverage float64
	var winningTeam string

	for _, team := range teams {
		average := team.calculateAverage()
		fmt.Printf("Rata-rata skor tim %s: %.2f\n", team.Name, average)
		if average > highestAverage {
			highestAverage = average
			winningTeam = team.Name
		}
	}

	var countWinningTeams int
	for _, team := range teams {
		if team.calculateAverage() >= 100 && team.calculateAverage() == highestAverage {
			countWinningTeams++
		}
	}

	if countWinningTeams == 2 {
		fmt.Println("Hasil seri!")
	} else if countWinningTeams == 1 {
		fmt.Printf("Tim %s adalah pemenang!\n", winningTeam)
	} else {
		fmt.Println("Tidak ada tim yang memenangkan trofi.")
	}
}
