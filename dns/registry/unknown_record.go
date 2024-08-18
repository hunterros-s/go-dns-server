package registry

import (
	"log"

	"github.com/hunterros-s/go-dns-server/dns"
)

type unknown struct {
	Domain  string
	QType   dns.QueryType
	DataLen uint16
	TTL     uint32
}

func New_unknown_record(record dns.RecordInfo, b dns.Buffer) (dns.Record, error) {
	b.Step(record.GetRDataLength())

	return &unknown{
		Domain:  record.GetQName(),
		QType:   record.GetQType(),
		DataLen: record.GetRDataLength(),
		TTL:     record.GetTTL(),
	}, nil
}

func (r *unknown) Write(buffer dns.Buffer) error {
	log.Panicln("not writing unknown record")
	return nil
}
