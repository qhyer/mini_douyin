package fsm

// RelationState is relation state type
type RelationState string

// RelationEvent is relation event type
type RelationEvent string

// consts for relation
var (
	// states
	StateNoRelation = RelationState("no_relation")
	StateFollowing  = RelationState("following")
	StateFriend     = RelationState("friend")
	// events
	EventAddFollowing = RelationEvent("add_following")
	EventDelFollowing = RelationEvent("del_following")
	EventBeFriend     = RelationEvent("be_friend")
)

// RelationEventHandler is used to handle all state change for relation
type RelationEventHandler interface {
	AddFollowing(*Event)
	DelFollowing(*Event)
}

// DefaultHandler is the default RelationEventHandler
var DefaultHandler = &defaultHandlerImpl{}

type defaultHandlerImpl struct{}

func (*defaultHandlerImpl) AddFollowing(*Event) {}
func (*defaultHandlerImpl) DelFollowing(*Event) {}

// RelationStateMachine is used to describe all state change for relation
type RelationStateMachine struct {
	*FSM
}

// NewRelationStateMachine will create a RelationStateMachine
func NewRelationStateMachine(initial RelationState, handler RelationEventHandler) *RelationStateMachine {
	rs := &RelationStateMachine{
		FSM: NewFSM(
			string(StateNoRelation),
			Events{
				{
					Name: string(EventAddFollowing),
					Src: []string{
						string(StateNoRelation),
					},
					Dst: string(StateFollowing),
				},
				{
					Name: string(EventDelFollowing),
					Src: []string{
						string(StateFollowing),
						string(StateFriend),
					},
					Dst: string(StateNoRelation),
				},
			},
			Callbacks{
				string(EventAddFollowing): handler.AddFollowing,
				string(EventDelFollowing): handler.DelFollowing,
			},
		),
	}
	return rs
}

// Event is used to execute any events
func (r *RelationStateMachine) Event(event RelationEvent, args ...interface{}) error {
	return r.FSM.Event(string(event), args...)
}

// SetState is used to set state
func (r *RelationStateMachine) SetState(state RelationState) {
	r.FSM.SetState(string(state))
}
