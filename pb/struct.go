package pb

type Trans struct {
	FromID   int64
	TargetID int64
	Message  []byte
	Block    BlockInfo //  区块信息，每次转发增加区块信息
}
