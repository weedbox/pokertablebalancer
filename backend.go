package pokertablebalancer

import "github.com/weedbox/pokertablebalancer/psae"

// TODO: ASK Fred
func NewBackend(c *Competition) *psae.Backend {

	b := psae.NewBackend()

	b.AllocateTable = func() (*psae.TableState, error) {
		return nil, nil

		// resp, err := c.manager.api.CreateNakamaTable(c.id)
		// if err != nil {
		// 	return nil, err
		// }

		// ts := &psae.TableState{
		// 	ID:             tableID,
		// 	Players:        make(map[string]*psae.Player),
		// 	Status:         psae.TableStatus_Ready,
		// 	TotalSeats:     c.engine.Game().MaxPlayersPerTable,
		// 	AvailableSeats: c.engine.Game().MaxPlayersPerTable,
		// 	Statistics: &psae.TableStatistics{
		// 		NoChanges: 0,
		// 	},
		// }

		// return ts, nil

	}

	b.JoinTable = func(tableId string, players []*psae.Player) error {

		// entries := make([]*api.MatchEntry, 0)
		// for _, p := range players {
		// 	entries = append(entries, &api.MatchEntry{
		// 		PlayerId: p.ID,
		// 		MatchId:  tableId,
		// 	})
		// }
		// _, err := c.manager.api.JoinNakamaTable(c.id, entries)
		// if err != nil {
		// 	return err
		// }

		return nil
	}
	b.BrokeTable = func(tableId string) error {

		// _, err := c.manager.api.CloseNakamaTable(c.id, tableId)
		// if err != nil {
		// 	return err
		// }

		return nil
	}

	return b
}
