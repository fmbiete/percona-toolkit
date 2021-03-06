// Package models contains the types for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

// AllDatabases represents a row from '[custom all_databases]'.
type AllDatabases struct {
	Datname Name // datname
}

// GetAllDatabases runs a custom query, returning results as AllDatabases.
func GetAllDatabases(db XODB) ([]*AllDatabases, error) {
	var err error

	// sql query
	var sqlstr = `SELECT datname ` +
		`FROM pg_database ` +
		`WHERE datistemplate = false`

	// run query
	XOLog(sqlstr)
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AllDatabases{}
	for q.Next() {
		ad := AllDatabases{}

		// scan
		err = q.Scan(&ad.Datname)
		if err != nil {
			return nil, err
		}

		res = append(res, &ad)
	}

	return res, nil
}
