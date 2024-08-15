package record

import (
	"github.com/hunterros-s/go-dns-server/pkg/buffer"
	"github.com/hunterros-s/go-dns-server/pkg/dns/enum"
)

type UNKNOWNRecord struct {
	domain   string
	qtype    enum.QueryType
	data_len uint16
	ttl      uint32
}

func NewUNKNOWNRecord(record *Base, b *buffer.PacketBuffer) (*UNKNOWNRecord, error) {
	return &UNKNOWNRecord{
		domain:   record.QNAME,
		qtype:    record.QTYPE,
		data_len: record.RDATALength,
		ttl:      record.TTL,
	}, nil
}
