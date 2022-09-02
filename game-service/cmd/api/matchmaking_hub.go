package main

type match_message struct {
	data []byte
	room string
}

type match_subscription struct {
	conn *match_conn
}

type hub struct {
	rooms      map[string]map[*match_conn]bool
	broadcast  chan match_message
	register   chan match_subscription
	unregister chan match_subscription
}

var match_h = hub{
	broadcast:  make(chan match_message),
	register:   make(chan match_subscription),
	unregister: make(chan match_subscription),
	rooms:      make(map[string]map[*match_conn]bool),
}

func (h *hub) run() {}