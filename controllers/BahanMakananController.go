package controllers

import (
	"aplikasiresep/database"
	"aplikasiresep/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllBahan(c *fiber.Ctx) error {
	var bahan []models.BahanMakanan
	database.DB.Find(&bahan)
	if len(bahan) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status": "error!",
			"data":   nil,
		})
	}
	return c.JSON(&bahan)
}

func AddBahan(c *fiber.Ctx) error {
	bahan := new(models.BahanMakanan)

	if err := c.BodyParser(bahan); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	InputRaw := models.BahanMakanan{
		NamaBahan: "NamaBahan",
		Kuantitas: "Kuantitas",
		Resep_Id:  bahan.Resep_Id,
	}

	database.DB.Create(&InputRaw)
	return c.JSON(&InputRaw)
}

func UpdateBahan(c *fiber.Ctx) error {
	id := c.Params("id")
	bahan := new(models.BahanMakanan)
	database.DB.First(&bahan, id)

	if bahan.NamaBahan == "" {
		return c.Status(500).SendString("Produk tidak tersedia!")
	}

	if err := c.BodyParser(&bahan); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	database.DB.Save(&bahan)
	return c.JSON(&bahan)
}

func DeleteBahanbyId(c *fiber.Ctx) error {
	id := c.Params("id")
	bahan := new(models.BahanMakanan)
	database.DB.First(&bahan, id)

	if bahan.NamaBahan == "" {
		return c.Status(500).SendString("Produk tidak tersedia!")
	}

	database.DB.Delete(&bahan)
	return c.JSON("data telah terhapus!")
}

func GettBahanById(c *fiber.Ctx) error {
	id := c.Params("id")
	var bahan models.BahanMakanan
	database.DB.Where("id_resep = ?", id).Find(&bahan)
	return c.JSON(&bahan)
}
