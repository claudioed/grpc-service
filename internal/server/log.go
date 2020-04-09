package server

import (
	"fmt"
	api "github.com/claudioed/grpc-service/api/v1"
	"sync"
)

type Log struct {
	mu sync.Mutex
	records []*api.Record
}

var ErrOffsetNotFound = fmt.Errorf("offset not found")

// creates new log instance
func NewLog() *Log  {
	return &Log{}
}

func (l *Log) Append(record *api.Record) (uint64,error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	record.Offset = uint64(len(l.records))
	l.records = append(l.records,record)
	return record.Offset,nil
}

func (l *Log) Read(offset uint64) (*api.Record, error) {
	l.mu.Lock()
	defer  l.mu.Unlock()
	if offset >= uint64(len(l.records)){
		return &api.Record{},ErrOffsetNotFound
	}
	return l.records[offset],nil
}