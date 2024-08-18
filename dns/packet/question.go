package packet

import "github.com/hunterros-s/go-dns-server/dns"

type Question struct {
	Name  string
	QType dns.QueryType
}

func NewQuestion(name string, qtype dns.QueryType) dns.Question {
	return &Question{
		Name:  name,
		QType: qtype,
	}
}

func (q *Question) Write(buffer dns.Buffer) error {
	if err := buffer.WriteQName(q.Name); err != nil {
		return err
	}

	if err := buffer.WriteU16(uint16(q.QType)); err != nil {
		return err
	}

	if err := buffer.WriteU16(1); err != nil {
		return err
	}

	return nil
}

func (q *Question) GetName() string {
	return q.Name
}

func (q *Question) GetQType() dns.QueryType {
	return q.QType
}

func (q *Question) SetName(name string) {
	q.Name = name
}

func (q *Question) SetQType(qType dns.QueryType) {
	q.QType = qType
}
