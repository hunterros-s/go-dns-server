package registry

import "github.com/hunterros-s/go-dns-server/dns"

func init() {
	RegisterRecordFactory(dns.CNAME, newCNAMERecord)
}

type cname struct {
	Domain string
	Host   string
	TTL    uint32
}

func newCNAMERecord(info dns.RecordInfo, buffer dns.Buffer) (dns.Record, error) {
	c_name, err := buffer.ReadQName()
	if err != nil {
		return nil, err
	}

	return &cname{
		Domain: info.GetQName(),
		Host:   c_name,
		TTL:    info.GetTTL(),
	}, nil
}

func (r *cname) Write(buffer dns.Buffer) error {

	if err := buffer.WriteQName(r.Domain); err != nil {
		return err
	}
	if err := buffer.WriteU16(uint16(dns.CNAME)); err != nil {
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
