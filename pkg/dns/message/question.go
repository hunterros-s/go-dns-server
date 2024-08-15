package dns

import (
	"github.com/hunterros-s/go-dns-server/pkg/buffer"
	"github.com/hunterros-s/go-dns-server/pkg/dns/enum"
)

type Question struct {
	name  string
	qtype enum.QueryType
}

func NewQuestion(name string, qtype enum.QueryType) *Question {
	return &Question{
		name:  name,
		qtype: qtype,
	}
}

func (q *Question) Read(buffer *buffer.PacketBuffer) error {
	name, err := buffer.ReadQName()
	if err != nil {
		return err
	}
	num, err := buffer.ReadU16()
	if err != nil {
		return err
	}
	qtype := enum.QueryType(num)

	q.name = name
	q.qtype = qtype

	return nil
}
