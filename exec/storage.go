// Author: yangzq80@gmail.com
// Date: 2021-03-16
//
package exec

import (
	log "github.com/sirupsen/logrus"
	"github.com/xujiajun/nutsdb"
)

type Storage struct {
	db *nutsdb.DB
}

func NewStorage() *Storage {
	opt := nutsdb.DefaultOptions
	opt.Dir = "/tmp/nutsdb" //这边数据库会自动创建这个目录文件
	db, err := nutsdb.Open(opt)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	return &Storage{db: db}
}

func (s *Storage) GetAll(bucket string) {
	if err := s.db.View(
		func(tx *nutsdb.Tx) error {
			entries, err := tx.GetAll(bucket)
			if err != nil {
				return err
			}

			for _, entry := range entries {
				log.Println(string(entry.Key), string(entry.Value))
			}

			return nil
		}); err != nil {
		log.Println(err)
	}
}

func (s *Storage) Save(bucket string, key string, value string) {
	if err := s.db.Update(
		func(tx *nutsdb.Tx) error {
			if err := tx.Put(bucket, []byte(key), []byte(value), 0); err != nil {
				return err
			}
			return nil
		}); err != nil {
		log.Fatal(err)
	}
}
