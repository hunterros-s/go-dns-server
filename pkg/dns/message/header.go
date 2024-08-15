package dns

import (
	"github.com/hunterros-s/go-dns-server/pkg/buffer"
	"github.com/hunterros-s/go-dns-server/pkg/dns/enum"
)

type Header struct {
	id uint16 // 16 bits

	response             bool              // 1 bit
	opcode               uint8             // 4 bits
	authoritative_answer bool              // 1 bit
	truncated_message    bool              // 1 bit
	recursion_desired    bool              // 1 bit
	recursion_available  bool              // 1 bit
	z                    bool              // 1 bit
	rescode              enum.ResponseCode // 4 bits

	questions             uint16 // 16 bits
	answers               uint16 // 16 bits
	authoritative_entries uint16 // 16 bits
	resource_entries      uint16 // 16 bits
}

func NewHeader() *Header {
	return &Header{
		id: 0,

		response:             false,
		opcode:               0,
		authoritative_answer: false,
		truncated_message:    false,
		recursion_desired:    false,
		recursion_available:  false,
		z:                    false,
		rescode:              enum.NOERROR,

		questions:             0,
		answers:               0,
		authoritative_entries: 0,
		resource_entries:      0,
	}
}

func (h *Header) Read(buffer *buffer.PacketBuffer) error {
	var err error

	// Read the fixed-size fields
	h.id, err = buffer.ReadU16()
	if err != nil {
		return err
	}

	flags, err := buffer.ReadU16()
	if err != nil {
		return err
	}

	// Extract bit fields from the flags
	h.response = (flags & 0x8000) != 0              // QR (bit 15)
	h.opcode = uint8((flags >> 11) & 0x0F)          // Opcode (bits 11-14)
	h.authoritative_answer = (flags & 0x0400) != 0  // AA (bit 10)
	h.truncated_message = (flags & 0x0200) != 0     // TC (bit 9)
	h.recursion_desired = (flags & 0x0100) != 0     // RD (bit 8)
	h.recursion_available = (flags & 0x0080) != 0   // RA (bit 7)
	h.z = (flags & 0x0070) != 0                     // Reserved (bits 4-6)
	h.rescode = enum.ResponseCode((flags & 0x000F)) // RCODE (bits 0-3)

	// Read fixed-size fields from the buffer
	h.questions, err = buffer.ReadU16()
	if err != nil {
		return err
	}

	h.answers, err = buffer.ReadU16()
	if err != nil {
		return err
	}

	h.authoritative_entries, err = buffer.ReadU16()
	if err != nil {
		return err
	}

	h.resource_entries, err = buffer.ReadU16()
	if err != nil {
		return err
	}

	return nil
}
