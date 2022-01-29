package race

import (
	"fmt"
	"sync"
	"time"
)

const NUMBER_RUNNERS int = 4

func Runner(waitGroup *sync.WaitGroup, nameTeam string, baton chan int, arrivalOrder chan string) {
	runner := <-baton

	fmt.Printf("[%s] - Corredor %d correndo com o bastão\n", nameTeam, runner)
	nextRunner := runner + 1
	if nextRunner <= NUMBER_RUNNERS{
		go Runner(waitGroup, nameTeam, baton, arrivalOrder) //Puting the next runner in his position
	}
	time.Sleep(time.Microsecond * time.Duration(4 * 100)) // Running with the baton

	if runner == NUMBER_RUNNERS {
		fmt.Printf("[%s] - Finalizou a corrida! \n", nameTeam)
		arrivalOrder <- nameTeam
		waitGroup.Done()
		return
	}

	fmt.Printf("[%s] - Corredou %d passou o bastão para o corredor %d\n", nameTeam, runner, nextRunner)
	baton <- nextRunner
}