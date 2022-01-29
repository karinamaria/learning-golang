package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	NUMBER_TEAMS int = 5
	NUMBER_RUNNERS int = 4
)

func Runner(waitGroup *sync.WaitGroup, nameTeam string, baton chan int, arrivalOrder chan string) {
	runner := <-baton

	fmt.Printf("[%s] - Corredor %d correndo com o bastão\n", nameTeam, runner)

	// Running with the baton
	time.Sleep(time.Microsecond * time.Duration(4 * 100))

	if runner == NUMBER_RUNNERS {
		fmt.Printf("[%s] - Finalizou a corrida! \n", nameTeam)
		arrivalOrder <- nameTeam
		waitGroup.Done()
		return
	}

	nextRunner := runner + 1
	//Puting the next runner in his position
	go Runner(waitGroup, nameTeam, baton, arrivalOrder)
	fmt.Printf("[%s] - Corredou %d passou o bastão para o corredor %d\n", nameTeam, runner, nextRunner)
	baton <- nextRunner
}

func main() {
	var racingLanes []chan int
	for i := 0; i < NUMBER_TEAMS; i++ {
		racingLanes = append(racingLanes, make(chan int))
	}
	//The race will only end when all teams finish the full course
	var waitGroup sync.WaitGroup
	waitGroup.Add(NUMBER_TEAMS)

	arrivalOrder := make(chan string, NUMBER_TEAMS) //Buffered Channel to store arrival order
	for i := 0; i < NUMBER_TEAMS; i++ {
		nameTeam := fmt.Sprint("Equipe ", (i + 1))
		go Runner(&waitGroup, nameTeam, racingLanes[i], arrivalOrder) //Put the fist runners of each team in his position
	}

	for i := 0; i < NUMBER_TEAMS; i++ {
		racingLanes[i] <- 1 //Send the baton to first runner of each lane and start the run
	}
	waitGroup.Wait()

	fmt.Println("\n =============\nCorrida finalizada")
	for i := 0; i < NUMBER_TEAMS; i++ {
		fmt.Printf("%dº Lugar: %s\n", (i + 1), <-arrivalOrder)
	}
}