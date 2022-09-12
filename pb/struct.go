package pb

type Trans struct {
	FromID   int64
	TargetID int64
	Message  []byte
	Block    BlockInfo //  BlockInfo，each transfer will add a new block to the origin chain
}

// --------------- Const Event --------------------

type SystemEvent string
type UserEvent string
type BlockEvent string
