package controllers

import (
	"aplikasiresep/database"
	"aplikasiresep/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllKategori(c *fiber.Ctx) error {
	var kategori []models.KategoriMakanan
	database.DB.Find(&kategori)
	if len(kategori) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status": "error!",
			"data":   nil,
		})
	}

	return c.JSON(&kategori)
}

func AddKategori(c *fiber.Ctx) error {
	kategori := new(models.KategoriMakanan)

	if err := c.BodyParser(kategori); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	InputRaw := models.KategoriMakanan{
		NamaKategori: "NamaKategori",
	}

	database.DB.Create(&InputRaw)
	return c.JSON(&InputRaw)
}

func UpdateKategori(c *fiber.Ctx) error {
	id := c.Params("id")
	kategori := new(models.KategoriMakanan)
	database.DB.First(&kategori, id)

	if kategori.NamaKategori == "" {
		return c.Status(500).SendString("Produk tidak tersedia!")
	}

	if err := c.BodyParser(&kategori); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	database.DB.Save(&kategori)
	return c.JSON(&kategori)
}

func DeleteKategori(c *fiber.Ctx) error {
	id := c.Params("id")
	kategori := new(models.KategoriMakanan)
	database.DB.First(&kategori, id)

	if kategori.NamaKategori == "" {
		return c.Status(500).SendString("Produk tidak tersedia!")
	}

	database.DB.Delete(&kategori)
	return c.JSON("data telah terhapus!")
}

func GetKategoriById(c *fiber.Ctx) error {
	id := c.Params("id")
	var kategori models.KategoriMakanan
	database.DB.Find(&kategori, id)
	return c.JSON(&kategori)
}
