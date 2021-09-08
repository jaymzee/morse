# morse code package for go
this package can interface with any led, buzzer, etc. to send morse code
messages.  To add morse code capability to the output just implement the
following methods:

``` go
func (buzz *Buzzer) On() { ... }
func (buzz *Buzzer) Off() { ... }
func (buzz *Buzzer) Write(xs []byte) (int, error) {
        return morse.Send(buzz, xs)
}
```

The only characters accepted are A thru Z, a period and a space.
Send a message to the output using io or fmt:
```go
io.WriteString(out, "HELLO WORLD.")
fmt.Fprint(out, "HELLO WORLD.")
```
