package main

import (
	"os"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"io/ioutil"
	 "bytes"
)

//this function post messages to slack https://postman-assignment.slack.com/messages/CDUBBCUV7/
func post(url string, jsonData string) string {
	var jsonStr = []byte(jsonData)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}


func main() {

	//yes i understand the consequences of open sourcing this url
  url := "https://hooks.slack.com/services/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"

	fmt.Println("Opening connection")
	post(url, " {\"text\":\"Opening connection\"} " )
	db, err := sql.Open("mysql", "root:"+os.Getenv("mySQLPassword")+"@tcp("+os.Getenv("mySQLIPAddress")+":"+ os.Getenv("mySQLIPPort") +")/")

	if err != nil {
		fmt.Println("Error: ", err)
		panic(err.Error()) // Just for example purpose. We should use proper error handling instead of panic
	}
	fmt.Println("Connection opened successfully!")
	post(url, " {\"text\":\"Connection opened successfully!\"} " )

	// Last thing to do, close connection
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Execute the query
	var (
		Id int
		LastName string
		FirstName string
		Address string
		City string
	)

	// DROP database DATABASE_NAME
	// _, err = db.Exec("DROP database IF EXIST testdb")
	// if err!=nil{
	// 	log.Fatal(err)
	// }

	//// Create Database if not exist DROP database DATABASE_NAME
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS testdb")
	if err!=nil{
		log.Fatal(err)
	}

	//// Use our database
	_, err = db.Exec("USE testdb")
	if err!=nil{
		log.Fatal(err)
	}

	//// Create Table if not exists
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Persons (PersonID int NOT NULL AUTO_INCREMENT, LastName varchar(255), FirstName varchar(255), Address varchar(255), City varchar(255), PRIMARY KEY (PersonID));")
	if err!=nil{
		log.Fatal(err)
	}


	///////////// Read Operation ///////////////////
	rows, err := db.Query("SELECT * FROM Persons")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer rows.Close()
	// Parametrize
	//rows, err := rowsQuery.Query(3)
	for rows.Next() {
		err := rows.Scan(&Id,  &LastName, &FirstName, &Address, &City)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(Id, LastName, FirstName, Address, City)
	}

	//////////// Insertion ////////////////////////
	stmt, err := db.Prepare("INSERT INTO Persons(LastName, FirstName, Address, City) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec("Lavania", "Susmit", "Rajasthan", "Jaipur")
	if err != nil {
		log.Fatal(err)
	}

	// META
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	if lastId >= 7{
		fmt.Println("Database Alert Max Autoincrement value reached")
		post(url, " {\"text\":\"Database Alert Max Autoincrement value reached\"} " )
	}


	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)


	//post("https://hooks.slack.com/services/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", " {\"text\":\"Hello, World! from golang\"} " )

}
