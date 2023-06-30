package module

import (
	"context"
	"fmt"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"github.com/alnoviantirs/perwalian/model"
	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func InsertPerwalian(db *mongo.Database, col string, time model.Waktu, lokasi string, walidosen model.Dosen, biodata model.Mahasiswa) (insertedID primitive.ObjectID, err error) {
	perwalian := bson.M{
		"time":   			 time,
		"lokasi":  		   lokasi,
		"walidosen":     walidosen,
		"biodata":			 biodata,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), perwalian)
	if err != nil {
		fmt.Printf("InsertPerwalian: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func InsertDosen(db *mongo.Database, col string, nama string, jabatan string) (insertedID primitive.ObjectID, err error) {
	dosen := bson.M{
		"nama":   			 nama,
		"jabatan":  		 jabatan,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), dosen)
	if err != nil {
		fmt.Printf("InsertDosen: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func InsertMahasiswa(db *mongo.Database, col string, nama string, phone_number string, jurusan string) (insertedID primitive.ObjectID, err error) {
	mahasiswa := bson.M{
		"nama":   			 nama,
		"phone_number":  phone_number,
		"jurusan":     		jurusan,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), mahasiswa)
	if err != nil {
		fmt.Printf("InsertMahasiswa: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func InsertWaktu(db *mongo.Database, col string, jam string, hari string, tanggal string) (insertedID primitive.ObjectID, err error) {
	waktu := bson.M{
		"jam":   			 jam,
		"hari":  		   hari,
		"tanggal":     tanggal,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), waktu)
	if err != nil {
		fmt.Printf("InsertWaktu %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func InsertLocation(db *mongo.Database, col string, nama_lokasi string, alamat string) (insertedID primitive.ObjectID, err error) {
	location:= bson.M{
		"nama_lokasi":    nama_lokasi,
		"alamat":  		  alamat,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), location)
	if err != nil {
		fmt.Printf("InsertLocation %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func InsertRuangan(db *mongo.Database, col string, lokasi_ruangan string) (insertedID primitive.ObjectID, err error) {
	ruangan:= bson.M{
		"lokasi_ruangan":    lokasi_ruangan,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), ruangan)
	if err != nil {
		fmt.Printf("InsertRuangan %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func UpdatePerwalian(db *mongo.Database, col string, id primitive.ObjectID, time model.Waktu, lokasi string, walidosen model.Dosen, biodata model.Mahasiswa) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"time":   			 time,
			"lokasi":  		   lokasi,
			"walidosen":     walidosen,
			"biodata":			 biodata,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdatePerwalian: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No document found with the specified ID")
		return
	}
	return nil
}
func UpdateDosen(db *mongo.Database, col string, id primitive.ObjectID, nama string, jabatan string) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"nama":   			 nama,
			"jabatan":  		 jabatan,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateDosen: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No document found with the specified ID")
		return
	}
	return nil
}

func UpdateMahasiswa(db *mongo.Database, col string, id primitive.ObjectID, nama string, phone_number string, jurusan string) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"nama":   			 nama,
			"phone_number":  phone_number,
			"jurusan":     		jurusan,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateMahasiswa: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No document found with the specified ID")
		return
	}
	return nil
}

func UpdateRuangan(db *mongo.Database, col string, id primitive.ObjectID, lokasi_ruangan string) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"lokasi_ruangan":    lokasi_ruangan,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateRuangan: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No document found with the specified ID")
		return
	}
	return nil
}


func DeletePerwalianByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	perwalian := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := perwalian.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}
func DeleteDosenByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	dosen := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := dosen.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}
func DeleteMahasiswaByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	mahasiswa := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := mahasiswa.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}
func DeleteRuanganByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	ruangan := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := ruangan.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

// func InsertPerwalian(db *mongo.Database, col string, time model.Waktu, lokasi string, walidosen model.Dosen, biodata model.Mahasiswa) (InsertedID interface{}) {
// 	var perwalian model.Perwalian
// 	perwalian.Time = time
// 	perwalian.Lokasi = lokasi
// 	perwalian.WaliDosen = walidosen
// 	perwalian.Biodata = biodata
// 	return InsertOneDoc( db, col, perwalian)
// }
// func InsertDosen(db *mongo.Database, col string, nama string, jabatan string) (InsertedID interface{}) {
// 	var dosen model.Dosen
// 	dosen.Nama = nama
// 	dosen.Jabatan = jabatan
// 	return InsertOneDoc( db, col, dosen)
// }
// func InsertMahasiswa(db *mongo.Database, col string, nama string, phone_number string, jurusan string) (InsertedID interface{}) {
// 	var mahasiswa model.Mahasiswa
// 	mahasiswa.Nama = nama
// 	mahasiswa.PhoneNumber = phone_number
// 	mahasiswa.Jurusan = jurusan
// 	return InsertOneDoc(db, col, mahasiswa)
// }
// func InsertWaktu(db *mongo.Database, col string, jam string, hari string, tanggal string) (InsertedID interface{}) {
// 	var waktu model.Waktu
// 	waktu.Jam = jam
// 	waktu.Hari = hari
// 	waktu.Tanggal = tanggal
// 	return InsertOneDoc(db, col, waktu)
// }
// func InsertLocation(db *mongo.Database, col string, nama_lokasi string, alamat string) (InsertedID interface{}) {
// 	var location model.Location
// 	location.Nama_lokasi = nama_lokasi
// 	location.Alamat = alamat
// 	return InsertOneDoc(db, col, location)
// }
// func InsertRuangan(db *mongo.Database, col string, lokasi_ruangan string) (InsertedID interface{}) {
// 	var ruangan model.Ruangan
// 	ruangan.Lokasi_ruangan = lokasi_ruangan
// 	return InsertOneDoc(db, col, ruangan)
// }

