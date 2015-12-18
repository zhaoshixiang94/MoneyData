package main

import (
	"encoding/json"
	"fmt"
	r "gopkg.in/dancannon/gorethink.v1"
	"io"
	"log"
	"os"
)

type Price struct {
	PurchasingPrice float64 `json:"Purchasing price"`
	SellingPrice    float64 `json:"Selling price"`
	MinimalAmount   int     `json:"Summ"`
	Phones          string  `json:"Phone number"`
}

func main() {

	fmt.Println("Connecting to RethinkDB")
	//Connect to RethinkDB
	session, err := r.Connect(r.ConnectOpts{
		Address:  "172.20.10.14:28015",
		Database: "MoneyData",
	})

	err = r.DB("MoneyData").TableDrop("Price").Exec(session)
	err = r.DB("MoneyData").TableCreate("Price").Exec(session)
	if err != nil {
		log.Fatal("Could not create table")
	}

	IFileName := "/Users/zhaoshixiang/Desktop/RBC_OUT-3.json"
	IFile, err := os.Open(IFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer IFile.Close()

	dec := json.NewDecoder(IFile)
	var res Price
	for {
		if err := dec.Decode(&res); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println()
		}
		/*fmt.Printf("Purchasing price: %f\n", res.PurchasingPrice)
		fmt.Printf("Selling price: %f\n", res.SellingPrice)
		fmt.Printf("MinimalAmount: %d\n", res.MinimalAmount)
		fmt.Printf("Phones: %s\n", res.Phones)*/
		_, err0 := r.DB("MoneyData").Table("Price").Insert(res).RunWrite(session)
		if err0 != nil {
			log.Fatal(err)
		}
	}
}
