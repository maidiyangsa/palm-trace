package helper

import (
	"rantai-pasok-chaincode/model/domain"
	"rantai-pasok-chaincode/model/web"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-protos-go/ledger/queryresult"
)

func ToWebResponse(status int, data interface{}, err error) *web.WebResponse {
	webResponse := &web.WebResponse{
		Status: status,
	}

	if data != nil {
		webResponse.Data = data
	}

	if err != nil {
		webResponse.Message = err.Error()
	}

	return webResponse
}

func ToKebunResponse(ctx contractapi.TransactionContextInterface, keyModification *queryresult.KeyModification, kebun *domain.Kebun) *web.KebunResponse {
	kebunResponse := &web.KebunResponse{
		Id:             kebun.Id,
		IdPetani:       kebun.IdPetani,
		Alamat:         kebun.Alamat,
		Latitude:       kebun.Latitude,
		Longitude:      kebun.Longitude,
		Luas:           kebun.Luas,
		NomorRspo:      kebun.NomorRspo,
		SertifikatRspo: kebun.SertifikatRspo,
		CreatedAt:      kebun.CreatedAt,
		UpdatedAt:      kebun.UpdatedAt,
	}

	if ctx != nil {
		kebunResponse.IdTransaksiBlockchain = ctx.GetStub().GetTxID()
	}

	if keyModification != nil {
		kebunResponse.IdTransaksiBlockchain = keyModification.GetTxId()
	}

	return kebunResponse
}

func ToKontrakResponse(ctx contractapi.TransactionContextInterface, keyModification *queryresult.KeyModification, kontrak *domain.Kontrak) *web.KontrakResponse {
	kontrakResponse := &web.KontrakResponse{
		Id:                kontrak.Id,
		IdPks:             kontrak.IdPks,
		IdKoperasi:        kontrak.IdKoperasi,
		Nomor:             kontrak.Nomor,
		TanggalPembuatan:  kontrak.TanggalPembuatan,
		TangalMulai:       kontrak.TangalMulai,
		TanggalSelesai:    kontrak.TanggalSelesai,
		Kuantitas:         kontrak.Kuantitas,
		Harga:             kontrak.Harga,
		Status:            kontrak.Status.String(),
		Pesan:             kontrak.Pesan,
		TanggalRespons:    kontrak.TanggalRespons,
		KuantitasTepenuhi: kontrak.KuantitasTepenuhi,
		KuantitasTersisa:  kontrak.KuantitasTersisa,
		CreatedAt:         kontrak.CreatedAt,
		UpdatedAt:         kontrak.UpdatedAt,
	}

	if ctx != nil {
		kontrakResponse.IdTransaksiBlockchain = ctx.GetStub().GetTxID()
	}

	if keyModification != nil {
		kontrakResponse.IdTransaksiBlockchain = keyModification.GetTxId()
	}

	return kontrakResponse
}

func ToDeliveryOrderResponse(ctx contractapi.TransactionContextInterface, keyModification *queryresult.KeyModification, deliveryOrder *domain.DeliveryOrder) *web.DeliveryOrderResponse {
	deliveryOrderResponse := &web.DeliveryOrderResponse{
		Id:               deliveryOrder.Id,
		IdKontrak:        deliveryOrder.IdKontrak,
		Nomor:            deliveryOrder.Nomor,
		TanggalPembuatan: deliveryOrder.TanggalPembuatan,
		Periode:          deliveryOrder.Periode,
		Kuantitas:        deliveryOrder.Kuantitas,
		Harga:            deliveryOrder.Harga,
		Rendemen:         deliveryOrder.Rendemen,
		Status:           deliveryOrder.Status.String(),
		Pesan:            deliveryOrder.Pesan,
		TanggalRespons:   deliveryOrder.TanggalRespons,
		CreatedAt:        deliveryOrder.CreatedAt,
		UpdatedAt:        deliveryOrder.UpdatedAt,
	}

	if ctx != nil {
		deliveryOrderResponse.IdTransaksiBlockchain = ctx.GetStub().GetTxID()
	}

	if keyModification != nil {
		deliveryOrderResponse.IdTransaksiBlockchain = keyModification.GetTxId()
	}

	return deliveryOrderResponse
}

