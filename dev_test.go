package perwalian

import (
	"fmt"
	model "github.com/alnoviantirs/perwalian/model"
	"github.com/alnoviantirs/perwalian/module"
	"testing"
)

func TestInsertPerwalian(t *testing.T) {
	time := model.Waktu{
		Jam : "09.00",
		Hari : "Selasa",
		Tanggal : "15 Maret 2023",
	}
	lokasi := "ULBI"
	walidosen := model.Dosen{
		Nama : "Rd. Nuraini Siti Fatonah, S.S., M.Hum. ",
		Jabatan : "Wali Dosen D4 Teknik Informatika",
	}
	biodata := model.Mahasiswa {
		Nama : "Al Novianti Ramadhani",
		PhoneNumber : "628456456211",
		Jurusan : "D4 Teknik Informatika",
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

