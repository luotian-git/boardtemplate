package model

type Dockerfile struct {
	Base       string
	Author     string
	CopyFrom   string
	CopyTo     string
	EntryPoint string
	Command    string
}
