package models

import "database/sql"

type Manufacturer struct {
	ID          uint16         `json:"id"`
	Name        string         `json:"name"`
	FoundedYear int            `json:"founded_year"`
	Nationality string         `json:"nationality"`
	Logo        sql.NullString `json:"logo"`
}
