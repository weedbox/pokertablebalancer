package pokertablebalancer

import "github.com/weedbox/pokertablebalancer/psae"

type Competition struct {
	id      string
	manager *seatManager
	engine  psae.PSAE
}

func NewCompetition(id string, sm *seatManager, g *psae.Game) *Competition {

	c := &Competition{
		id:      id,
		manager: sm,
	}

	if g != nil {
		// Preparing game configuration
		g = psae.NewGame()
		g.MaxPlayersPerTable = 9
		g.MinInitialPlayers = 4
		g.TableLimit = -1
	}

	p := psae.NewPSAE(
		psae.WithBackend(NewBackend(c)),
		psae.WithRuntime(NewRuntime()),
		psae.WithGame(g),
		//TODO: implement SeatMap with Redis
		//TODO: implement WaitingRoom with Redis
		//TODO: implement MatchQueue with JetStream
		//TODO: implement DispatchQueue with JetStream
		//TODO: implement ReleaseQueue with JetStream
	)

	c.engine = p

	return c
}

func (c *Competition) Close() error {
	return c.engine.Close()
}

func (c *Competition) Join(playerId string) error {

	player := &psae.Player{
		ID: playerId,
	}

	err := c.engine.Join(player)
	if err != nil {
		return err
	}

	return nil
}

func (c *Competition) DisallowRegistration() {
	c.engine.DisallowRegistration()
}

func (c *Competition) GetTableState(tableId string) (*psae.TableState, error) {
	return c.engine.GetTableState(tableId)
}

func (c *Competition) UpdateTable(state *psae.TableState) (*psae.TableState, error) {
	return c.engine.UpdateTableState(state)
}
