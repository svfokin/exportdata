package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Username string
	DBName   string
	SSLMode  string
	Password string
}

type Abonent struct {
	Id                sql.NullString
	Ls_reg            sql.NullString
	Uuid              sql.NullString
	Ncounter          sql.NullString
	Ls_gas            sql.NullString
	Id_ais            sql.NullString
	Database_name     sql.NullString
	Typecounter       sql.NullString
	Street_uuid       sql.NullString
	Fio               sql.NullString
	Adress            sql.NullString
	Id_turg           sql.NullString
	Id_rajon          sql.NullString
	Id_filial         sql.NullString
	Legal_org         sql.NullString
	Verification_date sql.NullString
	Ncounter_real     sql.NullString
	Equipment_uuid    sql.NullString
	Working           sql.NullString
	Date_remote       sql.NullString
	Date_amount       sql.NullString
	Amount            sql.NullString
	Equipment_name    sql.NullString
	Department_uuid   sql.NullString
	Update_date       sql.NullString
}

type AbonentStr struct {
	Id                string
	Ls_reg            string
	Uuid              string
	Ncounter          string
	Ls_gas            string
	Id_ais            string
	Database_name     string
	Typecounter       string
	Street_uuid       string
	Fio               string
	Adress            string
	Id_turg           string
	Id_rajon          string
	Id_filial         string
	Legal_org         string
	Verification_date string
	Ncounter_real     string
	Equipment_uuid    string
	Working           string
	Date_remote       string
	Date_amount       string
	Amount            string
	Equipment_name    string
	Department_uuid   string
	Update_date       string
}

// Функция преобразования из NullString
func SqlString(str sql.NullString) string {
	if str.Valid {
		return str.String
	}
	return "_"
}

func NewSelectDB(db *sqlx.DB) ([]Abonent, error) {
	abonent := []Abonent{}
	err := db.Select(&abonent, "SELECT * FROM abonents LIMIT 10")

	if err != nil {
		return nil, err
	}

	return abonent, nil
}

func NewQueryDB(db *sqlx.DB) error {

	rows, err := db.Queryx("SELECT * FROM abonents LIMIT 10")

	if err != nil {
		return err
	}

	var a Abonent
	var as AbonentStr

	/*
		r := reflect.ValueOf(a)

		for i := 0; i < r.NumField(); i++ {
			fmt.Printf("%s ", r.Type().Field(i).Name)
		}
	*/

	kol := 0

	for rows.Next() {
		err = rows.StructScan(&a)
		if err != nil {
			return err
		}
		kol++
		as.Id = SqlString(a.Id)
		as.Ls_reg = SqlString(a.Ls_reg)
		as.Uuid = SqlString(a.Uuid)
		as.Ncounter = SqlString(a.Ncounter)
		as.Ls_gas = SqlString(a.Ls_gas)
		as.Id_ais = SqlString(a.Id_ais)
		as.Database_name = SqlString(a.Database_name)
		as.Typecounter = SqlString(a.Typecounter)
		as.Street_uuid = SqlString(a.Street_uuid)
		as.Fio = SqlString(a.Fio)
		as.Adress = SqlString(a.Adress)
		as.Id_turg = SqlString(a.Id_turg)
		as.Id_rajon = SqlString(a.Id_rajon)
		as.Id_filial = SqlString(a.Id_filial)
		as.Legal_org = SqlString(a.Legal_org)
		as.Verification_date = SqlString(a.Verification_date)
		as.Ncounter_real = SqlString(a.Ncounter_real)
		as.Equipment_uuid = SqlString(a.Equipment_uuid)
		as.Working = SqlString(a.Working)
		as.Date_remote = SqlString(a.Date_remote)
		as.Date_amount = SqlString(a.Date_amount)
		as.Amount = SqlString(a.Amount)
		as.Equipment_name = SqlString(a.Equipment_name)
		as.Department_uuid = SqlString(a.Department_uuid)
		as.Update_date = SqlString(a.Update_date)
		fmt.Println("- - - - - - - -")
		fmt.Println(as)
	}

	aj, err := json.Marshal(as)
	if err != nil {
		return err
	}

	var au AbonentStr
	err = json.Unmarshal(aj, &au)
	if err != nil {
		return err
	}

	//	fmt.Printf("%#v\n", a)
	fmt.Println(kol)
	fmt.Println(aj)
	fmt.Println(au)
	fmt.Println(string(aj))

	return nil
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
