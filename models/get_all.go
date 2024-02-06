package models

import "Estudos/db"

func getAll() (todos []Todo, err error) {
	conn, err := db.OpenConnection()

	rows, err := conn.Query(`SELECT * FROM todos`)

	if err != nil {
		return
	}

	defer conn.Close()

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}

	return
}