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

//type DataKV map[string]interface{}
type DataKV struct {
	Data interface{} `form:"data" json:"data" xml:"data"  binding:"required"`
}

func NewStorage(dir string) *Storage {
	log.Println("Init JSON storage...", dir)
	//create dir
	mkdirIfNotExist(dir)
	node, _ := snowflake.NewNode(1)

	return &Storage{db: new(jsonstore.JSONStore), dir: dir, idNode: node}
}

//查询bucket中 key 全部
func (s *Storage) ReadAll(bucket string, key string) map[string]json.RawMessage {

	s.loadPersistent(bucket)

	return s.db.GetAll(regexp.MustCompile(key))
}

//查询单个
func (s *Storage) ReadOne(bucket string, key string) json.RawMessage {

	s.loadPersistent(bucket)

	var rs json.RawMessage

	_, rs = s.db.GetRawMessage(key)

	return rs
}

//保存key,value. bucket类似table
func (s *Storage) Create(bucket string, key string, kv DataKV) string {

	s.loadPersistent(bucket)

	//默认自增ID
	//id := key + ":" + strconv.Itoa(len(s.db.Keys())+1)
	id := key + ":" + s.idNode.Generate().String()

	err := s.db.Set(id, kv.Data)
	if err != nil {
		panic(err)
	}

	s.savePersistent(bucket)

	return id
}

// 根据key更新
func (s *Storage) Update(bucket string, key string, kv DataKV) error {

	s.loadPersistent(bucket)

	err := s.db.Set(key, kv.Data)
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
