package morse

import (
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

// Output the interface for any two state output such as led, buzzer, etc
type Output interface {
	On()
	Off()
}

// Send transmits the message on the output
func Send(out Output, bytes []byte) {
	for _, b := range bytes {
		switch b {
		case ' ':
			log.Println("SPC")
			time.Sleep(tInterword)
		case '.':
			log.Println("STOP")
			time.Sleep(tStop)
		default:
			sendByte(out, b)
		}
	}
}

func sendSymbol(out Output, t time.Duration) {
	out.On()
	time.Sleep(t)
	out.Off()
	time.Sleep(tDit)
}

func sendByte(out Output, b byte) {
	if code, found := Code[b]; found {
		log.Printf("%c %q\n", b, code)
		for _, sym := range code {
			switch sym {
			case '.':
				sendSymbol(out, tDit)
			case '-':
				sendSymbol(out, tDah)
			default:
				panic("invalid character in morse code table")
			}
		}
		time.Sleep(tInterchar)
	}
	// bytes that are not in the map Code are ignored
}
