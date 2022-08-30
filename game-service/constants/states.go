package states

type State int

const (
	Pending State = iota
	Running
	Completed
	Cancelled
)

func (s State) Run() {
	if (s == Pending) {
		s = Running;
	}
}

func (s State) Complete() {
	if (s == Running) {
		s = Completed;
	}
}

func (s State) Cancel() {
	if (s == Running) {
		s = Cancelled;
	}
}