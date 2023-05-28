package controller

import (
	"context"
	"encoding/json"
	"example/go/mongodb/model"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "user"
const colName = "students"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)
	fmt.Println("ClintOptom TYPE:", reflect.TypeOf(clientOption))

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connect sucessful")
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection is ready")
}
func getAllStudent() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var students []primitive.M
	for cur.Next(context.Background()) {
		var student primitive.M
		if err = cur.Decode(&student); err != nil {
			log.Fatal(err)
		} else {
			students = append(students, student)
		}
	}
	defer cur.Close(context.Background())
	return students
}
func getOneStudent(studentId string) primitive.M {
	id, _ := primitive.ObjectIDFromHex(studentId)
	filter := bson.M{"_id": id}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	var student primitive.M
	if err = cur.Decode(&student); err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	return student
}
func createOneStudent(student model.Student) interface{} {
	inesrted, err := collection.InsertOne(context.Background(), student)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Insered:", student, "with ", inesrted.InsertedID)
	return inesrted.InsertedID
}

func updateStudent(studentId string, student model.Student) int64 {
	id, _ := primitive.ObjectIDFromHex(studentId)
	filter := bson.M{"_id": id}
	//warning this will be eror
	update := bson.M{"$set": bson.M{"firstname": student.Name, "lastname": student.SurName, "age": student.Age}}
	result, err := collection.UpdateMany(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return result.ModifiedCount
}

func deleteStudent(studentId string) int64 {
	id, _ := primitive.ObjectIDFromHex(studentId)
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	return result.DeletedCount
}

func clearStudent() int64 {
	result, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	return result.DeletedCount
}

func GetAllStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	ls := getAllStudent()

	json.NewEncoder(w).Encode(ls)
}

func GetOneStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	params := mux.Vars(r)
	data := getOneStudent(params["id"])

	json.NewEncoder(w).Encode(data)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PATCH")
	params := mux.Vars(r)
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}
	var student model.Student
	_ = json.NewDecoder(r.Body).Decode(&student)
	if !student.IsValid() {
		data, _ := json.MarshalIndent(student, "", "\t")
		dataString := string(data)
		json.NewEncoder(w).Encode("You send wrong data\n" + dataString)
		return
	}

	updateCount := updateStudent(params["id"], student)

	json.NewEncoder(w).Encode(updateCount)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	data := deleteStudent(params["id"])

	json.NewEncoder(w).Encode(data)
}

func ClearStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	data := clearStudent()

	json.NewEncoder(w).Encode(data)
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}
	var student model.Student
	_ = json.NewDecoder(r.Body).Decode(&student)
	if !student.IsValid() {
		data, _ := json.MarshalIndent(student, "", "\t")
		dataString := string(data)
		json.NewEncoder(w).Encode("You send wrong data \n" + dataString)
		return
	}

	data := createOneStudent(student)

	json.NewEncoder(w).Encode(data)
}
