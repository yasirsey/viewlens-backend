package repositories

import (
	"database/sql"
	"fmt"
)

type ViewRepository struct {
	db *sql.DB
}

func NewViewRepository(db *sql.DB) *ViewRepository {
	return &ViewRepository{db}
}

func (vr *ViewRepository) ListViews() ([]string, error) {
	query := `
		SELECT TABLE_NAME 
		FROM INFORMATION_SCHEMA.VIEWS 
		WHERE TABLE_NAME LIKE 'VEX_%'
	`
	rows, err := vr.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var views []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, err
		}
		views = append(views, tableName)
	}
	return views, nil
}

func (vr *ViewRepository) GetViewDetails(viewName string) ([]string, []map[string]interface{}, error) {
	// Get columns
	columnQuery := fmt.Sprintf(`
		SELECT COLUMN_NAME 
		FROM INFORMATION_SCHEMA.COLUMNS 
		WHERE TABLE_NAME = '%s'`, viewName)

	columnRows, err := vr.db.Query(columnQuery)
	if err != nil {
		return nil, nil, err
	}
	defer columnRows.Close()

	var columns []string
	for columnRows.Next() {
		var columnName string
		if err := columnRows.Scan(&columnName); err != nil {
			return nil, nil, err
		}
		columns = append(columns, columnName)
	}

	// Get data
	dataQuery := fmt.Sprintf("SELECT * FROM %s", viewName)
	dataRows, err := vr.db.Query(dataQuery)
	if err != nil {
		return nil, nil, err
	}
	defer dataRows.Close()

	var data []map[string]interface{}
	for dataRows.Next() {
		row := make(map[string]interface{})
		columnData := make([]interface{}, len(columns))
		columnPointers := make([]interface{}, len(columns))
		for i := range columnData {
			columnPointers[i] = &columnData[i]
		}

		if err := dataRows.Scan(columnPointers...); err != nil {
			return nil, nil, err
		}
		for i, colName := range columns {
			row[colName] = columnData[i]
		}
		data = append(data, row)
	}

	return columns, data, nil
}
