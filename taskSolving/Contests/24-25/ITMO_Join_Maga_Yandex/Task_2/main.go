package main

import (
	"fmt"
)

type State struct {
	a, b, c, d int
}

func (s State) String() string {
	return fmt.Sprintf("%d%d%d%d", s.a, s.b, s.c, s.d)
}

func getNextStates(state State) []State {
	var nextStates []State

	// Case 1: Fix 2 performance bugs in Wi
	if state.a >= 2 {
		nextStates = append(nextStates, State{state.a - 2, state.b, state.c + 1, state.d + 1})
	}

	// Case 2: Fix 2 memory bugs in Wi
	if state.b >= 2 {
		nextStates = append(nextStates, State{state.a, state.b - 2, state.c + 1, state.d + 1})
	}

	// Case 3: Fix 2 performance bugs in Li
	if state.c >= 2 {
		nextStates = append(nextStates, State{state.a + 1, state.b + 1, state.c - 2, state.d})
	}

	// Case 4: Fix 2 memory bugs in Li
	if state.d >= 2 {
		nextStates = append(nextStates, State{state.a + 1, state.b + 1, state.c, state.d - 2})
	}

	return nextStates
}

func main() {
	initialState := State{2, 0, 2, 3}
	visited := make(map[string]bool)
	queue := []State{initialState}
	visited[initialState.String()] = true

	for len(queue) > 0 {
		currentState := queue[0]
		fmt.Println(currentState)
		queue = queue[1:]

		for _, nextState := range getNextStates(currentState) {
			stateStr := nextState.String()
			if !visited[stateStr] {
				visited[stateStr] = true
				queue = append(queue, nextState)
			}
		}
	}

	fmt.Println("Количество уникальных состояний:", len(visited))
	//fmt.Println(visited)
}
