package dns

import (
	"github.com/hunterros-s/go-dns-server/pkg/buffer"
	"github.com/hunterros-s/go-dns-server/pkg/dns/enum"
	"github.com/hunterros-s/go-dns-server/pkg/dns/record"
)

type DNSRecord interface{}

type RecordFactory func(*record.Base, *buffer.PacketBuffer) (DNSRecord, error)

var queryTypeToRecordFactory = map[enum.QueryType]RecordFactory{
	enum.A: func(d *record.Base, pb *buffer.PacketBuffer) (DNSRecord, error) { return record.NewARecord(d, pb) },
	enum.UNKNOWN: func(d *record.Base, pb *buffer.PacketBuffer) (DNSRecord, error) {
		return record.NewUNKNOWNRecord(d, pb)
	},
}

func GetNewRecordFunc(qt enum.QueryType) (RecordFactory, bool) {
	f, ok := queryTypeToRecordFactory[qt]
	return f, ok
}
