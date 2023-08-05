package web

type PembayaranCreateRequest struct {
	Id                    string `json:"id"`
	IdTransaksi           string `json:"idTransaksi"`
	JenisUser             string `json:"jenisUser"`
	Nomor                 string `json:"nomor"`
	Tanggal               string `json:"tanggal"`
	NomorRekeningPengirim string `json:"nomorRekeningPengirim"`
	NomorRekeningPenerima string `json:"nomorRekeningPenerima"`
	CidBuktiPembayaran    string `json:"cidBuktiPembayaran"`
	CreatedAt             string `json:"createdAt"`
	UpdatedAt             string `json:"updatedAt"`
}

type PembayaranFindAllRequest struct {
	IdTransaksi string `json:"idTransaksi"`
}
