package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type 	Dosen struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama        string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Jabatan     string             `bson:"jabatan,omitempty" json:"jabatan,omitempty"`
}

type Mahasiswa struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	PhoneNumber string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Jurusan      string             `bson:"jurusan,omitempty" json:"jurusan,omitempty"`
}

type Perwalian struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Time         Waktu             `bson:"time,omitempty" json:"time,omitempty"`	
	Lokasi       string             `bson:"lokasi,omitempty" json:"lokasi,omitempty"`
	WaliDosen    Dosen              `bson:"walidosen,omitempty" json:"walidosen,omitempty"`
	Biodata      Mahasiswa          `bson:"biodata,omitempty" json:"biodata,omitempty"`
	Lokasi_ruangan      Ruangan          `bson:"lokasi_ruangan,omitempty" json:"lokasi_ruangan,omitempty"`
}

type Waktu struct {
	Jam      string   `bson:"jam,omitempty" json:"jam,omitempty"`	
	Hari     string   `bson:"hari,omitempty" json:"hari,omitempty"`
	Tanggal  string   `bson:"tanggal,omitempty" json:"tanggal,omitempty"`
}

type Location struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_lokasi    string             `bson:"nama_lokasi,omitempty" json:"nama_lokasi,omitempty"`
	Alamat         string             `bson:"alamat,omitempty" json:"alamat,omitempty"`
}
type Ruangan struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Lokasi_ruangan    string          `bson:"lokasi_ruangan,omitempty" json:"lokasi_ruangan,omitempty"`
}

type Admin struct{
	ID              primitive.ObjectID 	`bson:"_id,omitempty" json:"_id,omitempty"`
	Username 		string             	`bson:"username,omitempty" json:"username,omitempty"`
	Password        string          	`bson:"password,omitempty" json:"password,omitempty"`
}