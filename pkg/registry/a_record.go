package registry

import (
	"net"

	"github.com/hunterros-s/go-dns-server/pkg/domain"
)

func init() {
	RegisterRecordFactory(domain.A, new_a_record)
}

type a struct {
	Domain string
	Addr   net.Addr
	TTL    uint32
}

// NewARecord creates a new ARecord from the base record and buffer.
func new_a_record(info domain.RecordInfo, buffer domain.Buffer) (domain.Record, error) {
	rawAddr, err := buffer.ReadU32()
	if err != nil {
		return nil, err
	}

	addr := &net.IPAddr{
		IP: net.IPv4(
			byte(rawAddr>>24),
			byte(rawAddr>>16),
			byte(rawAddr>>8),
			byte(rawAddr),
		),
	}

	return &a{
		Domain: info.GetQName(),
		Addr:   addr,
		TTL:    info.GetTTL(),
	}, nil
}
