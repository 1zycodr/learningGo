package main

import (
	"fmt"
	"sync"
)

type Income struct {
	Title  string
	Amount int
}

func main() {
	var totalAmount int
	var wg sync.WaitGroup
	var mutex sync.Mutex

	incomes := []Income{
		{
			Title:  "Job 1",
			Amount: 10,
		},
		{
			Title:  "Job 2",
			Amount: 500,
		},
		{
			Title:  "Job 3",
			Amount: 50,
		},
		{
			Title:  "Job 4",
			Amount: 100,
		},
	}

	for i := 0; i < 52; i++ {
		for _, income := range incomes {
			wg.Add(1)
			go func(amount int, wg *sync.WaitGroup, m *sync.Mutex) {
				defer wg.Done()
				m.Lock()
				totalAmount += amount
				m.Unlock()
			}(income.Amount, &wg, &mutex)
		}
	}
	wg.Wait()

	fmt.Println(totalAmount)
}
