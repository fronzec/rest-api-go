package dao

// TODO import model from models directory from current project
import (
	"github.com/fronzec/rest-api-go/dao"
	. "github.com/fronzec/rest-api-go/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

type MoviesDAO struct {
	Database     string
	Server       string
	AuthDatabase string
	AuthUsername string
	AuthPassword string
}

var db *mgo.Database

const (
	COLLECTION = "movies"
)

func (m *MoviesDAO) Connect() {
	print(">>>>> Connecting to database")
	mongoDBDialInfo := &mgo.DialInfo{
		Database: dao.AuthDatabase,
		Username: dao.AuthUsername,
		Password: dao.AuthPassword,
		Addrs:    []string{dao.Server},
		Timeout:  60 * time.Second,
	}

	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *MoviesDAO) FindAll() ([]Movie, error) {
	var movies []Movie
	err := db.C(COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}

func (m *MoviesDAO) FindById(id string) (Movie, error) {
	var movie Movie
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}

func (m *MoviesDAO) Insert(movie Movie) error {
	err := db.C(COLLECTION).Insert(&movie)
	return err
}

func (m *MoviesDAO) Delete(movie Movie) error {
	err := db.C(COLLECTION).Remove(&movie)
	return err
}

func (m *MoviesDAO) Update(movie Movie) error {
	err := db.C(COLLECTION).UpdateId(movie.ID, &movie)
	return err
}
