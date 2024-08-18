package registry

import (
	"net"

	"github.com/hunterros-s/go-dns-server/dns"
)

func init() {
	RegisterRecordFactory(dns.A, new_a_record)
}

type a struct {
	Domain string
	Addr   net.IP
	TTL    uint32
}

// NewARecord creates a new ARecord from the base record and buffer.
func new_a_record(info dns.RecordInfo, buffer dns.Buffer) (dns.Record, error) {
	rawAddr, err := buffer.ReadU32()
	if err != nil {
		return nil, err
	}

	addr := net.IPv4(
		byte(rawAddr>>24),
		byte(rawAddr>>16),
		byte(rawAddr>>8),
		byte(rawAddr),
	)

	return &a{
		Domain: info.GetQName(),
		Addr:   addr,
		TTL:    info.GetTTL(),
	}, nil
}

func (r *a) Write(buffer dns.Buffer) error {

	if err := buffer.WriteQName(r.Domain); err != nil {
		return err
	}

	if err := buffer.WriteU16(uint16(dns.A)); err != nil {
		return err
	}

	if err := buffer.WriteU16(1); err != nil {
		return err
	}

	if err := buffer.WriteU16(uint16(r.TTL)); err != nil {
		return err
	}

	if err := buffer.WriteU16(4); err != nil {
		return err
	}

	ip := r.Addr.To4()
	if err := buffer.WriteByte(ip[0]); err != nil {
		return err
	}
	if err := buffer.WriteByte(ip[1]); err != nil {
		return err
	}
	if err := buffer.WriteByte(ip[2]); err != nil {
		return err
	}
	if err := buffer.WriteByte(ip[3]); err != nil {
		return err
	}

	return nil
}
