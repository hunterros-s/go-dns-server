package registry

import "github.com/hunterros-s/go-dns-server/pkg/domain"

type unknown struct {
	Domain  string
	QType   domain.QueryType
	DataLen uint16
	TTL     uint32
}

func New_unknown_record(record domain.RecordInfo, b domain.Buffer) (domain.Record, error) {
	return &unknown{
		Domain:  record.GetQName(),
		QType:   record.GetQType(),
		DataLen: record.GetRDataLength(),
		TTL:     record.GetTTL(),
	}, nil
}
