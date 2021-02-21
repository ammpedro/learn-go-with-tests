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
