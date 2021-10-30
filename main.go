package main

import (
	"database/sql"
	"fmt"
	"github.com/mhaqiw/haqi-coba-golang/repository"
	"github.com/mhaqiw/haqi-coba-golang/usecase"
	"log"
	"time"

	"github.com/mhaqiw/haqi-coba-golang/handler"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/mhaqiw/haqi-coba-golang/util"
)

func main() {
	db := initDB()
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	e := echo.New()
	ps := repository.NewMProductStockRepository(db)
	p := repository.NewProductRepository()
	w := repository.NewWarehouseRepository()
	timeoutContext := time.Duration(10) * time.Second
	productStockUsecase := usecase.NewProductStockUsecase(ps,p,w,timeoutContext)
	handler.NewProductStockHandler(e,productStockUsecase)

	log.Fatal(e.Start("localhost:9090"))
}

func initDB() *sql.DB {
	postgresHost := util.MustHaveEnv("POSTGRES_HOST")
	postgresPort := util.MustHaveEnv("POSTGRES_PORT")
	postgresUser := util.MustHaveEnv("POSTGRES_USER")
	postgresPassword := util.MustHaveEnv("POSTGRES_PASSWORD")
	postgresDbname := util.MustHaveEnv("POSTGRES_DB")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		postgresHost, postgresPort, postgresUser, postgresPassword, postgresDbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
