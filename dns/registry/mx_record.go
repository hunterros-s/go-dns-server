package registry

import "github.com/hunterros-s/go-dns-server/dns"

func init() {
	RegisterRecordFactory(dns.MX, newMXRecord)
}

type mx struct {
	Domain   string
	Priority uint16
	Host     string
	TTL      uint32
}

func newMXRecord(info dns.RecordInfo, buffer dns.Buffer) (dns.Record, error) {
	priority, err := buffer.ReadU16()
	if err != nil {
		return nil, err
	}

	mx_, err := buffer.ReadQName()
	if err != nil {
		return nil, err
	}

	return &mx{
		Domain:   info.GetQName(),
		Priority: priority,
		Host:     mx_,
		TTL:      info.GetTTL(),
	}, nil
}

func (r *mx) Write(buffer dns.Buffer) error {

	if err := buffer.WriteQName(r.Domain); err != nil {
		return err
	}
	if err := buffer.WriteU16(uint16(dns.MX)); err != nil {
		return err
	}
	if err := buffer.WriteU16(1); err != nil {
		return err
	}
	if err := buffer.WriteU32(r.TTL); err != nil {
		return err
	}

	pos := buffer.Pos()
	if err := buffer.WriteU16(0); err != nil {
		return err
	}

	if err := buffer.WriteQName(r.Host); err != nil {
		return err
	}

	size := buffer.Pos() - (pos + 2)
	buffer.SetU16(pos, size)

	return nil
}
