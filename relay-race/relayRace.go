package main

import (
	"fmt"
	"sync"
	"time"
)

const NUMBER_TEAMS int = 5
const NUMBER_RUNNERS int = 4

var waitGroup sync.WaitGroup

func main() {
	racingLanes := make([]chan int, NUMBER_TEAMS)
	for i := 0; i < NUMBER_TEAMS; i++ {
		racingLanes[i] = make(chan int)
	}
	//The race will only end when all teams finish the full course
	waitGroup.Add(NUMBER_TEAMS)

	for i := 0; i < NUMBER_TEAMS; i++ {
		nameTeam := fmt.Sprint("Equipe ", (i + 1))
		go Runner(nameTeam, racingLanes[i]) //Put the fist runners of each team in his position
		racingLanes[i] <- 1                 //Send the lane to first runner of each lane and start the run
	}

	waitGroup.Wait()
	fmt.Printf("\n ============= \nCorrida finalizada")
}

func Runner(team string, lane chan int) {
	runner := <-lane

	fmt.Printf("[%s] - Corredor %d correndo com o bastão\n", team, runner)

	// Running in the lane
	time.Sleep(4 * 100 * time.Millisecond)

	if runner == NUMBER_RUNNERS {
		fmt.Printf("[%s] - Finalizou a corrida! \n", team)
		waitGroup.Done()
		return
	}

	nextRunner := runner + 1
	//Puting the next runner in his position
	go Runner(team, lane)

	fmt.Printf("[%s] - Corredou %d passou o bastão para o corredor %d\n", team, runner, nextRunner)

	lane <- nextRunner
}
