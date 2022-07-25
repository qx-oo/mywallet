package db

import (
	"fmt"
	"path"

	"github.com/boltdb/bolt"
)

const DB_NAME = "MyWallet"

type DB struct {
	filename string
	db       *bolt.DB
}

func NewDB(dir string) (*DB, error) {
	filename := path.Join(dir, "wallet.db")
	db, err := bolt.Open(filename, 0600, nil)
	if err != nil {
		return nil, err
	}
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(DB_NAME))
		return err
	})
	if err != nil {
		return nil, err
	}
	return &DB{filename: filename, db: db}, nil
}

func (cli *DB) Exists(name string) (ret bool, err error) {
	err = cli.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(DB_NAME))
		val := b.Get([]byte(name))
		if val == nil {
			ret = false
		} else {
			ret = true
		}
		return nil
	})
	return
}

func (cli *DB) GetAddress(name string) (address string, err error) {
	err = cli.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(DB_NAME))
		v := b.Get([]byte(name))
		if v == nil {
			err = fmt.Errorf("key %s not exists", name)
		}
		address = string(v)
		return err
	})
	return
}

func (cli *DB) SaveAddress(name string, address string) error {
	return cli.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(DB_NAME))
		err := b.Put([]byte(name), []byte(address))
		return err
	})
}

func (cli *DB) Delete(name string) error {
	return cli.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(DB_NAME))
		return b.Delete([]byte(name))
	})
}

func (cli *DB) GetAll() (map[string]string, error) {
	data := map[string]string{}
	err := cli.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(DB_NAME))
		b.ForEach(func(k, v []byte) error {
			data[string(k)] = string(v)
			return nil
		})
		return nil
	})
	return data, err
}

func (cli *DB) Close() {
	cli.db.Close()
}
