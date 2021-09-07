package morse

import (
	"fmt"
	"log"
	"time"
)

const (
	tDit       = 100 * time.Millisecond
	tDah       = 300 * time.Millisecond
	tInterchar = 200 * time.Millisecond
	tInterword = 400 * time.Millisecond
	tStop      = 1000 * time.Millisecond
)

// Output is an interface for any two state output such as led, buzzer, etc
type Output interface {
	On()
	Off()
}

// Send transmits the message on the output in morse code and returns
// the number of bytes translated successfully.
func Send(out Output, bytes []byte) (int, error) {
	for n, b := range bytes {
		switch b {
		case ' ':
			log.Println("SPC")
			time.Sleep(tInterword)
		case '.':
			log.Println("STOP")
			time.Sleep(tStop)
		default:
			if err := sendByte(out, b); err != nil {
				return n, err
			}
		}
	}
	return len(bytes), nil
}

func sendSymbol(out Output, t time.Duration) {
	out.On()
	time.Sleep(t)
	out.Off()
	time.Sleep(tDit)
}

func sendByte(out Output, b byte) error {
	if code, found := Code[b]; found {
		log.Printf("%c %q\n", b, code)
		for _, sym := range code {
			switch sym {
			case '.':
				sendSymbol(out, tDit)
			case '-':
				sendSymbol(out, tDah)
			default:
				panic("only . and - allowed in morse code table!")
			}
		}
		time.Sleep(tInterchar)
		return nil
	}
	return fmt.Errorf("cannot translate byte %#02x to morse code", b)
}
