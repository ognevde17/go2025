package main

import "fmt"
import "math/rand"
import "time"

type Book struct {
	name        string
	author      string
	pages_count int
	id          int
}

func (book *Book) Get_Name() string {
	return book.name
}
func (book *Book) Set_Id(new_id int) {
	book.id = new_id
}

type Storage struct {
	Books map[int]Book
}

func NewStorage() Storage {
	return Storage{Books: make(map[int]Book)}
}

func GenerateRandomID() int {
	return rand.Intn(1_000_000_000)
}

func (storage *Storage) AddBook(book Book) {
	storage.Books[book.id] = book
}
func (storage *Storage) GetBookbyId(id int) Book {
	return storage.Books[id]
}

type Library struct {
	Names map[string]int
	St    Storage
}

func NewLibrary() Library {
	return Library{Names: make(map[string]int), St: NewStorage()}
}

func (library *Library) AddBookWithId(book Book) int {
	id := GenerateRandomID()
	book.Set_Id(id)
	library.Names[book.Get_Name()] = id
	library.St.AddBook(book)
	return id
}
func (library *Library) GetBookbyName(name string) Book {
	return library.St.GetBookbyId(library.Names[name])
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	lib := NewLibrary()

	b := Book{name: "Go 2025", author: "Anon", pages_count: 100}
	b1 := Book{name: "Go 20255", author: "Ano", pages_count: 10}

	id0 := lib.AddBookWithId(b)
	id1 := lib.AddBookWithId(b1)

	got := lib.GetBookbyName("Go 20255")
	if got.name != "Go 20255" || got.author != "Ano" {
		panic("GetBookbyName returned wrong book")
	}
	fmt.Println("OK:", id0, id1)
}
