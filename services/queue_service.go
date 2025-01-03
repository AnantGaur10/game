package services

import (
	"fmt"
	"game/types"
	"sync"
)

type QueueService struct {
	queue []*types.Player
	mu    sync.Mutex
}

func NewQueueService() *QueueService {
	return &QueueService{
		queue: make([]*types.Player, 0),
	}
}

// Pop a player from the queue
func (qs *QueueService) pop() *types.Player {
	qs.mu.Lock()
	defer qs.mu.Unlock()

	if len(qs.queue) == 0 {
		return nil
	}

	var lastPlayer *types.Player = qs.queue[len(qs.queue)-1]
	qs.queue = qs.queue[:len(qs.queue)-1]
	return lastPlayer
}

// Add a player to the queue
func (qs *QueueService) Push(player *types.Player) bool {
	qs.mu.Lock() 
	qs.queue = append(qs.queue, player)
	qs.mu.Unlock()
	return len(qs.queue)>=2

}

// Get two players and send them to the Room for interaction
func (qs *QueueService) GetPlayersForRoom() ([]types.Player, error) {
	qs.mu.Lock()
	defer qs.mu.Unlock()

	if len(qs.queue) < 2 {
		return nil, fmt.Errorf("not enough players in the queue")
	}

	// Pop two players from the queue
	player1 := *qs.pop()
	player2 := *qs.pop()

	return []types.Player{player1, player2}, nil
}
