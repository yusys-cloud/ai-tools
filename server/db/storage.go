// Author: yangzq80@gmail.com
// Date: 2021-03-16
//
package db

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/snowflake"
	log "github.com/sirupsen/logrus"
	"github.com/xujiajun/utils/filesystem"
	"github.com/yusys-cloud/ai-tools/server/db/jsonstore"
	"os"
	"regexp"
)

type Storage struct {
	db     *jsonstore.JSONStore
	dir    string
	idNode *snowflake.Node
}

type Data struct {
	K string      `json:"k"`
	V interface{} `json:"v"`
}

func NewStorage(dir string) *Storage {
	log.Println("Init JSON storage...", dir)
	//create dir
	mkdirIfNotExist(dir)
	node, _ := snowflake.NewNode(1)

	return &Storage{db: new(jsonstore.JSONStore), dir: dir, idNode: node}
}

//查询bucket中 key 全部
func (s *Storage) ReadAll(bucket string, key string) []Data {

	s.loadPersistent(bucket)

	rs := s.db.GetAll(regexp.MustCompile(key))

	return convertMapToArray(rs)
}

//查询单个
func (s *Storage) ReadOne(bucket string, key string) Data {

	s.loadPersistent(bucket)

	_, rs := s.db.GetRawMessage(key)

	var f interface{}

	json.Unmarshal(rs, &f)

	return Data{key, f}
}
func (s *Storage) ReadOneRaw(bucket string, key string) []byte {

	s.loadPersistent(bucket)

	_, rs := s.db.GetRawMessage(key)

	return rs
}

//保存key,value. bucket类似table
func (s *Storage) Create(bucket string, key string, value interface{}) string {

	s.loadPersistent(bucket)

	//默认自增ID
	id := key + ":" + s.idNode.Generate().String()

	err := s.db.Set(id, value)
	if err != nil {
		panic(err)
	}

	s.savePersistent(bucket)

	return id
}

// 根据key更新
func (s *Storage) Update(bucket string, key string, value interface{}) error {

	s.loadPersistent(bucket)

	err := s.db.Set(key, value)
	if err != nil {
		panic(err)
	}

	s.savePersistent(bucket)

	return err
}
func (s *Storage) UpdateMarshalValue(bucket string, key string, value []byte) error {

	s.loadPersistent(bucket)

	err := s.db.SetMarshalValue(key, value)
	if err != nil {
		panic(err)
	}

	s.savePersistent(bucket)

	return err
}

// 根据key删除
func (s *Storage) Delete(bucket string, key string) {

	s.loadPersistent(bucket)

	s.db.Delete(key)

	s.savePersistent(bucket)
}

func (s *Storage) loadPersistent(bucket string) {
	if ss, err := jsonstore.Open(s.getFileName(bucket)); err == nil {
		s.db = ss
	}
}

func (s *Storage) savePersistent(bucket string) {
	// Saving will automatically gzip if .gz is provided
	if err := jsonstore.Save(s.db, s.getFileName(bucket)); err != nil {
		log.Error(err)
		panic(err)
	}
}

func (s *Storage) getFileName(bucket string) string {
	return s.dir + "/" + bucket + ".json.gz"
}

func mkdirIfNotExist(rootDir string) error {
	if ok := filesystem.PathIsExist(rootDir); !ok {
		if err := os.MkdirAll(rootDir, os.ModePerm); err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func convertMapToArray(raw map[string]json.RawMessage) []Data {
	var datas []Data
	for k, v := range raw {
		datas = append(datas, Data{k, v})
	}
	return datas
}
