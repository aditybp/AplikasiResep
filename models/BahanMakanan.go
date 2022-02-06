package models

type BahanMakanan struct {
	Id        uint   `json:"Id"`
	NamaBahan string `json:"NamaBahan"`
	Kuantitas string `json:"Kuantitas"`
	Resep_Id  int    `json:"IdResep"`
}
