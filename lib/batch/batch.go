package batch

import (
	"fmt"
	"runtime"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	jobs := make(chan int64, n)
	results := make(chan user)

	for w := 0; w < int(pool); w++ {
		go worker(jobs, results)
	}

	for j := 0; j < int(n); j++ {
		jobs <- int64(j)
	}

	for a := 0; a < int(n); a++ {
		res = append(res, <-results)
	}

	fmt.Println("NumGoroutine after:", runtime.NumGoroutine())
	return res
}

func worker(jobs <-chan int64, results chan<- user) {
	for j := range jobs {
		results <- getOne(j)
	}
}
