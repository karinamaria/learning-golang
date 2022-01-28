package main

import (
	"fmt"
	"sync"
	"github.com/karinamaria/learning-golang/relay-race/src/race"
)

const NUMBER_TEAMS int = 5

func main() {
	var racingLanes []chan int
	for i := 0; i < NUMBER_TEAMS; i++ {
		racingLanes = append(racingLanes, make(chan int))
	}

	//The race will only end when all teams finish the full course
	var wg sync.WaitGroup
	wg.Add(NUMBER_TEAMS)

	arrivalOrder := make(chan string, NUMBER_TEAMS) //Buffered Channel to store arrival order
	for i := 0; i < NUMBER_TEAMS; i++ {
		nameTeam := fmt.Sprint("Equipe ", (i + 1))
		go race.Runner(&wg, nameTeam, racingLanes[i], arrivalOrder) //Put the fist runners of each team in his position
	}

	for i := 0; i < NUMBER_TEAMS; i++ {
		racingLanes[i] <- 1 //Send the baton to first runner of each lane and start the run
	}
	wg.Wait()

	fmt.Println("\n =============\nCorrida finalizada")
	for i := 0; i < NUMBER_TEAMS; i++ {
		fmt.Printf("%dº Lugar: %s\n", (i + 1), <-arrivalOrder)
	}
}