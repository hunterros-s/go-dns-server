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

// Bytes returns a copy of the buffer's data up to the current position.
func (pb *PacketBuffer) Bytes() []byte {
	return append([]byte{}, pb.buffer[:pb.position]...)
}

// pos returns the current position in the buffer
func (pb *PacketBuffer) pos() ByteIndex {
	return pb.position
}
func (pb *PacketBuffer) Pos() uint16 {
	return uint16(pb.pos())
}

// step moves the current position forward by the specified number of steps
func (pb *PacketBuffer) step(steps ByteIndex) {
	pb.position += steps
}
func (pb *PacketBuffer) Step(steps uint16) {
	pb.step(ByteIndex(steps))
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

// Write writes a single byte to the buffer at the current position and advances the position.
// Returns an error if the position is beyond the buffer's end.
func (pb *PacketBuffer) WriteByte(val byte) error {
	if pb.position >= MAX_BUFFER_SIZE {
		return fmt.Errorf("end of buffer")
	}
	pb.buffer[pb.position] = val
	pb.position++
	return nil
}

func (pb *PacketBuffer) Set(pos uint16, val uint8) error {
	if pos >= MAX_BUFFER_SIZE {
		return fmt.Errorf("out of bounds position")
	}
	pb.buffer[ByteIndex(pos)] = val

	return nil
}

func (pb *PacketBuffer) SetU16(pos uint16, val uint16) error {
	err := pb.Set(pos, uint8(val>>8))
	if err != nil {
		return err
	}
	err = pb.Set(pos+1, uint8(val&0xFF))
	if err != nil {
		return err
	}

	return nil
}

// WriteU8 writes a single byte (u8) to the buffer.
// Returns an error if writing would go beyond the buffer's end.
func (pb *PacketBuffer) WriteU8(val uint8) error {
	return pb.WriteByte(byte(val))
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

// WriteU16 writes two bytes (u16) to the buffer in big-endian order.
// Returns an error if writing would go beyond the buffer's end.
func (pb *PacketBuffer) WriteU16(val uint16) error {
	if err := pb.WriteByte(byte(val >> 8)); err != nil {
		return err
	}
	if err := pb.WriteByte(byte(val & 0xFF)); err != nil {
		return err
	}
	return nil
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

// WriteU32 writes four bytes (u32) to the buffer in big-endian order.
// Returns an error if writing would go beyond the buffer's end.
func (pb *PacketBuffer) WriteU32(val uint32) error {
	if err := pb.WriteByte(byte((val >> 24) & 0xFF)); err != nil {
		return err
	}
	if err := pb.WriteByte(byte((val >> 16) & 0xFF)); err != nil {
		return err
	}
	if err := pb.WriteByte(byte((val >> 8) & 0xFF)); err != nil {
		return err
	}
	if err := pb.WriteByte(byte(val & 0xFF)); err != nil {
		return err
	}
	return nil
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

// WriteQName writes a domain name (QNAME) to the buffer. It splits the QNAME into labels and writes each label's length and bytes.
// Returns an error if any label exceeds 63 characters or if writing would go beyond the buffer's end.
func (pb *PacketBuffer) WriteQName(qname string) error {
	labels := strings.Split(qname, ".")
	for _, label := range labels {
		len := len(label)
		if len > 0x3F {
			return fmt.Errorf("single label exceeds 63 characters of length")
		}

		// Write the length of the label
		if err := pb.WriteU8(uint8(len)); err != nil {
			return err
		}

		// Write the bytes of the label
		for i := 0; i < len; i++ {
			if err := pb.WriteU8(label[i]); err != nil {
				return err
			}
		}
	}

	// Write the null byte to terminate the QNAME
	if err := pb.WriteU8(0); err != nil {
		return err
	}

	return nil
}
