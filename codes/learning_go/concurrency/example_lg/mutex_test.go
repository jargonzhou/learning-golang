package example_lg

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

// Example of sync.Mutex, syncRWMutex. - 'Learning Go': 12. Concurrency in Go, P.313
// in-memory scoreboard for multiplayer game.

type Scoreboard map[string]int

func (s Scoreboard) Update(name string, val int) {
	s[name] = val
}

func (s Scoreboard) Read(name string) (int, bool) {
	val, ok := s[name]
	return val, ok
}

// channel based version

func scoreboardManager(ctx context.Context, in <-chan func(Scoreboard)) {
	scorebard := Scoreboard{}
	for {
		select {
		case <-ctx.Done():
			return
		case f := <-in: // process input
			f(scorebard)
		}
	}
}

type ChannelScorebardManager chan func(Scoreboard)

func NewChannelScoreboardManager(ctx context.Context) ChannelScorebardManager {
	ch := make(ChannelScorebardManager)
	go scoreboardManager(ctx, ch)
	return ch
}

func (csm ChannelScorebardManager) Update(name string, val int) {
	csm <- func(s Scoreboard) {
		s.Update(name, val)
	}
}

func (csm ChannelScorebardManager) Read(name string) (int, bool) {
	type Result struct {
		out int
		ok  bool
	}
	resultCh := make(chan Result)

	csm <- func(s Scoreboard) {
		out, ok := s.Read(name)
		resultCh <- Result{out, ok}
	}

	result := <-resultCh
	return result.out, result.ok
}

// mutex based version
type MutexScoreboardManager struct {
	l          sync.RWMutex
	scoreboard Scoreboard
}

func NewMutexScoreboardManager() *MutexScoreboardManager {
	return &MutexScoreboardManager{
		scoreboard: Scoreboard{},
	}
}

func (msm *MutexScoreboardManager) Update(name string, val int) {
	msm.l.Lock()
	defer msm.l.Unlock()

	msm.scoreboard.Update(name, val)
}

func (msm *MutexScoreboardManager) Read(name string) (int, bool) {
	msm.l.RLock()
	defer msm.l.RUnlock()

	val, ok := msm.scoreboard.Read(name)
	return val, ok
}

func TestMutexScoreboardManager(t *testing.T) {
	msm := NewMutexScoreboardManager()
	teams := []string{"A", "B", "C"}

	var wg sync.WaitGroup
	wg.Add(len(teams))

	// spawn len(teams) goroutines
	for _, v := range teams {
		go func(team string) {
			defer wg.Done()

			// read and write 10 times
			for i := 0; i < 10; i++ {
				curScore, found := msm.Read(team)
				if !found {
					curScore = 10
				} else {
					curScore += len(team)
				}
				msm.Update(team, curScore)
			}
		}(v)
	}

	// wait for job done
	wg.Wait()

	for _, v := range teams {
		score, found := msm.Read(v)
		fmt.Println(v, score, found)
	}
}
