package persistence

import (
	. "internal/entities"
)

type LanguageDAO interface {
	FindAll() []Language
	Find(code string) *Language
	Exists(code string) bool
	Create(entity *Language) bool
	Delete(code string) bool
	Update(entity *Language) bool
	Save(entity *Language)
}

func GetLanguageDAO() LanguageDAO {
	return NewLanguageDAOBolt()
}
