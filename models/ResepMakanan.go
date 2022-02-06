package models

type ResepMakanan struct {
	Id          uint   `json:"Id"`
	NamaResep   string `json:"NamaResep"`
	Kategori_id int    `json:"IdKategori"`
}
