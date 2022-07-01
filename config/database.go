package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Threx-code/go-api/utils"
	_ "github.com/go-sql-driver/mysql"
)

var server, err = utils.LoadConfig("../../")

var (
	connection *sql.DB
)

// const (
// 	username     = "go_user"
// 	password     = "go_password"
// 	hostname     = "127.0.0.1:3306"
// 	dbname       = "go_api"
// 	root         = "root"
// 	rootPassword = ""
// )

func dns(dbName string) string {
	if err != nil {
		log.Fatal(err)
	}

	if dbName == "" {
		return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", server.DBRoot, server.RootPassword, server.DBHost, "")
	}
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", server.DBUser, server.DBPassword, server.DBHost, server.DBName)
}

func CreateDatabase() {
	// open a connection to the databse
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", dns(""))
	if err != nil {
		log.Printf("Error %s when opening DB\n", err.Error())
		return
	}
	defer db.Close()

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+server.DBName)
	if err != nil {
		log.Printf("%s when creating database\n", err.Error())
		return
	}

	// create user with priviledges
	_, err = db.ExecContext(ctx, "CREATE USER IF NOT EXISTS '"+server.DBUser+"' @'localhost' IDENTIFIED BY '"+server.DBPassword+"'")
	if err != nil {
		log.Printf("%s error creating user", err.Error())
		return
	}

	_, err = db.ExecContext(ctx, "GRANT ALL PRIVILEGES ON "+server.DBName+".* TO '"+server.DBName+"'@'localhost'")
	if err != nil {
		log.Printf("%s when creating user", err.Error())
		return
	}

	_, err = db.ExecContext(ctx, "FLUSH PRIVILEGES")
	if err != nil {
		log.Printf("%s when flushing", err.Error())
		return
	}

	_, err = res.RowsAffected()
	if err != nil {
		log.Printf("error %s when fetching rows\n", err.Error())
		return
	}

	db.Close()
}

func Connect() {
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", dns(server.DBName))
	if err != nil {
		log.Printf("%s error when connecting to database "+server.DBName, err.Error())
		return
	}

	// setting the connection pool limit that are allowed within our application
	db.SetMaxOpenConns(20)
	//setting the limit of idle connection
	db.SetMaxIdleConns(20)
	// close the idle connections pool
	db.SetConnMaxLifetime(time.Minute * 5)

	// ping the database server
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("%s error pinging database ", err.Error())
		return
	}

	connection = db
}

func GetDB() *sql.DB {
	return connection
}

func RunMigration(tableName string, query string) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	Connect()
	db := GetDB()
	res, err := db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("%s error creating table "+tableName, err.Error())
		return
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("%s error getting affected rows ", err.Error())
	}

	log.Printf("rows affected %d", rows)
}
