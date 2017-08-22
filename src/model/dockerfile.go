package model

type Dockerfile struct {
	Base     string
	CopyFrom string
	CopyTo   string
	Command  string
}
