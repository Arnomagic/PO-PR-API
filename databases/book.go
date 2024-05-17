package databases

import (
	"database/sql"
	"fmt"
	"hexapi/logs"
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/jmoiron/sqlx"
)

func NewBookDatabases(db *sqlx.DB) BookDb {
	return bookDb{db: db}
}
func (d bookDb) InsertBook(b Book) (*Book, error) {
	query := "INSERT INTO libraly_system.books (title, Author, publication_year, genre, members_id) VALUES ($1,$2,$3,$4,$5) RETURNING * ;"
	book := Book{}
	err := d.db.Get(&book, query, b.Title, b.Author, b.Publication_year, b.Genre)
	if err != nil {
		fmt.Println(err)
		return nil, ErrDB
	}
	return &book, nil
}
func (d bookDb) SelectAllBook(sr map[string]string) ([]Book, int, error) {
	query := "SELECT * FROM libraly_system.books "
	arg := []any{}
	if sr["srch"] != "" {
		query += "WHERE title LIKE '%' || $1 || '%' OR Author LIKE '%' || $1 || '%' "
		arg = append(arg, sr["srch"])
	}
	if sr["srt"] != "" {
		if sr["srt_fld"] != "" {
			if sr["srt_fld"] == "1" {
				query += "ORDER BY  " + "sub_name " + " " + sr["srt"]
			}
			if sr["srt_fld"] == "2" {
				query += "ORDER BY  " + "address " + " " + sr["srt"]
			}
			if sr["srt_fld"] == "3" {
				query += "ORDER BY  " + "contact " + " " + sr["srt"]
			}
			if sr["srt_fld"] == "4" {
				query += "ORDER BY  " + "tax_id " + " " + sr["srt"]
			}
		} else {
			query += "ORDER BY id " + sr["srt"]
		}
	}
	if sr["page"] != "" && sr["limit"] != "" {
		page, err := strconv.Atoi(sr["page"])
		if err != nil {
			logs.Log.Error(err.Error())
			return nil, 0, fiber.ErrBadRequest
		}
		limit, err := strconv.Atoi(sr["limit"])
		if err != nil {
			logs.Log.Error(err.Error())
			return nil, 0, fiber.ErrBadRequest
		}
		offset := (page - 1) * limit
		query += " OFFSET $" + strconv.Itoa(len(arg)+1) + " LIMIT $" + strconv.Itoa(len(arg)+2)
		arg = append(arg, offset, limit)
	}

	book := []Book{}
	err := d.db.Select(&book, query, arg...)
	if err != nil {
		fmt.Println(err)
		return nil, 0, ErrDB
	}
	if len(book) == 0 {
		return nil, 0, ErrNoRows
	}
	return book, 0, nil
}
func (d bookDb) SelectByIdBook(id int) (*Book, error) {
	query := "SELECT * FROM libraly_system.books WHERE books_id = $1;"
	books := Book{}
	err := d.db.Get(&books, query, id)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
		return nil, ErrDB
	}
	if err == sql.ErrNoRows {
		return nil, ErrNoRows
	}
	return &books, nil
}
func (d bookDb) UpdateByIdBook(b Book) (*Book, error) {
	fields_values := []struct {
		f string
		v any
	}{}
	if b.Title != "" {
		fields_values = append(fields_values, struct {
			f string
			v any
		}{"title", b.Title})
	}
	if b.Author != "" {
		fields_values = append(fields_values, struct {
			f string
			v any
		}{"author", b.Author})
	}
	if b.Publication_year != "" {
		fields_values = append(fields_values, struct {
			f string
			v any
		}{"publication_year", b.Publication_year})
	}
	if b.Genre != "" {
		fields_values = append(fields_values, struct {
			f string
			v any
		}{"genre", b.Genre})
	}
	query := "UPDATE libraly_system.books SET "
	field := ""
	argValue := []any{}
	for i, row := range fields_values {
		if i == 0 {
			field += "" + row.f + " = $" + fmt.Sprint((i + 1))
			argValue = append(argValue, row.v)
		} else {
			field += ", " + row.f + " = $" + fmt.Sprint((i + 1))
			argValue = append(argValue, row.v)
		}
	}
	query += field + " WHERE books_id = $" + fmt.Sprint(len(argValue)+1) + " RETURNING * ;"
	argValue = append(argValue, b.BookID)
	fmt.Println(query)
	fmt.Println(argValue...)
	book := Book{}
	err := d.db.Get(&book, query, argValue...)
	if err != nil {
		fmt.Println(err)
		return nil, ErrDB
	}
	return &book, nil
}
