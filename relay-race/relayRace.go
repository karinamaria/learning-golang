package main

import (
	"fmt"
	"sync"
	"time"
)

const NUMBER_TEAMS int = 5
const NUMBER_RUNNERS int = 4

var wg sync.WaitGroup

func main() {
	var racingLanes []chan int
	arrivalOrder := make(chan string, NUMBER_TEAMS) //Buffered Channel to store arrival order

	for i := 0; i < NUMBER_TEAMS; i++ {
		racingLanes = append(racingLanes, make(chan int))
	}
	//The race will only end when all teams finish the full course
	wg.Add(NUMBER_TEAMS)

	for i := 0; i < NUMBER_TEAMS; i++ {
		nameTeam := fmt.Sprint("Equipe ", (i + 1))
		go Runner(nameTeam, racingLanes[i], arrivalOrder) //Put the fist runners of each team in his position
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

func Runner(team string, baton chan int, arrivalOrder chan string) {
	runner := <-baton

	fmt.Printf("[%s] - Corredor %d correndo com o bastão\n", team, runner)

	// Running in the baton
	time.Sleep(4 * 100 * time.Millisecond)

	if runner == NUMBER_RUNNERS {
		fmt.Printf("[%s] - Finalizou a corrida! \n", team)
		arrivalOrder <- team
		wg.Done()
		return
	}

	nextRunner := runner + 1
	//Puting the next runner in his position
	go Runner(team, baton, arrivalOrder)

	fmt.Printf("[%s] - Corredou %d passou o bastão para o corredor %d\n", team, runner, nextRunner)

	baton <- nextRunner
}