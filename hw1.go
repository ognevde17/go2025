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

type Storage interface {
	AddBook(book Book)
	GetBookbyId(id int) Book
}
type Storage1 struct {
	Books map[int]Book
}

func NewStorage1() Storage1 {
	return Storage1{Books: make(map[int]Book)}
}

func GenerateRandomID() int {
	return rand.Intn(1_000_000_000)
}

func (storage *Storage1) AddBook(book Book) {
	storage.Books[book.id] = book
}
func (storage *Storage1) GetBookbyId(id int) Book {
	return storage.Books[id]
}

type Storage2 struct {
	Books []Book
}

func NewStorage2() Storage2 {
	return Storage2{Books: []Book{}}
}

func (storage *Storage2) AddBook(book Book) {
	storage.Books = append(storage.Books, book)
}

func (storage *Storage2) GetBookbyId(id int) Book {
	for _, book := range storage.Books {
		if book.id == id {
			return book
		}
	}
	return Book{}
}

type Library struct {
	Names map[string]int
	St    Storage
}

func NewLibrary(storage Storage) Library {
	return Library{Names: make(map[string]int), St: storage}
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
	fmt.Println("Testing Storage1 (map-based):")
	lib1 := NewLibrary(&Storage1{Books: make(map[int]Book)})

	b := Book{name: "Go 2025", author: "Anon", pages_count: 100}
	b1 := Book{name: "Go 20255", author: "Ano", pages_count: 10}

	id0 := lib1.AddBookWithId(b)
	id1 := lib1.AddBookWithId(b1)

	got := lib1.GetBookbyName("Go 20255")
	if got.name != "Go 20255" || got.author != "Ano" {
		panic("GetBookbyName returned wrong book")
	}
	fmt.Println("Storage1 OK:", id0, id1)

	fmt.Println("Testing Storage2 (slice-based):")
	lib2 := NewLibrary(&Storage2{Books: []Book{}})

	b2 := Book{name: "Go 2026", author: "Author2", pages_count: 200}
	b3 := Book{name: "Go 2027", author: "Author3", pages_count: 300}

	id2 := lib2.AddBookWithId(b2)
	id3 := lib2.AddBookWithId(b3)

	got2 := lib2.GetBookbyName("Go 2027")
	if got2.name != "Go 2027" || got2.author != "Author3" {
		panic("GetBookbyName returned wrong book")
	}
	fmt.Println("Storage2 OK:", id2, id3)
}
