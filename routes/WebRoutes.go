package routes

import (
	"aplikasiresep/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(App *fiber.App) {
	kategori := App.Group("/kategori")
	kategori.Get("/allkategori", controllers.GetAllKategori)
	kategori.Post("/addkategori", controllers.AddKategori)
	kategori.Patch("/updatekategori/:id", controllers.UpdateKategori)
	kategori.Delete("/deletekategori/:id", controllers.DeleteKategori)
	kategori.Get("/kategori/:id", controllers.GetKategoriById)

	bahan := App.Group("/bahan")
	bahan.Get("/allbahan", controllers.GetAllBahan)
	bahan.Post("/addbahan", controllers.AddBahan)
	bahan.Patch("/updatebahan/:id", controllers.UpdateBahan)
	bahan.Delete("/deletebahan/:id", controllers.DeleteBahanbyId)
	bahan.Get("/getbahan/:id", controllers.GettBahanById)

	resep := App.Group("/resep")
	resep.Get("/allresep", controllers.GetAllResep)
	resep.Post("/addresep", controllers.AddResep)
	resep.Patch("/updateresep/:id", controllers.UpdateResep)
	resep.Delete("/deleteresep/:id", controllers.DeleteResepbyId)
	resep.Get("/getresep/:id", controllers.GettResepById)

}
