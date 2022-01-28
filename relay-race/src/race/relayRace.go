package race

import (
	"fmt"
	"sync"
	"time"
)

const NUMBER_RUNNERS int = 4

func Runner(wg *sync.WaitGroup, team string, baton chan int, arrivalOrder chan string) {
	runner := <-baton

	fmt.Printf("[%s] - Corredor %d correndo com o bastão\n", team, runner)

	// Running with the baton
	time.Sleep(4 * 100 * time.Microsecond)

	if runner == NUMBER_RUNNERS {
		fmt.Printf("[%s] - Finalizou a corrida! \n", team)
		arrivalOrder <- team
		wg.Done()
		return
	}

	nextRunner := runner + 1
	//Puting the next runner in his position
	go Runner(wg, team, baton, arrivalOrder)

	fmt.Printf("[%s] - Corredou %d passou o bastão para o corredor %d\n", team, runner, nextRunner)

	baton <- nextRunner
}