package main

import "fmt"

func main()  {
	m := newMessage("Welcome to boulder Gopers")
	g := newGreeter(m)
	e:= newEvent(g)

	e.start()
}
func newMessage(msg string) message{
	return message {value:msg}
}
type message struct {
	value string
}
func (m message) string() string {
	return fmt.Sprintf "[OFFICIAL MESSAGE]: %s", m.value)
}
func newGreeter(m message) greeter {
	return greeter{message: m}
}

type greeter struct{
	message message
}
func (g greeter) greet() string{
	return g.message.String()
}
func newEvent (g greeter) event {
	return event{greeter: g}
}
type event struct {
	greeter greeter
}
func (e event) start() string{
	return e.greeter.greet()
}