package model

type Project struct {
	ProjectId   string `json:"projectId"`
	ProjectName string `json:"projectName"`
}

type Program struct {
	ProgramId      string    `json:"programId"`
	ProgramName    string    `json:"programName"`
	Project        []Project `json:"projects"`
	ProgramManager string    `json:"programmanager"`
}
type AddProject struct {
	ProgramName string `json:"programName"`
	ProjectId   string `json:"projectId"`
	ProjectName string `json:"projectName"`
}
