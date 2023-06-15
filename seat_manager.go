package pokertablebalancer

import (
	"github.com/weedbox/pokertablebalancer/psae"
)

type SeatManager interface {
	RegisterCompetition(cId string, game *psae.Game) *Competition
	UnregisterCompetition(cId string) error
	GetCompetition(cId string) *Competition
}

type seatManager struct {
	api          ApiManager
	competitions map[string]*Competition
}

func NewSeatManager(api ApiManager) SeatManager {
	return &seatManager{
		api:          api,
		competitions: make(map[string]*Competition),
	}
}

func (sm *seatManager) RegisterCompetition(cId string, g *psae.Game) *Competition {

	if c, ok := sm.competitions[cId]; ok {
		return c
	}

	c := NewCompetition(cId, sm, g)

	sm.competitions[cId] = c

	return c
}

func (sm *seatManager) UnregisterCompetition(cId string) error {

	if c, ok := sm.competitions[cId]; ok {
		c.Close()
		return nil
	}

	delete(sm.competitions, cId)

	return nil
}

func (sm *seatManager) GetCompetition(cId string) *Competition {

	if c, ok := sm.competitions[cId]; ok {
		return c
	}

	return nil
}
