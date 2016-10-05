package agent

// Agent lists the characteristics of each agent that needs to be implemented.
type Agent interface {
	Belief()
	Desire()
	Intention()

	Plans()
	Actions()
	Goals()
	Events()
}

// NewAgent ...
type NewAgent struct {
	Agent Agent
	name  string
}
