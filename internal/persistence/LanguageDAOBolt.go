package persistence

import (
	"encoding/json"
	. "internal/entities"
)

// DAO type
type languageDAOBolt struct {
	bucketName string
}

// DAO constructor
func NewLanguageDAOBolt() *languageDAOBolt {
	dao := new(languageDAOBolt)
	dao.bucketName = "languages"
	err := boltDB.CreateBucketIfNotExists(dao.bucketName)
	if err != nil {
		panic(err.Error())
	}
	return dao
}

// Fonctions
func (this *languageDAOBolt) FindAll() []Language {
	var res []Language
	var language Language

	resDB := boltDB.GetAll(this.bucketName)

	for _, languageDB := range resDB {
		err := json.Unmarshal([]byte(languageDB), &language)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, language)
	}

	return res
}

func (this *languageDAOBolt) Find(code string) *Language {
	res := new(Language)
	resDB := boltDB.Get(this.bucketName, code)

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

func (this *languageDAOBolt) Exists(code string) bool {
	return boltDB.Get(this.bucketName, code) != ""
}

func (this *languageDAOBolt) Create(entity *Language) bool {
	if boltDB.Get(this.bucketName, entity.Code) != "" {
		return false
	} else {
		js, err := json.Marshal(entity)
		if err != nil {
			return false
		}
		boltDB.Put(this.bucketName, entity.Code, string(js))
		return true
	}
}

func (this *languageDAOBolt) Delete(code string) bool {
	if boltDB.Get(this.bucketName, code) == "" {
		return false
	} else {
		boltDB.Delete(this.bucketName, code)
		return true
	}
}

func (this *languageDAOBolt) Update(entity *Language) bool {
	if boltDB.Get(this.bucketName, entity.Code) == "" {
		return false
	} else {
		js, err := json.Marshal(entity)
		if err != nil {
			return false
		}
		boltDB.Put(this.bucketName, entity.Code, string(js))
		return true
	}
}

func (this *languageDAOBolt) Save(entity *Language) {
	if boltDB.Get(this.bucketName, entity.Code) != "" {
		js, err := json.Marshal(entity)
		if err != nil {
			panic(err.Error())
		}
		boltDB.Put(this.bucketName, entity.Code, string(js))
	}
}
