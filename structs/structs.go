package structs

import "github.com/jinzhu/gorm"

// Buku -
type Buku struct {
	gorm.Model
	JudulBuku     string
	Isbn          string
	Penulis       string
	TahunTerbit   string
	JumlahHalaman string
}

// TableName -
func (Buku) TableName() string {
	return "buku"
}
