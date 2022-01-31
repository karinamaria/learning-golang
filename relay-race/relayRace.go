package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	NUMBER_TEAMS   int = 1
	NUMBER_RUNNERS int = 4
	METERS         int = 100
)

func Runner(waitGroup *sync.WaitGroup, nameTeam string, baton chan int) {
	runner := <-baton

	fmt.Printf("[%s] - Corredor %d correndo com o bastão\n", nameTeam, runner)
	nextRunner := runner + 1
	if nextRunner <= NUMBER_RUNNERS {
		go Runner(waitGroup, nameTeam, baton) //Puting the next runner in his position
	}
	time.Sleep(time.Millisecond * time.Duration(METERS)) // Running with the baton

	if runner == NUMBER_RUNNERS {
		fmt.Printf("[%s] - Finalizou a corrida! \n", nameTeam)
		waitGroup.Done()
		return
	}

	fmt.Printf("[%s] - Corredou %d passou o bastão para o corredor %d\n", nameTeam, runner, nextRunner)
	baton <- nextRunner
}

func main() {
	racingLanes := make(chan int) //Channel to pass baton between the runners of each team

	//The race will only end when all teams finish the full course
	var waitGroup sync.WaitGroup
	waitGroup.Add(NUMBER_TEAMS)

	nameTeam := fmt.Sprint("Equipe ", NUMBER_TEAMS)
	go Runner(&waitGroup, nameTeam, racingLanes) //Put the first runners of each team in his position
	racingLanes <- 1                             //Send the baton to first runner of each lane and start the run

	waitGroup.Wait()

	fmt.Println("\n =============\nCorrida finalizada")
}