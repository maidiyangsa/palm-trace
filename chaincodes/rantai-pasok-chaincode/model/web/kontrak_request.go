package web

import "rantai-pasok-chaincode/constant"

type KontrakCreateRequest struct {
	Id               string  `json:"id"`
	IdPks            string  `json:"idPks"`
	IdKoperasi       string  `json:"idKoperasi"`
	Nomor            string  `json:"nomor"`
	TanggalPembuatan string  `json:"tanggalPembuatan"`
	TangalMulai      string  `json:"tanggalMulai"`
	TanggalSelesai   string  `json:"tanggalSelesai"`
	Kuantitas        float32 `json:"kuantitas"`
	Harga            float64 `json:"harga"`
	CreatedAt        string  `json:"createdAt"`
	UpdatedAt        string  `json:"updatedAt"`
}

type KontrakConfirmRequest struct {
	Id                string                          `json:"id"`
	Status            constant.StatusPenawaranKontrak `json:"status"`
	Pesan             string                          `json:"pesan"`
	TanggalKonfirmasi string                          `json:"tanggalKonfirmasi"`
	UpdatedAt         string                          `json:"updatedAt"`
}

type KontrakUpdateKuantitasRequest struct {
	Id                 string  `json:"id"`
	KuantitasTerpenuhi float32 `json:"kuantitasTerpenuhi"`
	UpdatedAt          string  `json:"updatedAt"`
}

type KontrakFindAllRequest struct {
	IdPks      string                          `json:"idPks"`
	IdKoperasi string                          `json:"idKoperasi"`
	Status     constant.StatusPenawaranKontrak `json:"status"`
}
