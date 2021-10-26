package model

type CodeInfo struct {
	Code
	Dt
}

type Dt struct {
	LastDate string
}

type Code struct {
	Id        int
	Code      string
	Code_type int
}
