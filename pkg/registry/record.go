package registry

import "github.com/hunterros-s/go-dns-server/pkg/domain"

type RecordFactoryFunc func(domain.RecordInfo, domain.Buffer) (domain.Record, error)

var queryTypeToRecordFactory = make(map[domain.QueryType]RecordFactoryFunc)

func RegisterRecordFactory(qt domain.QueryType, factory RecordFactoryFunc) {
	queryTypeToRecordFactory[qt] = factory
}

func GetRecordFactory(qt domain.QueryType) (RecordFactoryFunc, bool) {
	factory, ok := queryTypeToRecordFactory[qt]
	return factory, ok
}

type RecordRegistry struct{}

func (rf *RecordRegistry) Get(qt domain.QueryType) (RecordFactoryFunc, bool) {
	factory, ok := queryTypeToRecordFactory[qt]
	return factory, ok
}
