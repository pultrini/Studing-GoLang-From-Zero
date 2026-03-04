package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go callDatabase(&wg)
	go callAPI(&wg)
	go processosInternos(&wg)
	wg.Wait()
}

func callDatabase(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("finalizado database")
	wg.Done()
}

func callAPI(wg *sync.WaitGroup) {
	time.Sleep(2 * time.Second)
	fmt.Println("finalizado api")
	wg.Done()

}
func processosInternos(wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("finalizado interno")
	wg.Done()
}
