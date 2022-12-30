package ai

import (
	"math/rand"
	"time"

	"github.com/kitsunemikan/six-purrpurrs/game"
	"github.com/kitsunemikan/six-purrpurrs/geom"
)

type ObstructingPlayer struct {
	Me   game.PlayerID
	rand *rand.Rand
}

func NewObstructingPlayer(id game.PlayerID) game.PlayerAgent {
	source := rand.NewSource(time.Now().UnixMicro())
	return &ObstructingPlayer{
		Me:   id,
		rand: rand.New(source),
	}
}

func (p *ObstructingPlayer) MakeMove(b *game.BoardState) geom.Offset {
	// Collect shifts
	dirs := make([]int, len(game.StrikeDirs))
	for i := range dirs {
		dirs[i] = i
	}

	for i := 0; i < len(dirs); i++ {
		swapID := p.rand.Intn(len(dirs)-1) + 1
		dirs[0], dirs[swapID] = dirs[swapID], dirs[0]
	}

	for opponentCell := range b.PlayerCells()[p.Me.Other()] {
		for i := 0; i < len(dirs); i++ {
			cell := opponentCell.Add(game.StrikeDirs[dirs[i]].Offset())
			if _, ok := b.UnoccupiedCells()[cell]; ok {
				return cell
			}

			cell = opponentCell.Sub(game.StrikeDirs[dirs[i]].Offset())
			if _, ok := b.UnoccupiedCells()[cell]; ok {
				return cell
			}
		}
	}

	// If all opponent's cells are obstructed, choose unoccupied at random
	for cell := range b.UnoccupiedCells() {
		return cell
	}

	panic("obstructing player: no unoccupied cells were present at all!")
}