package packet

import (
	"github.com/hunterros-s/go-dns-server/pkg/domain"
)

type Packet struct {
	Header      Header
	Questions   []domain.Question
	Answers     []domain.Record
	Authorities []domain.Record
	Resources   []domain.Record
}

func NewPacket() domain.Packet {
	return &Packet{
		Header:      newHeader(),
		Questions:   make([]domain.Question, 0),
		Answers:     make([]domain.Record, 0),
		Authorities: make([]domain.Record, 0),
		Resources:   make([]domain.Record, 0),
	}
}

func (p *Packet) GetQuestions() []domain.Question {
	return p.Questions
}

func (p *Packet) GetAnswers() []domain.Record {
	return p.Answers
}

func (p *Packet) GetAuthorities() []domain.Record {
	return p.Authorities
}

func (p *Packet) GetResources() []domain.Record {
	return p.Resources
}

func (p *Packet) AppendQuestion(question domain.Question) {
	p.Questions = append(p.Questions, question)
}

func (p *Packet) AppendAnswer(answer domain.Record) {
	p.Answers = append(p.Answers, answer)
}

func (p *Packet) AppendAuthority(authority domain.Record) {
	p.Authorities = append(p.Authorities, authority)
}

func (p *Packet) AppendResource(resource domain.Record) {
	p.Resources = append(p.Resources, resource)
}
