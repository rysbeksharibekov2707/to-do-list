package database

func CreateTasksTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title TEXT,
		details TEXT,
		done BOOLEAN
	)`
	_, err := DB.Exec(query)
	return err
}
