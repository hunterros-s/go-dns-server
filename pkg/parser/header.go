package parser

import "github.com/hunterros-s/go-dns-server/pkg/domain"

func (p *Parser) readHeader(packet domain.Packet, buffer domain.Buffer) error {
	var err error

	// Read the ID
	id, err := buffer.ReadU16()
	if err != nil {
		return err
	}
	packet.SetID(id)

	// Read the flags
	flags, err := buffer.ReadU16()
	if err != nil {
		return err
	}

	// Extract and set the bit fields from the flags
	packet.SetResponse((flags & 0x8000) != 0)              // QR (bit 15)
	packet.SetOpcode(uint8((flags >> 11) & 0x0F))          // Opcode (bits 11-14)
	packet.SetAuthoritativeAnswer((flags & 0x0400) != 0)   // AA (bit 10)
	packet.SetTruncatedMessage((flags & 0x0200) != 0)      // TC (bit 9)
	packet.SetRecursionDesired((flags & 0x0100) != 0)      // RD (bit 8)
	packet.SetRecursionAvailable((flags & 0x0080) != 0)    // RA (bit 7)
	packet.SetZ(uint8((flags >> 4) & 0x07))                // Reserved (bits 4-6)
	packet.SetRescode(domain.ResponseCode(flags & 0x000F)) // RCODE (bits 0-3)

	// Read and set the counts
	questionsCount, err := buffer.ReadU16()
	if err != nil {
		return err
	}
	packet.SetQuestionsCount(questionsCount)

	answersCount, err := buffer.ReadU16()
	if err != nil {
		return err
	}
	packet.SetAnswersCount(answersCount)

	authoritativeEntriesCount, err := buffer.ReadU16()
	if err != nil {
		return err
	}
	packet.SetAuthoritativeEntriesCount(authoritativeEntriesCount)

	resourceEntriesCount, err := buffer.ReadU16()
	if err != nil {
		return err
	}
	packet.SetResourceEntriesCount(resourceEntriesCount)

	return nil
}
