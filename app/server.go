package main

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerStore returns scores
type PlayerStore interface {
	GetPlayerScore(name string) int
}

// PlayerServer stores scores
type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, p.store.GetPlayerScore(player))
}
