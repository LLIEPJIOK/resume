package docs

import (
	"fmt"
	"strings"
	"unicode/utf8"

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

const (
	firstName      = "first_name"
	lastName       = "last_name"
	position       = "position"
	skills         = "skills"
	projects       = "projects"
	education      = "education"
	languageSkills = "language_skills"

	experience     = "experience.years"
	experienceTech = "experience.technologies"
	experienceDB   = "experience.databases"
	experienceDev  = "experience.dev_ops"
	experienceColl = "experience.collaboration_tools"
	experienceVC   = "experience.version_control"

	workHistory       = "work_history"
	workHistoryPeriod = "work_history.period"
	workHistoryRole   = "work_history.role"
	workHistoryProj   = "work_history.project"
	workHistoryResp   = "work_history.responsibilities"
	workHistoryTech   = "work_history.technologies"
)

func (r *Resume) Validate() (errs map[string]string) {
	errs = make(map[string]string)
	errs = r.validateTopFields(errs)
	errs = r.validateExperience(errs)
	errs = r.validateWorkHistory(errs)

	return errs
}

func (r *Resume) validateTopFields(errs map[string]string) map[string]string {
	if r.FirstName == "" {
		errs[firstName] = "first name is empty"
	}

	if r.LastName == "" {
		errs[lastName] = "last name is empty"
	}

	if utf8.RuneCountInString(r.LastName) != 2 || r.LastName[len(r.LastName)-1] != '.' {
		errs[lastName] = "last name must contain only 1 letter"
	}

	if r.Position == "" {
		errs[position] = "position is empty"
	}

	if len(r.Skills) == 0 {
		errs[skills] = "skills are empty"
	}

	if len(r.Projects) == 0 {
		errs[projects] = "projects are empty"
	}

	if r.Education == "" {
		errs[education] = "education is empty"
	}

	if r.LanguageSkills == "" {
		errs[languageSkills] = "language skills are empty"
	}

	return errs
}

func (r *Resume) validateExperience(errs map[string]string) map[string]string {
	errs = r.validateExperienceYears(errs)

	if len(r.Experience.Technologies) == 0 {
		errs[experienceTech] = "technologies are empty"
	}

	if len(r.Experience.Databases) == 0 {
		errs[experienceDB] = "databases are empty"
	}

	if len(r.Experience.DevOps) == 0 {
		errs[experienceDev] = "dev ops are empty"
	}

	if len(r.Experience.CollaborationTools) == 0 {
		errs[experienceColl] = "collaboration tools are empty"
	}

	if len(r.Experience.VersionControls) == 0 {
		errs[experienceVC] = "version control are empty"
	}

	return errs
}

func (r *Resume) validateExperienceYears(errs map[string]string) map[string]string {
	var minDate *mydate.Date

	for _, wh := range r.WorkHistory {
		if minDate == nil || wh.Start.Less(*minDate) {
			minDate = &wh.Start
		}
	}

	if minDate == nil {
		return errs
	}

	current := mydate.Current()
	totalMonths := current.Since(*minDate)

	if (totalMonths-1)/12+1 < r.Experience.Years || r.Experience.Years < (totalMonths)/12 {
		errs[experience] = "experience years do not match work history"
	}

	return errs
}

func (r *Resume) validateWorkHistory(errs map[string]string) map[string]string {
	if len(r.WorkHistory) == 0 {
		errs[workHistory] = "work history is empty"
		return errs
	}

	for i, wh := range r.WorkHistory {
		if wh.Start.Current() {
			errs[fmt.Sprintf("%s.%d", workHistoryPeriod, i+1)] = "start date cannot be current"
		}

		if wh.End.Less(wh.Start) {
			errs[fmt.Sprintf("%s.%d", workHistoryPeriod, i+1)] = "end date is before start date"
		}

		if wh.Role == "" {
			errs[fmt.Sprintf("%s.%d", workHistoryRole, i+1)] = "role is empty"
		}

		if wh.Project == "" {
			errs[fmt.Sprintf("%s.%d", workHistoryProj, i+1)] = "project is empty"
		}

		if len(wh.Responsibilities) == 0 {
			errs[fmt.Sprintf("%s.%d", workHistoryResp, i+1)] = "responsibilities are empty"
		}

		if len(wh.Technologies) == 0 {
			errs[fmt.Sprintf("%s.%d", workHistoryTech, i+1)] = "technologies are empty"
		}
	}

	return errs
}
