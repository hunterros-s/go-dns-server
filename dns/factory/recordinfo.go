package factory

import "github.com/hunterros-s/go-dns-server/dns"

type RecordInfo struct {
	QName       string
	QType       dns.QueryType
	QClass      uint16
	TTL         uint32
	RDataLength uint16
}

func NewRecordInfo(buffer dns.Buffer) (dns.RecordInfo, error) {
	domainname, err := buffer.ReadQName()
	if err != nil {
		return nil, err
	}

	qtype_num, err := buffer.ReadU16()
	if err != nil {
		return nil, err
	}
	qtype := dns.QueryType(qtype_num)

	qclass, err := buffer.ReadU16()
	if err != nil {
		return nil, err
	}

	ttl, err := buffer.ReadU32()
	if err != nil {
		return nil, err
	}

	datalen, err := buffer.ReadU16()
	if err != nil {
		return nil, err
	}

	return &RecordInfo{
		QName:       domainname,
		QType:       qtype,
		QClass:      qclass,
		TTL:         ttl,
		RDataLength: datalen,
	}, nil
}

func (ri *RecordInfo) GetQName() string {
	return ri.QName
}

func (ri *RecordInfo) GetQType() dns.QueryType {
	return ri.QType
}

func (ri *RecordInfo) GetQClass() uint16 {
	return ri.QClass
}

func (ri *RecordInfo) GetTTL() uint32 {
	return ri.TTL
}

func (ri *RecordInfo) GetRDataLength() uint16 {
	return ri.RDataLength
}
