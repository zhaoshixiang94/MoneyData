package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	//"time"
	r "gopkg.in/dancannon/gorethink.v1"
)

type Money struct {
	Kind     string  `json:"Kind"`
	Exchange float64 `json:"Exchange"`
}

func main() {
	fmt.Println("Connecting to RethinkDB")

	session, err := r.Connect(r.ConnectOpts{
		Address:  "172.20.10.14:28015",
		Database: "MoneyData",
	})
	if err != nil {
		log.Fatal("Could not connect")
	}

	err = r.DB("MoneyData").TableDrop("data").Exec(session)
	err = r.DB("MoneyData").TableCreate("data").Exec(session)
	if err != nil {
		log.Fatal("Could not create table")
	}

	/*err = r.DB("MoneyData").Table("data").IndexCreate("Data").Exec(session)
	if err != nil {
		log.Fatal("Could not create index")
	}*/

	var moneydata Money
	IFileName := "/Users/zhaoshixiang/Desktop/curentKYPC.json"
	IFile, err := os.Open(IFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer IFile.Close()

	dec := json.NewDecoder(IFile)
	for {
		if err := dec.Decode(&moneydata); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println()
		}
		fmt.Printf("Kind: %f\n", moneydata.Kind)
		fmt.Printf("Exchange: %f\n", moneydata.Exchange)
		_, err0 := r.DB("MoneyData").Table("data").Insert(moneydata).RunWrite(session)
		if err0 != nil {
			log.Fatal(err0)
		}
	}
}

/*file, err := os.Open("/Users/zhaoshixiang/Desktop/data.txt") //打开文件
defer file.Close()                                           //打开文件出错处理
if nil == err {
	buff := bufio.NewReader(file) //读入缓存
	for {
		line, err := buff.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		//fmt.Print(line) //可以对一行进行处理
		moneydata.Kind = line[0:3]
		moneydata.Exchange = line[4:len(line)]
		//fmt.Println(moneydata.Kind)
		//fmt.Println(moneydata.Exchange)
		_, err0 := r.DB("MoneyData").Table("data").Insert(moneydata).RunWrite(session)
		if err0 != nil {
			log.Fatal(err0)
		}
	}
}*/
/*_, err1 := r.DB("MoneyData").Table("data").Updata(moneydata).RunWrite(session)
if err1 != nil {
	log.Fatal(err1)
}
time.Sleep(100 * time.Millisecond)*/
