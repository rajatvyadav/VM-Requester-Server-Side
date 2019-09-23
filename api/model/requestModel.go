package model

type VMRequest struct {
	RequestId      string        `json:"requestId"`
	ProgramManager string        `json:"programManager"`
	ProgramName    string        `json:"programName"`
	ProjectName    string        `json:"projectName"`
	VmType         string        `json:"vmType"`
	VmResources    Vmresources   `json:"vmResources"`
	FirewallSetup  Firewallsetup `json:"firewallSetup"`
	ServerLocation string        `json:"serverLocation"`
	Tags           []string      `json:"tags"`
	ServerName     string        `json:"serverName"`
	PrivateIp      string        `json:"privateIp"`
	PublicIp       string        `json:"publicIp"`
	CurrentStatus  string        `json:"currentStatus"`
}

type Vmresources struct {
	RamSize         string   `json:"ramSize"`
	CpuCores        string   `json:"cpuCores"`
	OperatingSystem string   `json:"operatingSystem"`
	Disk            Diskinfo `json:"disk"`
}

type Diskinfo struct {
	Root    string `json:"root"`
	App     string `json:"app"`
	Data    string `json:"data"`
	Logs    string `json:"logs"`
	Backups string `json:"backups"`
}

type Firewallsetup struct {
	PIP              string   `json:"pip"`
	Ports            Portinfo `json:"ports"`
	InternetRequired string   `json:"internetRequired"`
}
type Portinfo struct {
	ApplicationName string `json:"applicationName"`
	PortNumber      string `json:"portNumber"`
	Direction       string `json:"direction"`
}

// Created global list which will be available throughout the application
var VMRequestList []VMRequest
