package cards

type NewCardEvent struct {
	NewCard Card
}

func (e NewCardEvent) Name() string {
	return "NewCardEvent"
}
