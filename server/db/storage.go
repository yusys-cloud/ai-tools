// Author: yangzq80@gmail.com
// Date: 2021-03-16
//
package db

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/xujiajun/nutsdb"
)

type Storage struct {
	db *nutsdb.DB
}

func NewStorage(dir string) *Storage {
	opt := nutsdb.DefaultOptions
	opt.Dir = dir //这边数据库会自动创建这个目录文件
	db, err := nutsdb.Open(opt)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	return &Storage{db: db}
}

//查询bucket中全部
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

//根据key前缀查询bucket中全部
func (s *Storage) GetAllByPrfix(bucket string, prefix string) {
	if err := s.db.View(
		func(tx *nutsdb.Tx) error {

			if entries, err := tx.PrefixScan(bucket, []byte(prefix), 100); err != nil {
				return err
			} else {
				for _, entry := range entries {
					fmt.Println(string(entry.Key), string(entry.Value))
				}
			}

			return nil
		}); err != nil {
		log.Println(err)
	}
}

//保存key,value
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

//根据key删除
func (s *Storage) Delete(bucket string, key string) {
	if err := s.db.Update(
		func(tx *nutsdb.Tx) error {
			if err := tx.Delete(bucket, []byte(key)); err != nil {
				return err
			}
			return nil
		}); err != nil {
		log.Fatal(err)
	}
}
