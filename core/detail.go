package core

type DetailStatus string

const None DetailStatus = "none"
const Error DetailStatus = "error"
const NotFound DetailStatus = "not_found"
const Found DetailStatus = "found"

type Detail struct {
	Id           string       `json:"id"`
	Status       DetailStatus `json:"status"`
	MessageError string       `json:"message_error"`
	Table        *Table       `json:"table"`
}

func NewDetail() *Detail {
	return &Detail{}
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
