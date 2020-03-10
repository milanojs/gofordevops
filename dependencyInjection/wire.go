// +build wireinject
package main 
import "github.com/google/go-cloud/wire"
func initializeEvent() event {
	wire.Build( newMessage, NewGreeter, newEvent)
	return event{}
}