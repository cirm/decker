package db

import (
	_ "github.com/lib/pq"
	"github.com/cirm/decker/env"
	"database/sql"
	"encoding/json"
	"log"
	"fmt"
)

func InitPg(c *env.AppContext) {
	var err error
	c.Db, err = sql.Open("postgres", "dbname=arco user=spark password=salasala host=postgres1.cydec port=5432 sslmode=disable")
	c.Db.SetMaxIdleConns(10)
	c.Db.SetMaxOpenConns(10)
	if err != nil {
		log.Fatal(err)
	}
	err = c.Db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Db connection opened.")
}

type JsonNullInt64 struct {
	sql.NullInt64
}

func (v JsonNullInt64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int64)
	} else {
		return json.Marshal(nil)
	}
}

func (v *JsonNullInt64) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *int64
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Int64 = *x
	} else {
		v.Valid = false
	}
	return nil
}

type JsonNullString struct {
	sql.NullString
}

func (v JsonNullString) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.String)
	} else {
		return json.Marshal(nil)
	}
}

func (v *JsonNullString) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *string
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.String = *x
	} else {
		v.Valid = false
	}
	return nil
}
