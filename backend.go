package pokertablebalancer

import "github.com/weedbox/pokertablebalancer/psae"

func NewBackend(c *Competition) *psae.Backend {

	b := psae.NewBackend()

	b.AllocateTable = func() (*psae.TableState, error) {

		resp, err := c.manager.api.AutoCreateTable(c.id)
		if err != nil {
			return nil, err
		}

		ts := &psae.TableState{
			ID:             resp.TableId,
			Players:        make(map[string]*psae.Player),
			Status:         psae.TableStatus_Ready,
			TotalSeats:     c.engine.Game().MaxPlayersPerTable,
			AvailableSeats: c.engine.Game().MaxPlayersPerTable,
			Statistics: &psae.TableStatistics{
				NoChanges: 0,
			},
		}

		return ts, nil

	}

	b.JoinTable = func(tableId string, players []*psae.Player) error {

		entries := make([]*TableEntry, 0)
		for _, p := range players {
			entries = append(entries, &TableEntry{
				PlayerId: p.ID,
				TableId:  tableId,
			})
		}
		_, err := c.manager.api.AutoJoinTable(c.id, entries)
		if err != nil {
			return err
		}

		return nil
	}
	b.BrokeTable = func(tableId string) error {

		_, err := c.manager.api.AutoCloseTable(c.id, tableId)
		if err != nil {
			return err
		}

		return nil
	}

	return b
}
