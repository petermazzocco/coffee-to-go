package mutexHandler

import (
	"log"
	"strconv"
	"sync"
	"time"
)

func baristaUsingMachine(m *sync.Mutex, wg *sync.WaitGroup, c chan string, i int) {
	defer wg.Done()
	m.Lock()         // The barista is using the machine and we're locking it down
	defer m.Unlock() // Once the barista is done, we unlock the machine for the next barista
	// simulate working time on machine
	time.Sleep(3 * time.Second)
	c <- "Barista " + strconv.Itoa(i) + " has used the coffee machine."
}

func MachineInUse() {
	machineInUse := make(chan string) // channel to handle the machine usage data
	var m sync.Mutex                  // prevent multiple goroutine from accessing the var at the same time
	var numBaristas = 2               // our baristas waiting to work
	var wg sync.WaitGroup             // wait group to wait for all goroutines to finish

	log.Println("Locking the machines for the baristas")
	for i := range numBaristas {
		wg.Add(1)
		go baristaUsingMachine(&m, &wg, machineInUse, i)
	}

	// Wait for the machine to be used and close the channel once it's done
	go func() {
		wg.Wait()
		close(machineInUse)
	}()

	// Wait for the machine to be used
	for msg := range machineInUse {
		log.Println(msg)
	}
}
