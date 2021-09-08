# morse code package for go
this package can interface with any led, buzzer, etc. to send morse code
messages.  To add morse code capability to the output just implement the
following methods:

``` go
func (buzz *Buzzer) On() { ... }
func (buzz *Buzzer) Off() { ... }
```

The only characters accepted are [A-Z0-9. ].
Send a message to the output using io or fmt:
```go
// create a Sender
sender := morse.NewSender(led)
// pass the Sender to anything that accepts a Writer
io.WriteString(sender, "HELLO WORLD.")
fmt.Fprint(sender, "HELLO WORLD.")
```
