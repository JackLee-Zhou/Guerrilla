package rpc

type Trans struct {
	FromID   int64
	TargetID int64
	Message  []byte
	//	加入BlockInfo 每次转发增加区块信息
}
