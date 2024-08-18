package parser

import "github.com/hunterros-s/go-dns-server/dns"

func (p *Parser) readHeader(packet dns.Packet, buffer dns.Buffer) error {
	var err error

	// Read the ID
	id, err := buffer.ReadU16()
	if err != nil {
		return err
	}
	header := packet.GetHeader()

	header.SetID(id)

	// Read the flags
	flags, err := buffer.ReadU16()
	if err != nil {
		return err
	}

	// Extract and set the bit fields from the flags
	header.SetResponse((flags & 0x8000) != 0)            // QR (bit 15)
	header.SetOpcode(uint8((flags >> 11) & 0x0F))        // Opcode (bits 11-14)
	header.SetAuthoritativeAnswer((flags & 0x0400) != 0) // AA (bit 10)
	header.SetTruncatedMessage((flags & 0x0200) != 0)    // TC (bit 9)
	header.SetRecursionDesired((flags & 0x0100) != 0)    // RD (bit 8)
	header.SetRecursionAvailable((flags & 0x0080) != 0)  // RA (bit 7)
	header.SetZ(uint8((flags >> 4) & 0x07))              // Reserved (bits 4-6)
	header.SetRescode(dns.ResponseCode(flags & 0x000F))  // RCODE (bits 0-3)

	// Read and set the counts
	questionsCount, err := buffer.ReadU16()
	if err != nil {
		return err
	}
	header.SetQuestionsCount(questionsCount)

	answersCount, err := buffer.ReadU16()
	if err != nil {
		return err
	}
	header.SetAnswersCount(answersCount)

	authoritativeEntriesCount, err := buffer.ReadU16()
	if err != nil {
		return err
	}
	header.SetAuthoritativeEntriesCount(authoritativeEntriesCount)

	resourceEntriesCount, err := buffer.ReadU16()
	if err != nil {
		return err
	}
	header.SetResourceEntriesCount(resourceEntriesCount)

	return nil
}
