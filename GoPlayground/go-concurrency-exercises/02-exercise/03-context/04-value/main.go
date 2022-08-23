package main

import (
	"context"
	"fmt"
)

type database map[string]bool
type userIDKeyType string

var db database = database{
	"jane": true,
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	processRequest(ctx, "jane")
}

func processRequest(ctx context.Context, userid string) {
	// TODO: send userID information to checkMemberShip through context for
	// map lookup.
	vctx := context.WithValue(ctx, userIDKeyType("userIDKey"), "jane")
	ch := checkMemberShip(vctx)
	status := <-ch
	fmt.Printf("membership status of userid : %s : %v\n", userid, status)
}

// checkMemberShip - takes context as input.
// extracts the user id information from context.
// spins a goroutine to do map lookup
// sends the result on the returned channel.
func checkMemberShip(vctx context.Context) <-chan bool {
	ch := make(chan bool)
	go func() {
		defer close(ch)
		// do some database lookup
		userID := vctx.Value(userIDKeyType("userIDKey")).(string)
		status := db[userID]
		ch <- status
	}()
	return ch
}
