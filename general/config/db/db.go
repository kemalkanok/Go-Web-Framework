package db

func NewDb() *db {
m := new(db)
m.port := 8080
return m
}
