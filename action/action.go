package action

const ADD_MODEL string = "ADD_MODEL"
const UPDATE_MODEL string = "UPDATE_MODEL"
const DELETE_MODEL string = "DELETE_MODEL"

type Action struct {
	Type string
	Payload interface{}
}

type ActionEvent struct {
	Data chan *Action
}

func NewActionEvent() *ActionEvent {
	return &ActionEvent{
		Data: make(chan *Action),
	}
}

func(ae *ActionEvent) Add(action *Action) {
	ae.Data <- action
}

func(ae *ActionEvent) AddEvent(event string, payload interface{}) {
	ae.Add(&Action{
		Type: event,
		Payload: payload,
	})
}
