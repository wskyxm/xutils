package xdao

import "fmt"

const Sqlite = "sqlite"
const Pgsql  = "pgsql"
const Mysql  = "mysql"

type DsnParams struct {
	Type		string
	Host		string
	Port		string
	User		string
	Password	string
	DBName		string
}

func (s *DsnParams)String() string {
	switch s.Type {
	case Pgsql: return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", s.Host, s.Port, s.User, s.Password, s.DBName)
	case Mysql: return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", s.User, s.Password, s.Host, s.Port, s.DBName)
	default: return s.DBName
	}
}