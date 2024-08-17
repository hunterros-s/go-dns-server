package parser

import (
	"github.com/hunterros-s/go-dns-server/pkg/domain"
)

func (p *Parser) parseSections(buffer domain.Buffer, packet domain.Packet) error {
	if err := p.parseQuestions(buffer, packet); err != nil {
		return err
	}
	if err := p.parseAnswers(buffer, packet); err != nil {
		return err
	}
	if err := p.parseAuthorities(buffer, packet); err != nil {
		return err
	}
	if err := p.parseResources(buffer, packet); err != nil {
		return err
	}
	return nil
}

func (p *Parser) parseQuestions(buffer domain.Buffer, packet domain.Packet) error {
	for i := 0; i < int(packet.GetQuestionsCount()); i++ {
		question, err := p.questionfactory.New(buffer)
		if err != nil {
			return err
		}
		packet.AppendQuestion(question)
	}
	return nil
}

func (p *Parser) parseAnswers(buffer domain.Buffer, packet domain.Packet) error {
	for i := 0; i < int(packet.GetAnswersCount()); i++ {
		record, err := p.recordfactory.New(buffer)
		if err != nil {
			return err
		}
		packet.AppendAnswer(record)
	}
	return nil
}

func (p *Parser) parseAuthorities(buffer domain.Buffer, packet domain.Packet) error {
	for i := 0; i < int(packet.GetAuthoritativeEntriesCount()); i++ {
		record, err := p.recordfactory.New(buffer)
		if err != nil {
			return err
		}
		packet.AppendAuthority(record)
	}
	return nil
}

func (p *Parser) parseResources(buffer domain.Buffer, packet domain.Packet) error {
	for i := 0; i < int(packet.GetResourceEntriesCount()); i++ {
		record, err := p.recordfactory.New(buffer)
		if err != nil {
			return err
		}
		packet.AppendResource(record)
	}
	return nil
}
