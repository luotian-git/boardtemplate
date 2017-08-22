package model

type Service struct {
	ServiceName     string
	ExternalService bool
	PortList        []ServicePort
	SelectorList    []string
}

type ServicePort struct {
	Port       int
	TargetPort int
	NodePort   int
}