func GetMahasiswaFromNama(db *mongo.Database, col string, nama string) (mhs model.Mahasiswa) {
	mahasiswa := db.Collection(col)
	filter := bson.M{"nama": nama}
	err := mahasiswa.FindOne(context.TODO(), filter).Decode(&mhs)
	if err != nil {
		fmt.Printf("getMahasiswaFromNama: %v\n", err)
	}
	return mhs
}
func GetPerwalianFromID(_id primitive.ObjectID, db *mongo.Database, col string) (mahasiswa model.Perwalian, errs error) {
	perwalian := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := perwalian.FindOne(context.TODO(), filter).Decode(&mahasiswa)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return mahasiswa, fmt.Errorf("no data found for ID %s", _id)
		}
		return mahasiswa, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return mahasiswa, nil
}
func GetDosenFromID(_id primitive.ObjectID, db *mongo.Database, col string) (dosenid model.Dosen, errs error) {
	dosen := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := dosen.FindOne(context.TODO(), filter).Decode(&dosenid)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return dosenid, fmt.Errorf("no data found for ID %s", _id)
		}
		return dosenid, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return dosenid, nil
}
func GetMahasiswaFromID(_id primitive.ObjectID, db *mongo.Database, col string) (mahasiswaid model.Mahasiswa, errs error) {
	mahasiswa := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := mahasiswa.FindOne(context.TODO(), filter).Decode(&mahasiswaid)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return mahasiswaid, fmt.Errorf("no data found for ID %s", _id)
		}
		return mahasiswaid, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return mahasiswaid, nil
}
func GetRuanganFromID(_id primitive.ObjectID, db *mongo.Database, col string) (ruanganid model.Ruangan, errs error) {
	ruangan := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := ruangan.FindOne(context.TODO(), filter).Decode(&ruanganid)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ruanganid, fmt.Errorf("no data found for ID %s", _id)
		}
		return ruanganid, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return ruanganid, nil
}
func GetNohpFromNama(db *mongo.Database, col string, phone_number string) (phn model.Mahasiswa) {
	mahasiswa := db.Collection(col)
	filter := bson.M{"phone_number": phone_number}
	err := mahasiswa.FindOne(context.TODO(), filter).Decode(&phn)
	if err != nil {
		fmt.Printf("getNohpFromNama: %v\n", err)
	}
	return phn
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
func GetPerwalianFromLokasi(db *mongo.Database, col string, lokasi string) (per model.Perwalian) {
	perwalian := db.Collection(col)
	filter := bson.M{"lokasi": lokasi}
	err := perwalian.FindOne(context.TODO(), filter).Decode(&per)
	if err != nil {
		fmt.Printf("getPerwalianFromLokasi: %v\n", err)
	}
	return per
}

func GetAllPerwalian(db *mongo.Database, col string) (data []model.Perwalian) {
	perwalian := db.Collection(col)
	filter := bson.M{}
	cursor, err := perwalian.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func GetAllMahasiswa(db *mongo.Database, col string) (data []model.Mahasiswa) {
	perwalian := db.Collection(col)
	filter := bson.M{}
	cursor, err := perwalian.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func GetAllDosen(db *mongo.Database, col string) (data []model.Dosen) {
	perwalian := db.Collection(col)
	filter := bson.M{}
	cursor, err := perwalian.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func GetAllLocation(db *mongo.Database, col string) (data []model.Location) {
	perwalian := db.Collection(col)
	filter := bson.M{}
	cursor, err := perwalian.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func GetAllWaktu(db *mongo.Database, col string) (data []model.Waktu) {
	perwalian := db.Collection(col)
	filter := bson.M{}
	cursor, err := perwalian.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
func GetAllRuangan(db *mongo.Database, col string) (data []model.Ruangan) {
	perwalian := db.Collection(col)
	filter := bson.M{}
	cursor, err := perwalian.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

//login
func LoginAdmin(db *mongo.Database, col string, username string, password string) (authenticated bool, err error) {
	filter := bson.M{
		"username": username,
		"password": password,
	}

	result, err := db.Collection(col).CountDocuments(context.Background(), filter)
	if err != nil {
		fmt.Printf("LoginAdmin: %v\n", err)
		return false, err
	}

	if result == 1 {
		return true, nil
	}

	return false, nil
}

//login

// //login
func InsertAdmin(db *mongo.Database, col string, username string, password string) (insertedID primitive.ObjectID, err error) {
	admin := bson.M{
		"username":	username,
		"password": password,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), admin)
	if err != nil {
		fmt.Printf("InsertAdmin: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}