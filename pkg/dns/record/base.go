package record

import (
	"github.com/hunterros-s/go-dns-server/pkg/buffer"
	"github.com/hunterros-s/go-dns-server/pkg/dns/enum"
)

type Base struct {
	QNAME       string
	QTYPE       enum.QueryType
	QCLASS      uint16
	TTL         uint32
	RDATALength uint16
}

func (r *Base) read(buffer *buffer.PacketBuffer) error {
	domain, err := buffer.ReadQName()
	if err != nil {
		return err
	}

	qtype_num, err := buffer.ReadU16()
	if err != nil {
		return err
	}
	qtype := enum.QueryType(qtype_num)

	qclass, err := buffer.ReadU16()
	if err != nil {
		return err
	}

	ttl, err := buffer.ReadU32()
	if err != nil {
		return err
	}

	datalen, err := buffer.ReadU16()
	if err != nil {
		return err
	}

	r.QNAME = domain
	r.QTYPE = qtype
	r.QCLASS = qclass
	r.TTL = ttl
	r.RDATALength = datalen

	return nil
}

func NewBase(buffer *buffer.PacketBuffer) (*Base, error) {
	r := &Base{}
	err := r.read(buffer)
	if err != nil {
		return nil, err
	}
	return r, nil
}
