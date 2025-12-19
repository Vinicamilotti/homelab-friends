package store

import "fmt"

type DBHelper struct {
	Conn DBConnectionFacade
}

func NewDBHelper(conn DBConnectionFacade) *DBHelper {
	return &DBHelper{Conn: conn}
}

func (h *DBHelper) Insert(table string, fieldValue map[string]any) error {
	valuesParam := ""
	fields := ""
	values := []any{}
	for field, value := range fieldValue {
		fields += field + ", "
		valuesParam += "?, "
		values = append(values, value)
	}
	tx, err := h.Conn.Begin()
	if err != nil {
		return err
	}

	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, fields[:len(fields)-2], valuesParam[:len(valuesParam)-2])

	result, err := h.Conn.Exec(sql, values...)
	if err != nil {
		tx.Rollback()
		return err
	}
	r, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if r == 0 {
		return fmt.Errorf("no rows affected")
	}
	tx.Commit()
	return nil
}
