package registry

import (
	"encoding/binary"
	"net"

	"github.com/hunterros-s/go-dns-server/dns"
)

func init() {
	RegisterRecordFactory(dns.AAAA, newAAAARecord)
}

type aaaa struct {
	Domain string
	Addr   net.IP
	TTL    uint32
}

func newAAAARecord(info dns.RecordInfo, buffer dns.Buffer) (dns.Record, error) {
	raw1, err := buffer.ReadU32()
	if err != nil {
		return nil, err
	}
	raw2, err := buffer.ReadU32()
	if err != nil {
		return nil, err
	}
	raw3, err := buffer.ReadU32()
	if err != nil {
		return nil, err
	}
	raw4, err := buffer.ReadU32()
	if err != nil {
		return nil, err
	}

	addr := make([]byte, 16)
	binary.BigEndian.PutUint32(addr[0:4], raw1)
	binary.BigEndian.PutUint32(addr[4:8], raw2)
	binary.BigEndian.PutUint32(addr[8:12], raw3)
	binary.BigEndian.PutUint32(addr[12:16], raw4)

	// Create an IPv6 address from the 16-byte slice
	ip := net.IP(addr)

	return &aaaa{
		Domain: info.GetQName(),
		Addr:   ip,
		TTL:    info.GetTTL(),
	}, nil
}

func (r *aaaa) Write(buffer dns.Buffer) error {

	if err := buffer.WriteQName(r.Domain); err != nil {
		return err
	}
	if err := buffer.WriteU16(uint16(dns.AAAA)); err != nil {
		return err
	}
	if err := buffer.WriteU16(1); err != nil {
		return err
	}
	if err := buffer.WriteU32(r.TTL); err != nil {
		return err
	}
	if err := buffer.WriteU16(16); err != nil {
		return err
	}

	for octet := range r.Addr.To16() {
		if err := buffer.WriteU16(uint16(octet)); err != nil {
			return err
		}
	}

	return nil
}
