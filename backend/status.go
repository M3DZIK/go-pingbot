package backend

var Status StatusType

type StatusType struct {
	Success int64
	Error   int64
}
