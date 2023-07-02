package perwalian

import (
	"fmt"
	"testing"

	model "github.com/alnoviantirs/perwalian/model"
	"github.com/alnoviantirs/perwalian/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertPerwalian(t *testing.T) {
	time := model.Waktu{
		Jam : "10.00",
		Hari : "Rabu",
		Tanggal : "16 Maret 2023",
	}
	walidosen := model.Dosen{
		Nama : "Indra riksa herlambang,. S.Tr.Kom., M.Kom., SFPC",
		Jabatan : "Wali Dosen D4 Logistik Bisnis",
	}
	biodata := model.Mahasiswa {
		Nama : "No Ae Seol",
		PhoneNumber : "6284564646124",
		Jurusan : "D4 Logistik Bisnis",
	}
	ruangan := model.Ruangan{
		Lokasi_ruangan : "Ruang 101",
	}
	hasil, err:=module.InsertPerwalian(module.MongoConn, "perwalian", time, walidosen, biodata, ruangan)
	if err != nil {
		t.Errorf("Data berhasil disimpan dengan id %s", hasil.Hex())
	}
	fmt.Println(hasil)
}
func TestInsertDosen(t *testing.T) {
	nama := "Rd. Nuraini Siti Fatonah, S.S., M.Hum."
	jabatan := "Wali Dosen D4 Teknik Informatika"
	hasil, err :=module.InsertDosen(module.MongoConn, "dosen", nama, jabatan)
	if err != nil {
		t.Errorf("Data berhasil disimpan dengan id %s", hasil.Hex())
	}
	fmt.Println(hasil)
}
func TestInsertMahasiswa(t *testing.T) {
	nama := "Al Novianti Ramadhani"
	phone_number := "628456456211"
	jurusan := "D4 Teknik Informatika"
	hasil, err:=module.InsertMahasiswa(module.MongoConn, "mahasiswa", nama, phone_number, jurusan)
	if err != nil {
		t.Errorf("Data berhasil disimpan dengan id %s", hasil.Hex())
	}
	fmt.Println(hasil)
}
func TestInsertWaktu(t *testing.T) {
	jam := "09.00"
	hari := "Rabu"
	tanggal := "29 Maret 2023"
	hasil, err :=module.InsertWaktu(module.MongoConn, "waktu", jam, hari, tanggal)
	if err != nil {
		t.Errorf("Data berhasil disimpan dengan id %s", hasil.Hex())
	}
	fmt.Println(hasil)
}
func TestInsertLocation(t *testing.T) {
	nama_lokasi := "ULBI"
	alamat := "JL. Sarijadi no 54"
	hasil, err :=module.InsertLocation(module.MongoConn, "location", nama_lokasi, alamat)
	if err != nil {
		t.Errorf("Data berhasil disimpan dengan id %s", hasil.Hex())
	}
	fmt.Println(hasil)
}
func TestInsertRuangan(t *testing.T) {
	lokasi_ruangan := "Ruang 102"
	hasil, err:=module.InsertRuangan(module.MongoConn, "ruangan", lokasi_ruangan)
	if err != nil {
		t.Errorf("Data berhasil disimpan dengan id %s", hasil.Hex())
	}
	fmt.Println(hasil)
}

// func TestGetPerwalianFromID(t *testing.T) {
// 	_id := "6423f7398e67b4e741878f01"
// 	data:=module.GetPerwalianFromID(module.MongoConn, "perwalian", _id)
// 	fmt.Println(data)
// }

func TestGetPerwalianFromID(t *testing.T) {
	id := "642402a0e04e68c3f6724704"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	biodata, err := module.GetPerwalianFromID(objectID, module.MongoConn, "perwalian")
	if err != nil {
		t.Fatalf("error calling GetPerwalianFromID: %v", err)
	}
	fmt.Println(biodata)
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
func TestGetNohpFromNama(t *testing.T) {
	phone_number := "628456456266"
	mahasiswa:=module.GetNohpFromNama(module.MongoConn, "mahasiswa", phone_number)
	fmt.Println(mahasiswa)
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

func TestDeletePerwalianByID(t *testing.T) {
	id := "648adad330628a7f140274b1" // ID data yang ingin dihapus
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeletePerwalianByID(objectID, module.MongoConn, "perwalian")
	if err != nil {
		t.Fatalf("error calling DeletePerwalianByID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan GetPresensiFromID
	_, err = module.GetPerwalianFromID(objectID, module.MongoConn, "perwalian")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}

func TestDeleteDosenByID(t *testing.T) {
	id := "6423f66dacc4d0a3ca435d84" // ID data yang ingin dihapus
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeleteDosenByID(objectID, module.MongoConn, "dosen")
	if err != nil {
		t.Fatalf("error calling DeleteDosenByID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan GetPresensiFromID
	_, err = module.GetDosenFromID(objectID, module.MongoConn, "dosen")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}

func TestDeleteMahasiswaByID(t *testing.T) {
	id := "6423f66dacc4d0a3ca435d85" // ID data yang ingin dihapus
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeleteMahasiswaByID(objectID, module.MongoConn, "mahasiswa")
	if err != nil {
		t.Fatalf("error calling DeleteMahasiswaByID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan GetPresensiFromID
	_, err = module.GetMahasiswaFromID(objectID, module.MongoConn, "mahasiswa")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}

func TestDeleteRuanganByID(t *testing.T) {
	id := "6423f66dacc4d0a3ca435d88" // ID data yang ingin dihapus
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeleteRuanganByID(objectID, module.MongoConn, "ruangan")
	if err != nil {
		t.Fatalf("error calling DeleteRuanganByID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan GetPresensiFromID
	_, err = module.GetRuanganFromID(objectID, module.MongoConn, "ruangan")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}

//login
func TestLoginAdmin(t *testing.T) {
	username := "admin"
	password := "admin1234"

	authenticated, err := module.LoginAdmin(module.MongoConn, "admin", username, password)
	if err != nil {
		t.Errorf("Error authenticating admin: %v", err)
	}

	if authenticated {
		fmt.Println("Admin authenticated successfully")
	} else {
		t.Errorf("Admin authentication failed")
	}
}
//login

func TestInsertAdmin(t *testing.T) {
	username := "admin"
	password := "admin1234"

	insertedID, err := module.InsertAdmin(module.MongoConn, "admin", username, password)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}