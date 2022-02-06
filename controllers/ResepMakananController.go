package controllers

import (
	"aplikasiresep/database"
	"aplikasiresep/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllResep(c *fiber.Ctx) error {
	var resep []models.ResepMakanan
	database.DB.Find(&resep)
	if len(resep) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status": "error!",
			"data":   nil,
		})
	}
	return c.JSON(&resep)
}

func AddResep(c *fiber.Ctx) error {
	resep := new(models.ResepMakanan)

	if err := c.BodyParser(resep); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	InputRaw := models.ResepMakanan{
		NamaResep:   "NamaResep",
		Kategori_id: resep.Kategori_id,
	}

	database.DB.Create(&InputRaw)
	return c.JSON(&InputRaw)
}

func UpdateResep(c *fiber.Ctx) error {
	id := c.Params("id")
	resep := new(models.ResepMakanan)
	database.DB.First(&resep, id)

	if resep.NamaResep == "" {
		return c.Status(500).SendString("Produk tidak tersedia!")
	}

	if err := c.BodyParser(&resep); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	database.DB.Save(&resep)
	return c.JSON(&resep)
}

func DeleteResepbyId(c *fiber.Ctx) error {
	id := c.Params("id")
	resep := new(models.ResepMakanan)
	database.DB.First(&resep, id)

	if resep.NamaResep == "" {
		return c.Status(500).SendString("Produk tidak tersedia!")
	}

	database.DB.Delete(&resep)
	return c.JSON("data telah terhapus!")
}

func GettResepById(c *fiber.Ctx) error {
	id := c.Params("id")
	var resep models.ResepMakanan
	database.DB.First(&resep, id)
	return c.JSON(&resep)
}
