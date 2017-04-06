package main

import (
	"github.com/joelbraga/aztek/action"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var ch chan *action.Action

type Object struct {
	name string
	value int
}

func main(){
	go event()
	actionEvent := action.NewActionEvent()
	ch = actionEvent.Data
	actionEvent.AddEvent(action.ADD_MODEL, &Object{
		"XPTO",
		10,
	})

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	<-wait
}

func event() {
	for {
		select {
		case action := <- ch:
			fmt.Println(action.Type)
			fmt.Println(action.Payload)
		}
	}
}
