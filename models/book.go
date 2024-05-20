package models

import "goreads/db"

type Book struct {
	Id          int64
	Title       string
	ISBN        string
	Author      string
	ReleaseYear int
}

func (b *Book) Save() error {
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

func GetAllBooks() ([]Book, error) {
	cursor, err := db.GetDb().Query("SELECT * FROM books")

	if err != nil {
		return nil, err
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
			return nil, err
		}
		bookCollection = append(bookCollection, obj)
	}

	return bookCollection, nil
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
