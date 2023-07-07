package core

type DetailStatus string

type Detail struct {
	Id     string       `json:"id"`
	Status DetailStatus `json:"status"`
	Table  *Table       `json:"table"`
}

func (d *Detail) LogName() string {
	return "detail"
}

func (d *Detail) LogProperties() map[string]interface{} {
	return map[string]interface{}{
		"s_id":     d.Id,
		"s_status": d.Status,
		"o_table":  d.Table,
	}
}
