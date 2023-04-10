package perwalian

import (
	"fmt"
	model "github.com/alnoviantirs/perwalian/model"
	"github.com/alnoviantirs/perwalian/module"
	"testing"
)

func TestInsertPerwalian(t *testing.T) {
	time := model.Waktu{
		Jam : "10.00",
		Hari : "Rabu",
		Tanggal : "16 Maret 2023",
	}
	lokasi := "ULBI"
	walidosen := model.Dosen{
		Nama : "Indra riksa herlambang,. S.Tr.Kom., M.Kom., SFPC",
		Jabatan : "Wali Dosen D4 Logistik Bisnis",
	}
	biodata := model.Mahasiswa {
		Nama : "No Ae Seol",
		PhoneNumber : "6284564646124",
		Jurusan : "D4 Logistik Bisnis",
	}
	hasil:=module.InsertPerwalian(module.MongoConn, "perwalian", time, lokasi , walidosen, biodata )
	fmt.Println(hasil)
}
func TestInsertDosen(t *testing.T) {
	nama := "Rd. Nuraini Siti Fatonah, S.S., M.Hum."
	jabatan := "Wali Dosen D4 Teknik Informatika"
	hasil:=module.InsertDosen(module.MongoConn, "dosen", nama, jabatan)
	fmt.Println(hasil)
}
func TestInsertMahasiswa(t *testing.T) {
	nama := "Al Novianti Ramadhani"
	phone_number := "628456456211"
	jurusan := "D4 Teknik Informatika"
	hasil:=module.InsertMahasiswa(module.MongoConn, "mahasiswa", nama, phone_number, jurusan)
	fmt.Println(hasil)
}
func TestInsertWaktu(t *testing.T) {
	jam := "09.00"
	hari := "Rabu"
	tanggal := "29 Maret 2023"
	hasil:=module.InsertWaktu(module.MongoConn, "waktu", jam, hari, tanggal)
	fmt.Println(hasil)
}
func TestInsertLocation(t *testing.T) {
	nama_lokasi := "ULBI"
	alamat := "JL. Sarijadi no 54"
	hasil:=module.InsertLocation(module.MongoConn, "location", nama_lokasi, alamat)
	fmt.Println(hasil)
}
func TestInsertRuangan(t *testing.T) {
	lokasi_ruangan := "Ruang 102"
	hasil:=module.InsertRuangan(module.MongoConn, "ruangan", lokasi_ruangan)
	fmt.Println(hasil)
}

func TestGetMahasiswaFromNama(t *testing.T) {
	nama := "Al Novianti Ramadhani"
	biodata:=module.GetMahasiswaFromNama(module.MongoConn, "mahasiswa", nama)
	fmt.Println(biodata)
}
func TestGetDosenFromJabatan(t *testing.T) {
	jabatan := "Wali Dosen D4 Teknik Informatika"
	walidosen :=module.GetDosenFromJabatan(module.MongoConn, "dosen", jabatan)
	fmt.Println(walidosen)
}
func TestGetLocationFromAlamat(t *testing.T) {
	alamat := "JL. Sarijadi no 54"
	lokasi:=module.GetLocationFromAlamat(module.MongoConn, "location", alamat)
	fmt.Println(lokasi)
}
func TestGetRuanganFromLokasi_ruangan(t *testing.T) {
	lokasi_ruangan := "Ruang 102"
	ruangan:=module.GetRuanganFromLokasi_ruangan(module.MongoConn, "ruangan", lokasi_ruangan)
	fmt.Println(ruangan)
}
func TestWaktuFromJam(t *testing.T) {
	jam := "09.00"
	time:=module.GetWaktuFromJam(module.MongoConn, "waktu", jam)
	fmt.Println(time)
}
func TestPerwalianFromLokasi(t *testing.T) {
	lokasi := "ULBI"
	perwalian :=module.GetPerwalianFromLokasi(module.MongoConn, "perwalian", lokasi)
	fmt.Println(perwalian)
}

func TestGetAll(t *testing.T) {
	data :=module.GetAllPerwalian(module.MongoConn, "perwalian")
	fmt.Println(data)
}
func TestGetAllMahasiswa(t *testing.T) {
	data :=module.GetAllMahasiswa(module.MongoConn, "mahasiswa")
	fmt.Println(data)
}
func TestGetAllDosen(t *testing.T) {
	data :=module.GetAllDosen(module.MongoConn, "dosen")
	fmt.Println(data)
}
func TestGetAllLocation(t *testing.T) {
	data :=module.GetAllLocation(module.MongoConn, "location")
	fmt.Println(data)
}
func TestGetAllWaktu(t *testing.T) {
	data :=module.GetAllWaktu(module.MongoConn, "waktu")
	fmt.Println(data)
}
func TestGetAllRuangan(t *testing.T) {
	data :=module.GetAllRuangan(module.MongoConn, "ruangan")
	fmt.Println(data)
}