func ToTransaksiResponse(ctx contractapi.TransactionContextInterface, keyModification *queryresult.KeyModification, transaksi *domain.Transaksi) *web.TransaksiResponse {
	transaksiResponse := &web.TransaksiResponse{
		Id:                     transaksi.Id,
		IdDeliveryOrder:        transaksi.IdDeliveryOrder,
		IdPetani:               transaksi.IdPetani,
		Nomor:                  transaksi.Nomor,
		TanggalPembuatan:       transaksi.TanggalPembuatan,
		StatusKoperasi:         transaksi.StatusKoperasi.String(),
		PesanKoperasi:          transaksi.PesanKoperasi,
		TanggalResponsKoperasi: transaksi.TanggalResponsKoperasi,
		StatusPks:              transaksi.StatusPks.String(),
		PesanPks:               transaksi.PesanPks,
		TanggalResponsPks:      transaksi.TanggalResponsPks,
		Status:                 transaksi.Status.String(),
		CreatedAt:              transaksi.CreatedAt,
		UpdatedAt:              transaksi.UpdatedAt,
	}

	if ctx != nil {
		transaksiResponse.IdTransaksiBlockchain = ctx.GetStub().GetTxID()
	}

	if keyModification != nil {
		transaksiResponse.IdTransaksiBlockchain = keyModification.GetTxId()
	}

	return transaksiResponse
}

func ToTransaksiItemResponse(ctx contractapi.TransactionContextInterface, keyModification *queryresult.KeyModification, transaksiItem *domain.TransaksiItem) *web.TransaksiItemResponse {
	transaksiItemResponse := &web.TransaksiItemResponse{
		Id:          transaksiItem.Id,
		IdTransaksi: transaksiItem.IdTransaksi,
		IdKebun:     transaksiItem.IdKebun,
		Kuantitas:   transaksiItem.Kuantitas,
		Harga:       transaksiItem.Harga,
		CreatedAt:   transaksiItem.CreatedAt,
		UpdatedAt:   transaksiItem.UpdatedAt,
	}

	if ctx != nil {
		transaksiItemResponse.IdTransaksiBlockchain = ctx.GetStub().GetTxID()
	}

	if keyModification != nil {
		transaksiItemResponse.IdTransaksiBlockchain = keyModification.GetTxId()
	}

	return transaksiItemResponse
}

func ToPengirimanResponse(ctx contractapi.TransactionContextInterface, keyModification *queryresult.KeyModification, pengiriman *domain.Pengiriman) *web.PengirimanResponse {
	pengirimanResponse := &web.PengirimanResponse{
		Id:                   pengiriman.Id,
		IdTransaksi:          pengiriman.IdTransaksi,
		Nomor:                pengiriman.Nomor,
		Tanggal:              pengiriman.Tanggal,
		NamaSopir:            pengiriman.NamaSopir,
		NomorTeleponSopir:    pengiriman.NomorTeleponSopir,
		NamaKendaraan:        pengiriman.NamaKendaraan,
		NomorPolisiKendaraan: pengiriman.NomorPolisiKendaraan,
		CreatedAt:            pengiriman.CreatedAt,
		UpdatedAt:            pengiriman.UpdatedAt,
	}

	if ctx != nil {
		pengirimanResponse.IdTransaksiBlockchain = ctx.GetStub().GetTxID()
	}

	if keyModification != nil {
		pengirimanResponse.IdTransaksiBlockchain = keyModification.GetTxId()
	}

	return pengirimanResponse
}

func ToPenerimaanResponse(ctx contractapi.TransactionContextInterface, keyModification *queryresult.KeyModification, penerimaan *domain.Penerimaan) *web.PenerimaanResponse {
	pengirimanResponse := &web.PenerimaanResponse{
		Id:          penerimaan.Id,
		IdTransaksi: penerimaan.IdTransaksi,
		Nomor:       penerimaan.Nomor,
		Tanggal:     penerimaan.Tanggal,
		Kuantitas:   penerimaan.Kuantitas,
		CreatedAt:   penerimaan.CreatedAt,
		UpdatedAt:   penerimaan.UpdatedAt,
	}

	if ctx != nil {
		pengirimanResponse.IdTransaksiBlockchain = ctx.GetStub().GetTxID()
	}

	if keyModification != nil {
		pengirimanResponse.IdTransaksiBlockchain = keyModification.GetTxId()
	}

	return pengirimanResponse
}

func ToPembayaranResponse(ctx contractapi.TransactionContextInterface, keyModification *queryresult.KeyModification, pembayaran *domain.Pembayaran) *web.PembayaranResponse {
	pembayaranResponse := &web.PembayaranResponse{
		Id:          pembayaran.Id,
		IdTransaksi: pembayaran.IdTransaksi,
		Nomor:       pembayaran.Nomor,
		Tanggal:     pembayaran.Tanggal,
		Jumlah:      pembayaran.Jumlah,
		HashBukti:   pembayaran.HashBukti,
		CreatedAt:   pembayaran.CreatedAt,
		UpdatedAt:   pembayaran.UpdatedAt,
	}

	if ctx != nil {
		pembayaranResponse.IdTransaksiBlockchain = ctx.GetStub().GetTxID()
	}

	if keyModification != nil {
		pembayaranResponse.IdTransaksiBlockchain = keyModification.GetTxId()
	}

	return pembayaranResponse
}
