# morse code package for go
this package can interface with any led, buzzer, etc. to send morse code
messages.  You just need to implement the interface:
``` go
interface {
    On()
    Off()
}
```
and provide it and the bytes to transmit to Send()
