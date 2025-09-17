package chores

type NewChoreEvent struct {
	NewChore Chore
}

func (e NewChoreEvent) Name() string {
	return "NewChoreEvent"
}

type CompletedChoreEvent struct {
	ChoreId uint
}

func (e CompletedChoreEvent) Name() string {
	return "CompletedChoreEvent"
}
