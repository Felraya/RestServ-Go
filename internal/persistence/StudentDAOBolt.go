package persistence

import (
	"encoding/json"
	. "internal/entities"
	"strconv"
)

// DAO type
type studentDAOBolt struct {
	bucketName string
}

// DAO constructor
func NewStudentDAOBolt() *studentDAOBolt {

	dao := new(studentDAOBolt)
	dao.bucketName = "students"
	err := boltDB.CreateBucketIfNotExists(dao.bucketName)
	if err != nil {
		panic(err.Error())
	}
	return dao
}

// Fonctions
func (this *studentDAOBolt) FindAll() []Student {
	var res []Student
	var student Student

	resDB := boltDB.GetAll(this.bucketName)

	for _, studentDB := range resDB {
		err := json.Unmarshal([]byte(studentDB), &student)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, student)
	}

	return res
}

func (this *studentDAOBolt) Find(id int) *Student {
	res := new(Student)
	resDB := boltDB.Get(this.bucketName, strconv.Itoa(id))

	if resDB == "" {
		return nil
	} else {
		err := json.Unmarshal([]byte(resDB), &res)
		if err != nil {
			panic(err.Error())
		}
		return res
	}
}

func (this *studentDAOBolt) Exists(id int) bool {
	return boltDB.Get(this.bucketName, strconv.Itoa(id)) != ""
}

func (this *studentDAOBolt) Create(entity *Student) bool {
	if boltDB.Get(this.bucketName, strconv.Itoa(entity.Id)) == "" {
		return false
	} else {
		js, err := json.Marshal(entity)
		if err != nil {
			return false
		}
		boltDB.Put(this.bucketName, strconv.Itoa(entity.Id), string(js))
		return true
	}
}

func (this *studentDAOBolt) Delete(id int) bool {
	if boltDB.Get(this.bucketName, strconv.Itoa(id)) == "" {
		return false
	} else {
		boltDB.Delete(this.bucketName, strconv.Itoa(id))
		return true
	}
}

func (this *studentDAOBolt) Update(entity *Student) bool {
	if boltDB.Get(this.bucketName, strconv.Itoa(entity.Id)) == "" {
		return false
	} else {
		js, err := json.Marshal(entity)
		if err != nil {
			return false
		}
		boltDB.Put(this.bucketName, strconv.Itoa(entity.Id), string(js))
		return true
	}
}

func (this *studentDAOBolt) Save(entity *Student) {
	if boltDB.Get(this.bucketName, strconv.Itoa(entity.Id)) != "" {
		js, err := json.Marshal(entity)
		if err != nil {
			panic(err.Error())
		}
		boltDB.Put(this.bucketName, strconv.Itoa(entity.Id), string(js))
	}
}
