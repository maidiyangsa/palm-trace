package domain

import "rantai-pasok-chaincode/constant"

type Pembayaran struct {
	Id                    string             `json:"id"`
	AssetType             constant.AssetType `json:"assetType"`
	IdTransaksi           string             `json:"idTransaksi"`
	JenisUser             string             `json:"jenisUser"`
	Nomor                 string             `json:"nomor"`
	Tanggal               string             `json:"tanggal"`
	NomorRekeningPengirim string             `json:"nomorRekeningPengirim"`
	NomorRekeningPenerima string             `json:"nomorRekeningPenerima"`
	CidBuktiPembayaran    string             `json:"cidBuktiPembayaran"`
	CreatedAt             string             `json:"createdAt"`
	UpdatedAt             string             `json:"updatedAt"`
}
