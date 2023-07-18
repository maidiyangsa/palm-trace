package web

type ReferensiHargaResponse struct {
	IdTransaksiBlockchain string  `json:"idTransaksiBlockchain,omitempty" metadata:",optional"`
	Id                    string  `json:"id"`
	IdDinas               string  `json:"idDinas"`
	UmurTanam             int     `json:"umurTanam"`
	Harga                 float64 `json:"harga"`
	TanggalPembaruan      string  `json:"tanggalPembaruan"`
}
