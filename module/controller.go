package module

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"github.com/alnoviantirs/perwalian/model"
	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "perwalian_db",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

// func MongoConnect(dbname string) (db *mongo.Database) {
// 	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
// 	if err != nil {
// 		fmt.Printf("MongoConnect: %v\n", err)
// 	}
// 	return client.Database(dbname)
// }

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertPerwalian(db *mongo.Database, col string, time model.Waktu, lokasi string, walidosen model.Dosen, biodata model.Mahasiswa) (InsertedID interface{}) {
	var perwalian model.Perwalian
	perwalian.Time = time
	perwalian.Lokasi = lokasi
	perwalian.WaliDosen = walidosen
	perwalian.Biodata = biodata
	return InsertOneDoc( db, col, perwalian)
}
func InsertDosen(db *mongo.Database, col string, nama string, jabatan string) (InsertedID interface{}) {
	var dosen model.Dosen
	dosen.Nama = nama
	dosen.Jabatan = jabatan
	return InsertOneDoc( db, col, dosen)
}
func InsertMahasiswa(db *mongo.Database, col string, nama string, phone_number string, jurusan string) (InsertedID interface{}) {
	var mahasiswa model.Mahasiswa
	mahasiswa.Nama = nama
	mahasiswa.PhoneNumber = phone_number
	mahasiswa.Jurusan = jurusan
	return InsertOneDoc(db, col, mahasiswa)
}
func InsertWaktu(db *mongo.Database, col string, jam string, hari string, tanggal string) (InsertedID interface{}) {
	var waktu model.Waktu
	waktu.Jam = jam
	waktu.Hari = hari
	waktu.Tanggal = tanggal
	return InsertOneDoc(db, col, waktu)
}
func InsertLocation(db *mongo.Database, col string, nama_lokasi string, alamat string) (InsertedID interface{}) {
	var location model.Location
	location.Nama_lokasi = nama_lokasi
	location.Alamat = alamat
	return InsertOneDoc(db, col, location)
}
func InsertRuangan(db *mongo.Database, col string, lokasi_ruangan string) (InsertedID interface{}) {
	var ruangan model.Ruangan
	ruangan.Lokasi_ruangan = lokasi_ruangan
	return InsertOneDoc(db, col, ruangan)
}

func GetMahasiswaFromNama(db *mongo.Database, col string, nama string) (mhs model.Mahasiswa) {
	mahasiswa := db.Collection(col)
	filter := bson.M{"nama": nama}
	err := mahasiswa.FindOne(context.TODO(), filter).Decode(&mhs)
	if err != nil {
		fmt.Printf("getMahasiswaFromNama: %v\n", err)
	}
	return mhs
}
func GetDosenFromJabatan(db *mongo.Database, col string, jabatan string) (dsn model.Dosen) {
	dosen := db.Collection(col)
	filter := bson.M{"jabatan": jabatan}
	err := dosen.FindOne(context.TODO(), filter).Decode(&dsn)
	if err != nil {
		fmt.Printf("getDosenFromJabatan: %v\n", err)
	}
	return dsn
}
func GetLocationFromAlamat(db *mongo.Database, col string, alamat string) (loc model.Location) {
	location := db.Collection(col)
	filter := bson.M{"alamat": alamat}
	err := location.FindOne(context.TODO(), filter).Decode(&loc)
	if err != nil {
		fmt.Printf("getLocationFromAlamat: %v\n", err)
	}
	return loc
}
func GetRuanganFromLokasi_ruangan(db *mongo.Database, col string, lokasi_ruangan string) (rgn model.Ruangan) {
	ruangan := db.Collection(col)
	filter := bson.M{"lokasi_ruangan": lokasi_ruangan}
	err := ruangan.FindOne(context.TODO(), filter).Decode(&rgn)
	if err != nil {
		fmt.Printf("getRuanganFromLokasi_ruangan: %v\n", err)
	}
	return rgn
}
func GetWaktuFromJam(db *mongo.Database, col string, jam string) (wkt model.Waktu) {
	waktu := db.Collection(col)
	filter := bson.M{"jam": jam}
	err := waktu.FindOne(context.TODO(), filter).Decode(&wkt)
	if err != nil {
		fmt.Printf("getWaktuFromJam: %v\n", err)
	}
	return wkt
}
