package main

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	application := gofr.New()
	db := application.DB()
	if db == nil {
		return
	}

	query := `
 	   CREATE TABLE IF NOT EXISTS product (
	   id int primary key not null AUTO_INCREMENT,
	   name varchar (100),
	   type varchar (100))	   
	`

	if application.Config.Get("DB_DIALECT") == "mysql" {
		query = `
		IF NOT EXISTS
	(  SELECT [name]
		FROM user.tables
      WHERE [name] = 'product'
   ) CREATE TABLE product (id int primary key not null AUTO_INCREMENT,
	   name varchar (100),
	   type varchar (100)),	   
	`
	}

	if _, err := db.Exec(query); err != nil {
		application.Logger.Errorf("got error sourcing the schema: ", err)
	}

	os.Exit(m.Run())
}

func TestIntegration(t *testing.T) {
	go main()
	time.Sleep(time.Second * 5)

	req, _ := request.NewMock(http.MethodGet, "http://localhost:8000/product", nil)
	c := http.Client{}

	resp, err := c.Do(req)
	if err != nil {
		t.Errorf("TEST Failed.\tHTTP request encountered Err: %v\n", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Failed.\tExpected %v\tGot %v\n", http.StatusOK, resp.StatusCode)
	}

	_ = resp.Body.Close()
}
