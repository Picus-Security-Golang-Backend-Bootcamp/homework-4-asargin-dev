package factory

import (
	"errors"
	"fmt"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-asargin-dev/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (b *Repository) Migrate() {
	b.db.AutoMigrate(&models.Book{})
}

func (b *Repository) InsertData(books []models.Book) {
	for _, book := range books {
		b.db.Where(models.Book{ID: book.ID}).Attrs(models.Book{StockNumber: book.StockNumber, PageNumber: book.PageNumber, Price: book.Price, Name: book.Name, StockCode: book.StockCode, Isbn: book.Isbn, AuthorName: book.AuthorName}).FirstOrCreate(&book)
	}
}

func (b *Repository) FindAll() []models.Book {
	var books []models.Book
	b.db.Find(&books)

	return books
}

func (b *Repository) FindByBookName(bookName string) {
	var books []models.Book
	b.db.Where("Name LIKE ?", "%"+bookName+"%").Find(&books)
}

func (b *Repository) FindById(id int) error {
	var book *models.Book
	result := b.db.First(&book, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Printf("Book could not found : %d", id)
	}
	return nil
}

func (b *Repository) Update(book *models.Book) error {
	result := b.db.Save(book)

	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Updated")
	return nil
}

func (b *Repository) DeleteById(id int) error {
	result := b.db.Delete(&models.Book{}, id)

	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Deleted")
	return nil
}
