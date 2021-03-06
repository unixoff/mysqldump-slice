package relationship

import "database/sql"

type Relation struct {
	foreignTable string
	primaryTable string
	fkColumn string
}

func (rel *Relation) Parse(rows *sql.Rows) (err error) {
	return rows.Scan(&rel.foreignTable, &rel.primaryTable, &rel.fkColumn)
}

func (rel *Relation) FrTab() string {
	return rel.foreignTable
}

func (rel *Relation) PrTab() string {
	return rel.primaryTable
}

func (rel *Relation) FkCol() string {
	return rel.fkColumn
}
