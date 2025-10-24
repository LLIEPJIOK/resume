package docs

import (
	"strings"

	"github.com/LLIEPJIOK/resume/internal/domain/mydate"
)

type Resume struct {
	FirstName      string        `json:"first_name"`
	LastName       string        `json:"last_name"`
	Position       string        `json:"position"`
	Skills         []string      `json:"skills"`
	Projects       []string      `json:"projects"`
	Education      string        `json:"education"`
	LanguageSkills string        `json:"language_skills"`
	Experience     Experience    `json:"experience"`
	WorkHistory    []WorkHistory `json:"work_history"`
}

type Experience struct {
	Years              int      `json:"years"`
	Technologies       []string `json:"technologies"`
	Databases          []string `json:"databases"`
	DevOps             []string `json:"dev_ops"`
	CollaborationTools []string `json:"collaboration_tools"`
	VersionControls    []string `json:"version_control"`
}

type WorkHistory struct {
	Start            mydate.Date `json:"start"`
	End              mydate.Date `json:"end"`
	Role             string      `json:"role"`
	Project          string      `json:"project"`
	Responsibilities []string    `json:"responsibilities"`
	Technologies     []string    `json:"technologies"`
}

func (r *Resume) SetFullName(fullName string) *Resume {
	l := strings.Split(strings.TrimSpace(fullName), " ")
	r.FirstName = l[0]
	r.LastName = l[1]

	return r
}

func (r *Resume) SetPosition(position string) *Resume {
	r.Position = strings.TrimSpace(position)
	return r
}

func (r *Resume) SetSkills(skills []string) *Resume {
	r.Skills = skills
	return r
}

func (r *Resume) SetProjects(projects []string) *Resume {
	r.Projects = projects
	return r
}

func (r *Resume) SetEducation(education string) *Resume {
	r.Education = strings.TrimSpace(education)
	return r
}

func (r *Resume) SetLanguageSkills(languageSkills string) *Resume {
	r.LanguageSkills = strings.TrimSpace(languageSkills)
	return r
}

func (r *Resume) SetExperienceYears(years int) *Resume {
	r.Experience.Years = years
	return r
}

func (r *Resume) SetTechnologies(technologies []string) *Resume {
	r.Experience.Technologies = technologies
	return r
}

func (r *Resume) SetDatabases(databases []string) *Resume {
	r.Experience.Databases = databases
	return r
}

func (r *Resume) SetDevOps(devOps []string) *Resume {
	r.Experience.DevOps = devOps
	return r
}

func (r *Resume) SetCollaborationTools(collaborationTools []string) *Resume {
	r.Experience.CollaborationTools = collaborationTools
	return r
}

func (r *Resume) SetVersionControls(versionControls []string) *Resume {
	r.Experience.VersionControls = versionControls
	return r
}

func (r *Resume) AddWorkHistory(wh WorkHistory) *Resume {
	r.WorkHistory = append(r.WorkHistory, wh)
	return r
}
