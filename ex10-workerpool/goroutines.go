package goroutines

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

var syn sync.WaitGroup
var stop sync.Mutex

func Workers(worker_id int, exercises <-chan string) {

	defer syn.Done()
	worker_status := false
	stop.Unlock()
	for chan_exercise := range exercises {

		if worker_status == false {
			fmt.Printf("worker:%d spawning\n", worker_id)
			worker_status = true
		}

		sleep_time, _ := time.ParseDuration(chan_exercise + "s")
		fmt.Printf("worker:%d sleep:%s\n", worker_id, chan_exercise)
		time.Sleep(sleep_time)
	}

	if worker_status == true {
		fmt.Printf("worker:%d stopping\n", worker_id)
	}
}

func Run(poolSize int) {

	read_Stdin := bufio.NewScanner(os.Stdin)
	exercises := make(chan string, 100)

	for read_Stdin.Scan() {
		input := string(read_Stdin.Bytes())
		exercises <- input
	}

	for worker_id := 1; worker_id <= poolSize; worker_id++ {
		syn.Add(1)
		stop.Lock()
		go Workers(worker_id, exercises)
	}

	close(exercises)
	syn.Wait()
}
