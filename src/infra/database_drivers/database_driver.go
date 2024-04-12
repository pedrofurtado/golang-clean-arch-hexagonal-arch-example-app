package database_drivers

import (
	"fmt"
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type SqlDatabase struct {
	DB *sql.DB
}

func (db SqlDatabase) Close() error {
	fmt.Println("SqlDatabase::Close")
	return db.DB.Close()
}

func (db SqlDatabase) GetDB() *sql.DB {
	fmt.Println("SqlDatabase::GetDB")
	return db.DB
}

func Init() (SqlDatabase) {
	return SqlDatabase{
		newDatabase(),
	}
}

func newDatabase() (*sql.DB) {
	databaseDriver := os.Getenv("APP_ADAPTER_DATABASE_DRIVER")
	connectionUrl := ""

	switch databaseDriver {
		case "mysql":
			connectionUrl = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_NAME"))
		default:
			err := "Error Must be defined a database driver!"
			fmt.Println(err)
			panic(err)
	}

	db, err := sql.Open(databaseDriver, connectionUrl)
	if err != nil {
		fmt.Println("Error on database connection")
		panic(err)
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Error on database ping")
		panic(err)
	}

	fmt.Println("Database connection established successfully")
	return db
}
