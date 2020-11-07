package morse

import (
	"log"
	"strings"
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
func Send(out Output, message string) {
	for _, ch := range strings.ToUpper(message) {
		switch ch {
		case ' ':
			log.Println("SPC")
			time.Sleep(tInterword)
		case '.':
			log.Println("STOP")
			time.Sleep(tStop)
		default:
			sendChar(out, ch)
		}
	}
}

func sendSymbol(out Output, t time.Duration) {
	out.On()
	time.Sleep(t)
	out.Off()
	time.Sleep(tDit)
}

func sendChar(out Output, ch rune) {
	if code, found := Code[ch]; found {
		log.Printf("%c %q\n", ch, code)
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
	// characters that are not in Code are ignored
}
