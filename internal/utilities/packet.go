package utilities

import (
	"bytes"
	"encoding/hex"
	"strconv"
	"strings"
)

type Packet struct {
	packet string
	header string
}

func NewPacket() Packet {
	return Packet{}
}

func (p *Packet) GetLen() int {
	return len(p.packet)
}

func (p *Packet) SetPacket(pack string) {
	p.packet = pack
}

func (p *Packet) RemoveHeader() {
	if strings.HasPrefix(p.packet, "<") {
		endArrow := strings.Index(p.packet, ">") + 1
		endSlash := strings.Index(p.packet, "/>") + 2
		if endSlash < endArrow {
			p.packet = p.packet[endSlash:]
		} else {
			p.packet = p.packet[endArrow:]
		}
	} else if strings.HasPrefix(p.packet, "%") {
		packetHandled := strings.Split(p.packet, "%")
		if len(packetHandled) >= 4 {
			totalLen := len(packetHandled[1]) + len(packetHandled[2]) + len(packetHandled[3]) + 3
			p.packet = p.packet[totalLen:]
		}
	}
}

func (p *Packet) AddXML(field, value string, part int) {
	switch part {
	case 0:
		p.packet += "<" + field + ">" + value + "</" + field + ">"
	case 1:
		p.packet += "<" + field + ">"
	case 2:
		p.packet += "</" + field + ">"
	}
}

func (p *Packet) GetXML(field string) string {
	startTag := "<" + field + ">"
	endTag := "</" + field + ">"
	start := strings.Index(p.packet, startTag)
	if start == -1 {
		return ""
	}
	start += len(startTag)
	end := strings.Index(p.packet[start:], endTag)
	if end == -1 {
		return ""
	}
	return p.packet[start : start+end]
}

func (p *Packet) AddXMLSingle(part int, data ...string) {
	var builder strings.Builder
	if part == 2 {
		builder.WriteString("</")
	} else {
		builder.WriteString("<")
	}
	for i := 0; i < len(data); i += 2 {
		if i > 0 {
			builder.WriteString(" ")
		}
		builder.WriteString(data[i])
		if i+1 < len(data) {
			builder.WriteString("='")
			builder.WriteString(data[i+1])
			builder.WriteString("'")
		}
	}
	if part == 0 {
		builder.WriteString(" />")
	} else {
		builder.WriteString(">")
	}
	p.packet += builder.String()
}

func (p *Packet) GetXMLSingle(field string) string {
	start := strings.Index(p.packet, field+"='")
	if start == -1 {
		return ""
	}
	start += len(field) + 2
	end := strings.Index(p.packet[start:], "'")
	if end == -1 {
		return ""
	}
	return p.packet[start : start+end]
}

func (p *Packet) AddCDATA(data string) {
	p.packet += "<![CDATA[" + data + "]]>"
}

func (p *Packet) GetCDATA(str string) string {
	start := strings.Index(str, "CDATA[")
	if start == -1 {
		return ""
	}
	start += len("CDATA[")
	end := strings.Index(str[start:], "]")
	if end == -1 {
		return ""
	}
	return str[start : start+end]
}

func (p *Packet) AddString(s string) {
	p.packet += s
}

func (p *Packet) GetString(start, end int, nulled bool) string {
	if start >= len(p.packet) || end > len(p.packet) {
		return ""
	}
	result := p.packet[start:end]
	p.packet = p.packet[end:]
	if nulled {
		return result
	}
	return p.removeNullByte(result)
}

func (p *Packet) removeNullByte(s string) string {
	return strings.Split(s, "\x00")[0]
}

func (p *Packet) GetAsByte(val, num int) string {
	buf := new(bytes.Buffer)
	for i := 0; i < num; i++ {
		buf.WriteByte(byte((val >> (8 * i)) & 0xff))
	}
	return buf.String()
}

func (p *Packet) AddInt(val int) {
	p.packet += strconv.Itoa(val)
}

func (p *Packet) GetInt(bytec int) int {
	if bytec > len(p.packet) {
		return 0
	}
	s := p.packet[:bytec]
	bytes := []byte(s)
	var hexStr string
	for i := bytec - 1; i >= 0; i-- {
		b := bytes[i]
		hexByte := hex.EncodeToString([]byte{b})
		if len(hexByte) < 2 {
			hexByte = "0" + hexByte
		}
		hexStr += hexByte
	}
	p.packet = p.packet[bytec:]
	val, _ := strconv.ParseInt(hexStr, 16, 32)
	return int(val)
}

func (p *Packet) AddByte(b byte) {
	p.packet += string([]byte{b})
}

func (p *Packet) AddByte2(b1, b2 byte) {
	p.packet += string([]byte{b1, b2})
}

func (p *Packet) AddByte4(b1, b2, b3, b4 byte) {
	p.packet += string([]byte{b1, b2, b3, b4})
}

func (p *Packet) AddByteArray(b []byte) {
	p.packet += string(b)
}

func (p *Packet) GetPacket() string {
	return p.packet
}

func (p *Packet) Clean() {
	p.packet = ""
	p.header = ""
}
