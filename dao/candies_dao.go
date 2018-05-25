package dao

import (
	"fmt"
	"log"

	. "github.com/jafetbntz/candy_shop/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// CandiesDAO tiene como objetivo centralizar
type CandiesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "candies"
)

func (m *CandiesDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *CandiesDAO) GetAll() ([]Candy, error) {
	var candies []Candy
	err := db.C(COLLECTION).Find(bson.M{}).All(&candies)
	return candies, err
}

func (m *CandiesDAO) FindById(id string) (Candy, error) {
	var candy Candy
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&candy)
	return candy, err
}

func (m *CandiesDAO) Insert(candy Candy) error {
	fmt.Println("candy....")
	fmt.Println(candy.Id)

	err := db.C(COLLECTION).Insert(&candy)
	return err
}

func (m *CandiesDAO) Delete(candy Candy) error {
	err := db.C(COLLECTION).Remove(&candy)
	return err
}

func (m *CandiesDAO) Update(candy Candy) error {
	err := db.C(COLLECTION).UpdateId(candy.Id, &candy)
	return err
}
