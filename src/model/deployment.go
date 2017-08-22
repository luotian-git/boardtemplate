package model

type Deployment struct {
	DeploymentName string
	Replicas       int
	NFSvolumeList  []NFSvolume
	ContainerList  []Container
}

type NFSvolume struct {
	VolumeName string
	NFSServer  string
	VolumePath string
}

type Container struct {
	ContainerName  string
	BaseImage      string
	Workingdir     string
	Command        string
	ContainerCPU   string
	ContainerMem   string
	ContainerPorts []int
	VolumeList     []ContainerVolume
	EnvList        []ContainerEnv
}

type ContainerEnv struct {
	EnvName  string
	EnvValue string
}

type ContainerVolume struct {
	StorageName  string
	ContainerDir string
}
