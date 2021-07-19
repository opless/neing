package neing

import (
	"bufio"
	"errors"
	"strings"
)

type Protocol struct {
	reader bufio.Reader
	writer bufio.Writer
	isServer bool
	maxSize uint32
	myMaxSize uint32
	buffer []uint8
	bufferSize uint32
	// POSS: Job Queue
	// POSS: FIDs Map
}



func NewProtocol(reader *bufio.Reader, writer *bufio.Writer, isServer bool, maxSize uint32) *Protocol {
	return &Protocol{
		reader: *reader,
		writer: *writer,
		isServer: isServer,
		maxSize: maxSize,
		myMaxSize: maxSize,
		buffer: make([]uint8, maxSize),
	}
}

func (p *Protocol) Initialise()  {
	p.maxSize = p.myMaxSize
}

func (p *Protocol) ReadByteFromInput() (uint8,error) {
	var b, e = p.reader.ReadByte()
	return b, e
}
func (p *Protocol) ReadUInt16FromInput() (uint16, error) {
	var h,e = p.ReadByteFromInput()
	if e != nil {
		return 0, e
	}
	var l, f = p.ReadByteFromInput()
	if f != nil {
		return 0, f
	}
	var ret = (uint16(h) << 8) | uint16(l)
	return ret, nil
}

func (p Protocol) ReadUInt32FromInput() (uint32, error) {
	var ret uint32 = 0
	for i:=0;i <4;i++ {
		var x, e = p.ReadByteFromInput()
		if e != nil {
			return 0, e
		}
		ret = (ret << 8) | uint32(x)
	}
	return ret, nil
}

func (p *Protocol) ReadUInt8(ptr uint32) uint8 {
	return p.buffer[ptr]
}

func (p *Protocol) ReadUInt16(ptr uint32) uint16 {
	return (uint16(p.ReadUInt8(ptr+0)) << 8) |
		(uint16(p.ReadUInt8(ptr+1)))
}

func (p Protocol) ReadUInt32(ptr uint32) uint32 {
	return ((uint32(p.ReadUInt8(ptr+0))) << 24) |
		((uint32(p.ReadUInt8(ptr+1))) << 16) |
		((uint32(p.ReadUInt8(ptr+2))) << 8) |
		(uint32(p.ReadUInt8(ptr+3)))
}

func (p *Protocol) ReadString(ptr uint32) (string,uint32,error)  {
	size := uint32(p.ReadUInt16(ptr))
	if (ptr + size) >= p.bufferSize {
		return "", 0, errors.New(StringTooLong)
	}
	str := string(p.buffer[ptr+2:(ptr+2+size)])

	return str, size + 2, nil
}

func (p *Protocol) WriteUInt8(ptr uint32, value uint8) {
	p.buffer[ptr] = value
}

func (p *Protocol) WriteUInt16(ptr uint32, value uint16) {
	p.buffer[ptr] = uint8(value >> 8)
	p.buffer[ptr+1] = uint8(value & 0xFF)
}

func (p *Protocol) WriteUInt32(ptr uint32, value uint32) {
	p.buffer[ptr] = uint8((value >> 24) & 0xFF)
	p.buffer[ptr+1] = uint8((value >> 16) & 0xFF)
	p.buffer[ptr+2] = uint8((value >> 8) & 0xFF)
	p.buffer[ptr+3] = uint8(value & 0xFF)
}

func (p *Protocol) WriteString(ptr uint32, value string) (uint32,error) {
	octets := []uint8(value)
	size := uint16(len(octets))
	if (ptr+uint32(size)+2) >= p.maxSize {
		return 0, errors.New(StringTooLong)
	}
	var i uint32
	p.WriteUInt16(ptr,size)
	ptr += 2
	for i = 0;i < uint32(size);i ++ {
		p.buffer[ptr+i] = octets[i]
	}
	return uint32(size)+2, nil
}
func (p *Protocol) SendMessage() error {
	var i uint32
	for i =0; i < p.bufferSize; i++ {
		e := p.writer.WriteByte(p.buffer[i])
		if e != nil {
			return e
		}
	}
	return p.writer.Flush() // send it
}

