package pokertablebalancer

type CreateTableReq struct {
	CompetitionId string `json:"competition_id"`
}

type CreateTableResp struct {
	Success bool   `json:"success"`
	TableId string `json:"table_id"`
}

type CloseTableReq struct {
	CompetitionId string `json:"competition_id"`
	TableId       string `json:"table_id"`
}

type CloseTableResp struct {
	Success bool `json:"success"`
}

type TableEntry struct {
	PlayerId string `json:"player_id"`
	TableId  string `json:"table_id"`
}

type JoinTableReq struct {
	CompetitionId string        `json:"competition_id"`
	Tables        []*TableEntry `json:"tables"`
}

type JoinTableResp struct {
	Success bool `json:"success"`
}

type ApiManager interface {
	AutoCreateTable(competitionId string) (*CreateTableResp, error)
	AutoCloseTable(competitionId string, tableId string) (*CloseTableResp, error)
	AutoJoinTable(competitionId string, entries []*TableEntry) (*JoinTableResp, error)
}
