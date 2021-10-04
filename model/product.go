package model

type Product struct {
	ProductID  int    `json:"produkid" form:"produkid"`
	CategoryID int    `json:"kategoriid" form:"kategoriid"`
	Nama       string `json:"nama" form:"nama"`
	Deskripsi  string `json:"deskripsi" form:"deskripsi"`
	Harga      int    `json:"harga" form:"harga"`
}
