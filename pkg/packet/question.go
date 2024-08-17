package packet

import "github.com/hunterros-s/go-dns-server/pkg/domain"

type Question struct {
	Name  string
	QType domain.QueryType
}

func NewQuestion(name string, qtype domain.QueryType) domain.Question {
	return &Question{
		Name:  name,
		QType: qtype,
	}
}

func (q *Question) GetName() string {
	return q.Name
}

func (q *Question) GetQType() domain.QueryType {
	return q.QType
}

func (q *Question) SetName(name string) {
	q.Name = name
}

func (q *Question) SetQType(qType domain.QueryType) {
	q.QType = qType
}
