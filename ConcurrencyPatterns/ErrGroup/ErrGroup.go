package main

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

var (
	errFailure = errors.New("CustomError")
)

func main() {
	//Create errgroup
	group := errgroup.Group{}

	//Run first task.
	group.Go(func() error {
		time.Sleep(5 * time.Second)
		fmt.Println("Performing some task...I")
		return nil
	})

	//Run second task.
	group.Go(func() error {
		fmt.Println("Performing some task...II")
		return nil
	})

	//Run third task.
	group.Go(func() error {
		fmt.Println("Performing some task...III")
		//Note: Happy Path.
		return nil
		//Note: Try un-commenting below line to see failure.
		//return errFailure
	})

	//Wait for all goroutines to complete.
	if err := group.Wait(); err != nil {
		fmt.Printf("ErrGroup tasks ended up with an error: %v\n", err)
	} else {
		fmt.Println("All work done successfully.")
	}
}

/*
Positive Scenario:
Output:
Performing some task...III
Performing some task...II
Performing some task...I
All work done successfully.

Negative Scenario:
Output:
Performing some task...III
Performing some task...II
Performing some task...I
ErrGroup tasks ended up with an error: CustomError

*/
