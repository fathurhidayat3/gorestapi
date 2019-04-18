package controllers

import (
	"net/http"

	"github.com/fathurhidayat3/gorestapi/structs"
	"github.com/gin-gonic/gin"
)

// GetBuku to get one data with {id}
func (idb *InDB) GetBuku(c *gin.Context) {
	var (
		buku   structs.Buku
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&buku).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": buku,
			"count":  1,
		}
	}

	c.JSON(http.StatusOK, result)
}

// GetBukus get all data in buku
func (idb *InDB) GetBukus(c *gin.Context) {
	var (
		buku   []structs.Buku
		result gin.H
	)

	idb.DB.Find(&buku)
	if len(buku) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": buku,
			"count":  len(buku),
		}
	}

	c.JSON(http.StatusOK, result)
}

// CreateBuku create new data to the database
func (idb *InDB) CreateBuku(c *gin.Context) {
	var (
		buku   structs.Buku
		result gin.H
	)

	judulBuku := c.PostForm("judulBuku")
	jumlahHalaman := c.PostForm("jumlahHalaman")
	isbn := c.PostForm("isbn")
	penulis := c.PostForm("penulis")
	tahunTerbit := c.PostForm("tahunTerbit")

	buku.JudulBuku = judulBuku
	buku.JumlahHalaman = jumlahHalaman
	buku.Isbn = isbn
	buku.Penulis = penulis
	buku.TahunTerbit = tahunTerbit

	idb.DB.Create(&buku)

	result = gin.H{
		"result": buku,
	}
	c.JSON(http.StatusOK, result)
}

// UpdateBuku update data with {id} as query
func (idb *InDB) UpdateBuku(c *gin.Context) {
	id := c.Query("id")
	judulBuku := c.PostForm("judulBuku")
	jumlahHalaman := c.PostForm("jumlahHalaman")
	isbn := c.PostForm("isbn")
	penulis := c.PostForm("penulis")
	tahunTerbit := c.PostForm("tahunTerbit")

	var (
		buku    structs.Buku
		newBuku structs.Buku
		result  gin.H
	)

	err := idb.DB.First(&buku, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}

	newBuku.JudulBuku = judulBuku
	newBuku.JumlahHalaman = jumlahHalaman
	newBuku.Isbn = isbn
	newBuku.Penulis = penulis
	newBuku.TahunTerbit = tahunTerbit

	err = idb.DB.Model(&buku).Updates(newBuku).Error
	if err != nil {
		result = gin.H{
			"result": "update failed",
		}
	} else {
		result = gin.H{
			"result": "successfully updated data",
		}
	}
	c.JSON(http.StatusOK, result)
}

// DeleteBuku delete data with {id}
func (idb *InDB) DeleteBuku(c *gin.Context) {
	var (
		buku   structs.Buku
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&buku, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
	}
	err = idb.DB.Delete(&buku).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
	} else {
		result = gin.H{
			"result": "Data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}
