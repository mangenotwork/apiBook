package db

import (
	"apiBook/common/conf"
	"apiBook/common/log"
	"apiBook/common/utils"
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
)

var (
	DB            *LocalDB
	ISNULL        = fmt.Errorf("ISNULL")
	TableNotFound = fmt.Errorf("table notfound")
)

type LocalDB struct {
	Path   string
	Tables []string
	Conn   *bolt.DB
}

func NewLocalDB(tables []string) *LocalDB {
	path := ""
	if dbPath, ok := conf.Conf.YamlData["dbPath"]; ok {
		path = utils.AnyToString(dbPath)
	}
	return &LocalDB{
		Path:   path,
		Tables: tables,
	}
}

func GetDBConn() *bolt.DB {
	DB.Open()
	return DB.Conn
}

func Init() {
	DB = NewLocalDB(Tables)
	DB.Init()
}

func (ldb *LocalDB) Init() {
	db, err := bolt.Open(ldb.Path, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		_ = db.Close()
	}()
	for _, table := range ldb.Tables {
		err = db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(table))
			if b == nil {
				_, err = tx.CreateBucket([]byte(table))
				if err != nil {
					log.Panic(err)
				}
			}
			return nil
		})
		if err != nil {
			log.Panic(err)
		}
	}
}

func (ldb *LocalDB) Open() {
	ldb.Conn, _ = bolt.Open(ldb.Path, 0600, nil)
}

func (ldb *LocalDB) Close() {
	_ = ldb.Conn.Close()
}

func (ldb *LocalDB) Get(table, key string, data interface{}) error {
	ldb.Open()
	defer func() {
		_ = ldb.Conn.Close()
	}()
	return ldb.Conn.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(table))
		if b == nil {
			return TableNotFound
		}
		bt := b.Get([]byte(key))
		if len(bt) < 1 {
			return ISNULL
		}
		err := json.Unmarshal(bt, data)
		if err != nil {
			return err
		}
		return nil
	})
}

func (ldb *LocalDB) Set(table, key string, data interface{}) error {
	ldb.Open()

	defer func() {
		_ = ldb.Conn.Close()
	}()

	value, err := utils.AnyToJsonB(data)
	if err != nil {
		return err
	}

	return ldb.Conn.Update(func(tx *bolt.Tx) error {
	R:
		b := tx.Bucket([]byte(table))
		if b == nil {
			_, err = tx.CreateBucket([]byte(table))
			if err != nil {
				return err
			}
			goto R
		}
		err = b.Put([]byte(key), value)
		if err != nil {
			return err
		}
		return nil
	})
}

func (ldb *LocalDB) Delete(table, key string) error {
	ldb.Open()
	defer func() {
		_ = ldb.Conn.Close()
	}()
	return ldb.Conn.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(table))
		if b == nil {
			return fmt.Errorf("未获取到表")
		}
		if err := b.Delete([]byte(key)); err != nil {
			return err
		}
		return nil
	})
}

func (ldb *LocalDB) ClearTable(table string) error {
	ldb.Open()
	defer func() {
		_ = ldb.Conn.Close()
	}()
	return ldb.Conn.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket([]byte(table))
	})
}

func (ldb *LocalDB) Stats(table string) (bolt.BucketStats, error) {

	var stats bolt.BucketStats

	ldb.Open()

	defer func() {
		_ = ldb.Conn.Close()
	}()

	err := ldb.Conn.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(table))
		if b == nil {
			err := ldb.ClearTable(table)
			if err != nil {
				return err
			}
		}

		stats = b.Stats()

		return nil
	})

	return stats, err
}

func (ldb *LocalDB) AllKey(table string) ([]string, error) {
	keys := make([]string, 0)

	ldb.Open()

	defer func() {
		_ = ldb.Conn.Close()
	}()

	err := ldb.Conn.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(table))
		if b == nil {
			return TableNotFound
		}

		c := b.Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			keys = append(keys, string(k))
		}
		return nil
	})
	return keys, err
}

func (ldb *LocalDB) GetAll(table string, f func(k, v []byte)) error {
	ldb.Open()

	defer func() {
		_ = ldb.Conn.Close()
	}()

	err := ldb.Conn.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(table))
		if b == nil {
			return TableNotFound
		}

		return b.ForEach(func(k, v []byte) error {
			f(k, v)
			return nil
		})

	})
	return err
}
