package domain

import "rantai-pasok-chaincode/constant"

type DeliveryOrder struct {
	Id                string                                `json:"id"`
	AssetType         constant.AssetType                    `json:"assetType"`
	IdKontrak         string                                `json:"idKontrak"`
	Nomor             string                                `json:"nomor"`
	TanggalPembuatan  string                                `json:"tanggalPembuatan"`
	Periode           string                                `json:"periode"`
	Kuantitas         float32                               `json:"kuantitas"`
	Harga             float64                               `json:"harga"`
	Rendemen          float32                               `json:"rendemen"`
	Status            constant.StatusPenawaranDeliveryOrder `json:"status"`
	Pesan             string                                `json:"pesan"`
	TanggalRespons    string                                `json:"tanggalRespons"`
	KuantitasTepenuhi float32                               `json:"kuantitasTepenuhi"`
	KuantitasTersisa  float32                               `json:"kuantitasTersisa"`
	CreatedAt         string                                `json:"createdAt"`
	UpdatedAt         string                                `json:"updatedAt"`
}
