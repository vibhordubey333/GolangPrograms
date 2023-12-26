package main

import (
	"fmt"
	"sync"
	"time"
)

// Ball represents the ping pong ball.
type Ball struct {
	hits int
	mu   sync.Mutex
}

// Player represents a player in the ping-pong game.
type Player struct {
	name string
	ball *Ball
}

// Hit simulates a player hitting the ball.
func (p *Player) Hit() {
	p.ball.mu.Lock()
	defer p.ball.mu.Unlock()
	p.ball.hits++
	fmt.Printf("%s hits the ball. Total hits: %d\n", p.name, p.ball.hits)
}

// Game simulates the ping-pong game.
func Game(player1, player2 *Player, maxHits int) {
	var wg sync.WaitGroup
	ball := &Ball{}

	player1.ball = ball
	player2.ball = ball

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < maxHits; i++ {
			time.Sleep(time.Millisecond * 100)
			player1.Hit()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < maxHits; i++ {
			time.Sleep(time.Millisecond * 100)
			player2.Hit()
		}
	}()

	wg.Wait()

	// Determine the winner based on the total hits
	var winner *Player
	if player1.ball.hits > player2.ball.hits {
		winner = player1
	} else {
		winner = player2
	}

	fmt.Printf("%s wins with %d hits!\n", winner.name, winner.ball.hits)
}

func main() {
	player1 := &Player{name: "Player 1"}
	player2 := &Player{name: "Player 2"}

	// Simulate the game with a maximum of 10 hits
	Game(player1, player2, 1)
}
