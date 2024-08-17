package buffer

import (
	"fmt"
	"strings"
)

const (
	MAX_BUFFER_SIZE = 512
	MaxJumps        = 5
)

// ByteIndex represents a position within a packet buffer
type ByteIndex uint16

// PacketBuffer represents a fixed-size buffer for handling network packets
type PacketBuffer struct {
	buffer   [MAX_BUFFER_SIZE]byte // Fixed-size buffer to store packet data
	position ByteIndex             // Current position in the buffer
}

// NewPacketBuffer creates and initializes a new PacketBuffer
func NewPacketBuffer(b []byte) *PacketBuffer {
	pb := &PacketBuffer{}
	copy(pb.buffer[:], b)
	return pb
}

// pos returns the current position in the buffer
func (pb *PacketBuffer) pos() ByteIndex {
	return pb.position
}

// step moves the current position forward by the specified number of steps
func (pb *PacketBuffer) step(steps ByteIndex) {
	pb.position += steps
}

// seek sets the current position to the specified value
func (pb *PacketBuffer) seek(pos ByteIndex) {
	pb.position = pos
}

// read returns the byte at the current position and advances the position
// Returns an error if the position is beyond the buffer's end
func (pb *PacketBuffer) read() (byte, error) {
	if pb.position >= MAX_BUFFER_SIZE {
		return 0x0, fmt.Errorf("end of buffer")
	}
	result := pb.buffer[pb.position]
	pb.position += 1
	return result, nil
}

// get returns the byte at the specified position without advancing the current position
// Returns an error if the position is beyond the buffer's end
func (pb *PacketBuffer) get(pos ByteIndex) (byte, error) {
	if pos >= MAX_BUFFER_SIZE {
		return 0x0, fmt.Errorf("end of buffer")
	}
	return pb.buffer[pos], nil
}

// get_range returns a slice of bytes from the buffer, starting at 'start' and of length 'len'
// Returns an error if the requested range extends beyond the buffer's end
func (pb *PacketBuffer) get_range(start, len ByteIndex) ([]byte, error) {
	if start+len >= MAX_BUFFER_SIZE {
		return []byte{}, fmt.Errorf("end of buffer")
	}
	return pb.buffer[start : start+len], nil
}

// read_u16 reads two bytes from the current position, interprets them as a big-endian uint16, and advances the position
// Returns an error if reading would go beyond the buffer's end
func (pb *PacketBuffer) ReadU16() (uint16, error) {
	if pb.position+2 >= MAX_BUFFER_SIZE {
		return 0, fmt.Errorf("end of buffer")
	}

	byte1 := uint16(pb.buffer[pb.position])
	byte2 := uint16(pb.buffer[pb.position+1])

	result := byte1<<8 | byte2
	pb.step(2)

	return result, nil
}

// read_u32 reads four bytes from the current position, interprets them as a big-endian uint32, and advances the position
// Returns an error if reading would go beyond the buffer's end
func (pb *PacketBuffer) ReadU32() (uint32, error) {
	if pb.position+4 >= MAX_BUFFER_SIZE {
		return 0, fmt.Errorf("end of buffer")
	}

	byte1 := uint32(pb.buffer[pb.position])
	byte2 := uint32(pb.buffer[pb.position+1])
	byte3 := uint32(pb.buffer[pb.position+2])
	byte4 := uint32(pb.buffer[pb.position+3])

	pb.step(4)

	return byte1<<24 | byte2<<16 | byte3<<8 | byte4, nil
}

// read_qname reads a domain name from the buffer, handling DNS name compression
// It returns the domain name as a string and any error encountered
func (pb *PacketBuffer) ReadQName() (string, error) {
	pos := pb.pos()
	jumped := false
	const maxJumps = 5
	jumpsPerformed := 0
	outstring := ""
	delim := ""

	for {
		// Prevent infinite loops due to malformed packets
		if jumpsPerformed > maxJumps {
			return outstring, fmt.Errorf("limit of %d jumps exceeded", maxJumps)
		}

		// Read the length byte of the current label
		len, err := pb.get(pos)
		if err != nil {
			return outstring, err
		}

		// Check if this is a pointer (DNS name compression)
		if (len & 0xC0) == 0xC0 {
			// Update buffer position if this is the first jump
			if !jumped {
				pb.seek(pos + 2)
			}

			// Calculate the offset for the jump
			b2, err := pb.get(pos + 1)
			if err != nil {
				return outstring, err
			}
			offset := ((uint16(len) ^ 0xC0) << 8) | uint16(b2)
			pos = ByteIndex(offset)

			jumped = true
			jumpsPerformed++

			continue
		}
		// Handle a normal label
		pos++
		if len == 0 {
			break // End of domain name
		}

		// Append the delimiter (empty for the first label, "." for subsequent ones)
		outstring += delim

		// Read the label and append it to the output string
		label, err := pb.get_range(pos, ByteIndex(len))
		if err != nil {
			return outstring, err
		}
		outstring += strings.ToLower(string(label))

		delim = "."
		pos += ByteIndex(len)
	}

	// Update the buffer position if we didn't jump
	if !jumped {
		pb.seek(pos)
	}

	return outstring, nil
}
