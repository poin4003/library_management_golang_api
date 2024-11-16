package db

import (
	"fmt"
	"library_management/global"
	"library_management/models"
	"testing"
)

func TestInitPostgresql(t *testing.T) {
	InitPostgresql()

	if global.PDB == nil {
		fmt.Println("Failed to init postgres")
	}

	if !global.PDB.Migrator().HasTable(&models.Book{}) {
		fmt.Println("Table 'Book' was not created")
	}
}
