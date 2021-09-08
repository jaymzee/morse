package morse

import (
	"fmt"
	"log"
	"time"
)

// Output is an interface for any two state output such as led, buzzer, etc
type Output interface {
	On()
	Off()
}

// Sender sends morse code to an Output
type Sender struct {
	Out       Output
	Dit       time.Duration
	Dah       time.Duration
	Interchar time.Duration
	Interword time.Duration
	Stop      time.Duration
}

// NewSender returns a new Sender
func NewSender(out Output) *Sender {
	return &Sender{
		Out:       out,
		Dit:       100 * time.Millisecond,
		Dah:       300 * time.Millisecond,
		Interchar: 200 * time.Millisecond,
		Interword: 400 * time.Millisecond,
		Stop:      1000 * time.Millisecond,
	}
}

// Write implements io.Writer
// and transmits the message on the output in morse code and returns
// the number of bytes sent successfully.
func (s *Sender) Write(bytes []byte) (int, error) {
	for n, b := range bytes {
		switch b {
		case ' ':
			log.Println("SPC")
			time.Sleep(s.Interword)
		case '.':
			log.Println("STOP")
			time.Sleep(s.Stop)
		default:
			if err := s.WriteByte(b); err != nil {
				return n, err
			}
		}
	}
	return len(bytes), nil
}

// WriteByte transmits the character on the output in morse code and returns
func (s *Sender) WriteByte(b byte) error {
	if code, found := Code[b]; found {
		log.Printf("%c %q\n", b, code)
		for _, sym := range code {
			switch sym {
			case '.':
				s.sendSymbol(s.Dit)
			case '-':
				s.sendSymbol(s.Dah)
			default:
				panic("only . and - allowed in morse code table!")
			}
		}
		time.Sleep(s.Interchar)
		return nil
	}
	return fmt.Errorf("cannot translate byte %#02x to morse code", b)
}

func (s *Sender) sendSymbol(t time.Duration) {
	s.Out.On()
	time.Sleep(t)
	s.Out.Off()
	time.Sleep(s.Dit)
}
