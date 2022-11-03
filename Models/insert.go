package Models

import (
	"api-golang/db"
	"database/sql"
)

func Insert(todo Todo) (id int64, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}
	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	sql := `INSERT INTO todos (title,descrfiption,done) VALUES ($1,$2,$3) RETURN id`

	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)

	return
}