func (p *Protocol) ReadMessage() error {
	// read size, then type, and then hand off to appropriate handler.
	var e error
	var size uint32
	var verb uint8
	var tag uint16

	// read size (including size)
	size, e = p.ReadUInt32FromInput()
	if e != nil {
		return e
	}
	// make sure the message can fit in our (agreed) buffer
	if size >= p.maxSize {
		// we will drop the connection of a too big message.
		return errors.New(MessageTooBig)
	}
	// read verb
	verb , e = p.ReadByteFromInput()
	if e == nil {
		return e
	}
	// gate that verb
	var gate = verb %2
	if gate == 0 && !p.isServer {
		return errors.New(NotAServer)
	}
	if gate == 1 && p.isServer {
		return errors.New(NotAClient)
	}

	// read tag
	tag, e = p.ReadUInt16FromInput()
	if tag != NoTag {
		return errors.New(ExpectingNoTag)
	}

	// read into buffer
	p.bufferSize = size - 7
	var i uint32
	for i = 0; i < p.bufferSize; i ++ {
		p.buffer[i], e= p.ReadByteFromInput()
		if e != nil {
			return e
		}
	}

	// offload
	switch verb {
	case TAttach:
		return errors.New(NotImplemented)
	case TAuth:
		return errors.New(NotImplemented)
	case TClunk:
		return errors.New(NotImplemented)
	case TCreate:
		return errors.New(NotImplemented)
	case TError:
		return errors.New(NotImplemented)
	case TFlush:
		return errors.New(NotImplemented)
	case TOpen:
		return errors.New(NotImplemented)
	case TRead:
		return errors.New(NotImplemented)
	case TRemove:
		return errors.New(NotImplemented)
	case TStat:
		return errors.New(NotImplemented)
	case TVersion:
		return p.ServerNegotiateVersion()
	case TWalk:
		return errors.New(NotImplemented)
	case TWrite:
		return errors.New(NotImplemented)
	case RAttach:
		return errors.New(NotImplemented)
	case RAuth:
		return errors.New(NotImplemented)
	case RClunk:
		return errors.New(NotImplemented)
	case RCreate:
		return errors.New(NotImplemented)
	case RError:
		return errors.New(NotImplemented)
	case RFlush:
		return errors.New(NotImplemented)
	case ROpen:
		return errors.New(NotImplemented)
	case RRead:
		return errors.New(NotImplemented)
	case RRemove:
		return errors.New(NotImplemented)
	case RStat:
		return errors.New(NotImplemented)
	case RVersion:
		return errors.New(NotImplemented)
	case RWalk:
		return errors.New(NotImplemented)
	case RWrite:
		return errors.New(NotImplemented)
	default:
		return errors.New(NotImplemented)
	}
}

func (p Protocol) ServerNegotiateVersion() error {
	// everything is cleared out on a TVersion message
	p.Initialise()
	// unpack mSize
	mSizeRequested := p.ReadUInt32(0)
	// unpack version string
	versionString, _, e := p.ReadString(4)
	if e != nil {
		return e
	}
	if !strings.HasPrefix(versionString,"9P2000") {
		return errors.New(Not9P2000)
	}
	if mSizeRequested < p.myMaxSize {
		p.maxSize = mSizeRequested
	}

	return p.SendRVersion(p.maxSize,"9P2000")
}

func (p *Protocol) ClientNegotiateVersion() error {
	p.Initialise()
	return p.SendTVersion(p.maxSize,"9P2000")
}

func (p *Protocol) SendTVersion(maxSize uint32,version string) error {
	size := uint32(11)
	p.WriteUInt8(4,TVersion)
	p.WriteUInt16(5,NoTag)
	p.WriteUInt32(7,maxSize)
	c,e := p.WriteString(11,version)
	if e != nil {
		return e
	}
	size += c
	p.WriteUInt32(0,size)
	p.bufferSize = size
	return p.SendMessage()
	// 0123 4 56 789a b
	// mlen V tt msiz string...
}

func (p *Protocol) SendRVersion(maxSize uint32, version string) error {
	return p.SendTVersion(maxSize,version)
}


