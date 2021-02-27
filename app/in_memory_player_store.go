package main

//NewInMemoryPlayerStore stores player scores
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

//InMemoryPlayerStore is the player score object
type InMemoryPlayerStore struct {
	store map[string]int
}

//RecordWin records a player's wins
func (i InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

//GetPlayerScore returns a player's score
func (i InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

//GetLeague returns an array of players and wins
func (i InMemoryPlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}
