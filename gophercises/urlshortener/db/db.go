package db

import (
	"fmt"
	"github.com/boltdb/bolt"
)

type MyDB bolt.DB

func Open() (db *MyDB, err error) {
	boltdb, err := bolt.Open("my.db", 0600, &bolt.Options{})
	if err != nil {
		return nil, err
	}

	return (*MyDB)(boltdb), nil
}

func (db *MyDB) Close() error {
	boltdb := (*bolt.DB)(db)
	err := boltdb.Close()
	if err != nil {
		return err
	}

	return nil
}

func (db *MyDB) Insert(bucketName string, key string, value string) error {
	boltdb := (*bolt.DB)(db)
	err := boltdb.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return bucket.Put([]byte(key), []byte(value))
	})

	if err != nil {
		return err
	}
	return nil
}

func (db *MyDB) Read(bucketName string, key string) (value string, err error) {
	boltdb := (*bolt.DB)(db)
	err = boltdb.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return fmt.Errorf("get bucket: FAILED")
		}

		byteValue := b.Get([]byte(key))
		if byteValue == nil {
			return fmt.Errorf("read failed with the key %s", string(key))
		}

		value = string(byteValue)
		return nil
	})

	if err != nil {
		return "", err
	}
	return value, nil
}

func (db *MyDB) Dump() (dump string) {
	boltdb := (*bolt.DB)(db)
	boltdb.View(func(tx *bolt.Tx) error {
		tx.ForEach(func(bucketName []byte, b *bolt.Bucket) (err error) {
			c := b.Cursor()
			for k, v := c.First(); k != nil; k, v = c.Next() {
				dump = dump + fmt.Sprintln("Key: ", string(k), " Value: ", string(v))
			}
			return nil
		})
		return nil
	})

	return dump
}

func (db *MyDB) FillDB() {
	db.Insert("b", "/urlshort", "https://github.com/gophercises/urlshort")
	db.Insert("b", "/urlshort-final", "https://github.com/gophercises/urlshort/tree/solution")
	db.Insert("b", "/urlshort-godoc", "https://godoc.org/github.com/gophercises/urlshort")
	db.Insert("b", "/yaml-godoc", "https://godoc.org/gopkg.in/yaml.v2")
}

func (db *MyDB) GetMap() (dbData map[string]string) {
	dbData = make(map[string]string)
	boltdb := (*bolt.DB)(db)

	boltdb.View(func(tx *bolt.Tx) error {
		tx.ForEach(func(bucketName []byte, b *bolt.Bucket) (err error) {
			c := b.Cursor()
			for k, v := c.First(); k != nil; k, v = c.Next() {
				dbData[string(k)] = string(v)
			}
			return nil
		})
		return nil
	})

	return dbData
}
