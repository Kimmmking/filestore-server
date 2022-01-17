package store

import (
	"database/sql"
	cfg "filestore-server/config"
	"filestore-server/store"
	"filestore-server/store/factory"
	"fmt"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlStore struct {
	sync.RWMutex
	filemetas map[string]*store.FileMeta
}

var db *sql.DB

func init() {
	// TODO: connect
	db, _ = sql.Open("mysql", cfg.MySQLSource)
	db.SetMaxOpenConns(1000)
	err := db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to mysql, err:" + err.Error())
		os.Exit(1)
	}

	factory.Register("mysql", &MemStore{
		filemetas: make(map[string]*store.FileMeta),
	})
}

func (ms *MysqlStore) Create(filemeta *store.FileMeta) error {
	return nil
}

func (ms *MysqlStore) Update(filemeta *store.FileMeta) error {
	return nil
}

func (ms *MysqlStore) Get(string) (store.FileMeta, error) {
	return store.FileMeta{}, nil
}

func (ms *MysqlStore) GetAll() ([]store.FileMeta, error) {
	return []store.FileMeta{}, nil
}

func (ms *MysqlStore) Delete(name string) error {
	return nil
}
