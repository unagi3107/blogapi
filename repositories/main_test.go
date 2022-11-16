package repositories_test

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"

	_ "github.com/go-sql-driver/mysql"
)

var testDB *sql.DB

var (
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_NAME")
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func connectDB() error {
	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}
	return nil
}

func setup() error {
	if err := connectDB(); err != nil {
		return err
	}

	if err := cleanupDB(); err != nil {
		fmt.Println("cleanup", err)
		return err
	}

	if err := setupTestData(); err != nil {
		fmt.Println("setup", err)
		return err
	}

	return nil
}

func tearDown() {
	cleanupDB()
	defer testDB.Close()
}

func setupTestData() error {
	cmd := exec.Command("mysql", "-h", "127.0.0.1", "-u", "docker", "sampledb", "--password=docker", "-e", "source ./testdata/setup.sql")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func cleanupDB() error {
	cmd := exec.Command("mysql", "-h", "127.0.0.1", "-u", "docker", "sampledb", "--password=docker", "-e", "source ./testdata/cleanup.sql")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
