package packet

import "github.com/hunterros-s/go-dns-server/dns"

type Packet struct {
	Header      dns.Header
	Questions   []dns.Question
	Answers     []dns.Record
	Authorities []dns.Record
	Resources   []dns.Record
}

func NewPacket() dns.Packet {
	return &Packet{
		Header:      newHeader(),
		Questions:   make([]dns.Question, 0),
		Answers:     make([]dns.Record, 0),
		Authorities: make([]dns.Record, 0),
		Resources:   make([]dns.Record, 0),
	}
}

func (p *Packet) Write(buffer dns.Buffer) error {
	h := p.GetHeader()

	h.SetQuestionsCount(uint16(len(p.Questions)))
	h.SetAnswersCount(uint16(len(p.Answers)))
	h.SetAuthoritativeEntriesCount(uint16(len(p.Authorities)))
	h.SetResourceEntriesCount(uint16(len(p.Resources)))

	err := h.Write(buffer)
	if err != nil {
		return err
	}

	for _, q := range p.Questions {
		err = q.Write(buffer)
		if err != nil {
			return err
		}
	}

	for _, a := range p.Answers {
		err = a.Write(buffer)
		if err != nil {
			return err
		}
	}

	for _, a := range p.Authorities {
		err = a.Write(buffer)
		if err != nil {
			return err
		}
	}

	for _, r := range p.Resources {
		err = r.Write(buffer)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Packet) GetHeader() dns.Header {
	return p.Header
}

func (p *Packet) GetQuestions() []dns.Question {
	return p.Questions
}

func (p *Packet) GetAnswers() []dns.Record {
	return p.Answers
}

func (p *Packet) GetAuthorities() []dns.Record {
	return p.Authorities
}

func (p *Packet) GetResources() []dns.Record {
	return p.Resources
}

func (p *Packet) SetHeader(h dns.Header) {
	p.Header = h
}

func (p *Packet) AppendQuestion(question dns.Question) {
	p.Questions = append(p.Questions, question)

}

func (p *Packet) AppendAnswer(answer dns.Record) {
	p.Answers = append(p.Answers, answer)
}

func (p *Packet) AppendAuthority(authority dns.Record) {
	p.Authorities = append(p.Authorities, authority)
}

func (p *Packet) AppendResource(resource dns.Record) {
	p.Resources = append(p.Resources, resource)
}
