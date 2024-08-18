package registry

import "github.com/hunterros-s/go-dns-server/dns"

type RecordFactoryFunc func(dns.RecordInfo, dns.Buffer) (dns.Record, error)

var queryTypeToRecordFactory = make(map[dns.QueryType]RecordFactoryFunc)

func RegisterRecordFactory(qt dns.QueryType, factory RecordFactoryFunc) {
	queryTypeToRecordFactory[qt] = factory
}

func GetRecordFactory(qt dns.QueryType) (RecordFactoryFunc, bool) {
	factory, ok := queryTypeToRecordFactory[qt]
	return factory, ok
}

type RecordRegistry struct{}

func (rf *RecordRegistry) Get(qt dns.QueryType) (RecordFactoryFunc, bool) {
	factory, ok := queryTypeToRecordFactory[qt]
	return factory, ok
}
