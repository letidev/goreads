package models

import (
	"fmt"
	"goreads/db"
)

type Book struct {
	Id          int64
	Title       string
	ISBN        string
	Author      string
	ReleaseYear int
}

func (b *Book) SaveNew() error {
	statement, err := db.GetDb().Prepare(`
		INSERT INTO
		books
			(title, isbn, author, release_year)
		VALUES
			(?, ?, ?, ?)
	`)

	if err != nil {
		return err
	}

	defer statement.Close()

	result, err := statement.Exec(b.Title, b.ISBN, b.Author, b.ReleaseYear)

	if err != nil {
		return err
	}

	bookid, err := result.LastInsertId()
	b.Id = bookid

	return err
}

// Returns the selected rows, total number of rows, number of selected rows, total number of pages, and error
func GetBooksPage(page int) ([]Book, int, int, int, error) {
	var pageSize int = 5
	var offset int = (page - 1) * pageSize

	cursor, err := db.GetDb().Query(fmt.Sprintf("SELECT * FROM books LIMIT %d OFFSET %d", pageSize, offset))

	if err != nil {
		return nil, 0, 0, 0, err
	}

	countCursor, err := db.GetDb().Query("SELECT COUNT(*) FROM books")

	if err != nil {
		return nil, 0, 0, 0, err
	}

	var totalRows int
	// should execute only once
	for countCursor.Next() {
		err = countCursor.Scan(&totalRows)

		if err != nil {
			return nil, 0, 0, 0, err
		}
	}

	// for each select we need an empty collection
	// if it's a global one, the new items get appended
	// to the old items and get duplicated
	bookCollection := []Book{}

	for cursor.Next() {
		var obj Book

		err := cursor.Scan(
			&obj.Id,
			&obj.Title,
			&obj.ISBN,
			&obj.Author,
			&obj.ReleaseYear,
		)

		if err != nil {
			return nil, 0, 0, 0, err
		}
		bookCollection = append(bookCollection, obj)
	}

	var totalPages int = totalRows / pageSize
	if totalRows%pageSize != 0 {
		totalPages++
	}

	return bookCollection, totalRows, len(bookCollection), totalPages, nil
}

func GetOneBook(id int64) (*Book, error) {
	row := db.GetDb().QueryRow("SELECT * FROM books WHERE id=?", id)

	var obj Book

	err := row.Scan(
		&obj.Id,
		&obj.Title,
		&obj.ISBN,
		&obj.Author,
		&obj.ReleaseYear)

	if err != nil {
		return nil, err
	}

	return &obj, nil
}

func (b Book) SaveExisting() error {
	statement, err := db.GetDb().Prepare(`
		UPDATE books
		SET title=?, isbn=?, author=?, release_year=?
		WHERE id=?
	`)

	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(b.Title, b.ISBN, b.Author, b.ReleaseYear, b.Id)

	return err
}

func Delete(id int64) error {
	statement, err := db.GetDb().Prepare(`
		DELETE FROM books
		WHERE id=?
	`)

	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	return err
}
