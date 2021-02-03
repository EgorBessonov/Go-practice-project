package main

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type note struct {
	category  string
	id        int
	subject   string
	content   string
	createdAt string
	updatedAt string
}

func _updateNote(n note) (err error) {
	db, err := sql.Open("sqlite3", "notesdb.db")
	if err != nil {
		return err
	}
	defer db.Close()
	n.updatedAt = time.Now().String()
	_, err = db.Exec("UPDATE notes SET subject = $1, content = $2, updatedAt = $3, category = $4 WHERE id = $5", n.subject, n.content, n.updatedAt, n.category, n.id)
	if err != nil {
		return err
	}
	return
}

func _deleteNote(id int) (err error) {
	db, err := sql.Open("sqlite3", "notesdb.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM notes where id = $1", id)
	if err != nil {
		return err
	}
	return
}

func _addNote(n note) (err error) {
	n.createdAt = time.Now().String()
	n.updatedAt = time.Now().String()
	db, err := sql.Open("sqlite3", "notesdb.db")
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO notes (category, subject, content, createdAt, updatedAt) VALUES ($1, $2, $3, $4, $5)", n.category, n.subject, n.content, n.createdAt, n.updatedAt)
	if err != nil {
		return err
	}
	return err
}

func _getNotes() (notes []note, err error) {

	db, err := sql.Open("sqlite3", "notesdb.db")
	if err != nil {
		return notes, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM notes")
	if err != nil {
		return notes, err
	}
	defer rows.Close()

	for rows.Next() {
		n := note{}
		err := rows.Scan(&n.id, &n.category, &n.subject, &n.content, &n.createdAt, &n.updatedAt)
		if err != nil {
			return notes, err
		}
		notes = append(notes, n)
	}
	return notes, err
}

func _getNote(id int) (n note, err error) {
	db, err := sql.Open("sqlite3", "notesdb.db")
	if err != nil {
		return n, err
	}
	defer db.Close()
	row := db.QueryRow("SELECT * FROM notes WHERE id = $1", id)

	err = row.Scan(&n.id, &n.category, &n.subject, &n.content, &n.createdAt, &n.updatedAt)
	return
}
