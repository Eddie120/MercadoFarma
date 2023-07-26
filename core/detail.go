package core

type DetailStatus string

const None DetailStatus = "none"
const Error DetailStatus = "error"
const NotFound DetailStatus = "not_found"
const Found DetailStatus = "found"

type Detail struct {
	Id             string       `json:"id"`
	CanonicalQuery string       `json:"canonical_query"` // concatenate(ip::query::date)
	Status         DetailStatus `json:"status"`
	MessageError   string       `json:"message_error"`
	Table          *Table       `json:"table"`
}

func NewDetail() *Detail {
	return &Detail{
		Status: None,
	}
}

func (d *Detail) LogName() string {
	return "detail"
}

func (d *Detail) LogProperties() map[string]interface{} {
	return map[string]interface{}{
		"s_id":              d.Id,
		"s_canonical_query": d.CanonicalQuery,
		"s_status":          d.Status,
		"s_message_error":   d.MessageError,
		"o_table":           d.Table,
	}
}
