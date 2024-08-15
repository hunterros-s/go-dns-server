package record

import (
	"net"

	"github.com/hunterros-s/go-dns-server/pkg/buffer"
)

type ARecord struct {
	domain string
	addr   net.Addr
	ttl    uint32
}

func (r *ARecord) read(buffer *buffer.PacketBuffer) error {
	raw_addr, err := buffer.ReadU32()
	if err != nil {
		return err
	}
	b1 := byte(raw_addr >> 24)
	b2 := byte(raw_addr >> 16)
	b3 := byte(raw_addr >> 8)
	b4 := byte(raw_addr)

	r.addr = &net.IPAddr{IP: net.IPv4(b1, b2, b3, b4)}
	return nil
}

func NewARecord(record *Base, buffer *buffer.PacketBuffer) (*ARecord, error) {
	r := &ARecord{}

	r.domain = record.QNAME
	r.ttl = record.TTL

	err := r.read(buffer)
	if err != nil {
		return nil, err
	}

	return r, nil
}
