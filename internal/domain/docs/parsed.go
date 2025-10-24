package docs

import (
	"log/slog"
	"strconv"
	"strings"

	"github.com/LLIEPJIOK/resume/internal/domain/mydate"
	"github.com/LLIEPJIOK/resume/pkg/strutil"
)

type ParsedDocument struct {
	Title   string        `json:"title"`
	Content []ContentItem `json:"content"`
}

type ContentItem struct {
	Type  contentType  `json:"type"`
	Text  string       `json:"text,omitempty"`
	Table *ParsedTable `json:"table,omitempty"`
}

type ParsedTable struct {
	Rows [][]string `json:"rows"`
}

func validateParsedDocumentStructure(doc *ParsedDocument) error {
	if len(doc.Content) < expectedContentLen {
		return NewErrInvalidLength(expectedContentLen, len(doc.Content))
	}

	for i, c := range doc.Content {
		expectedType := expectedContentType(i)
		if c.Type != expectedType {
			return NewErrInvalidType(expectedType, i)
		}
	}

	return nil
}

func (d *ParsedDocument) ToResume() (*Resume, error) {
	if err := validateParsedDocumentStructure(d); err != nil {
		return nil, err
	}

	r := &Resume{}
	r = r.SetFullName(d.Content[0].Text)
	r = r.SetPosition(d.Content[1].Text)

	for _, row := range d.Content[2].Table.Rows {
		for _, cell := range row {
			switch {
			case strings.HasPrefix(cell, "Навыки"):
				skills := strings.Split(strings.TrimSpace(cell), "\n")
				r = r.SetSkills(skills[1:]) // skip header

			case strings.HasPrefix(cell, "Проекты"):
				projects := strings.Split(strings.TrimSpace(cell), "\n")
				r = r.SetProjects(projects[1:]) // skip header

			case strings.HasPrefix(cell, "Образование"):
				education := strings.TrimSpace(strings.TrimPrefix(cell, "Образование"))
				r = r.SetEducation(education)

			case strings.HasPrefix(cell, "Языковые навыки"):
				languageSkills := strings.TrimSpace(strings.TrimPrefix(cell, "Языковые навыки"))
				r = r.SetLanguageSkills(languageSkills)

			default:
				slog.Warn("Unexpected cell content", slog.String("cell", cell))
			}
		}
	}

	// doc.Content[3] is work experience header, skip it

	for exp := range strings.SplitSeq(d.Content[4].Table.Rows[0][0], "\n") {
		switch {
		case strings.HasPrefix(exp, "Опыт разработки"):
			re := getExperienceRegexp()
			matches := re.FindStringSubmatch(exp)
			if len(matches) > 1 {
				years, err := strconv.Atoi(matches[1])
				if err != nil {
					slog.Warn("Failed to parse years of experience", slog.String("cell", exp))
					continue
				}

				r = r.SetExperienceYears(years)
			}

		case strings.HasPrefix(exp, "Технологии"):
			technologies := strutil.SplitAnyCleanPrefix(exp, ",.", "Технологии:")
			r.SetTechnologies(technologies)

		case strings.HasPrefix(exp, "Базы данных"):
			databases := strutil.SplitAnyCleanPrefix(exp, ",.", "Базы данных:")
			r.SetDatabases(databases)

		case strings.HasPrefix(exp, "DevOps"):
			devOps := strutil.SplitAnyCleanPrefix(exp, ",.", "DevOps:")
			r.SetDevOps(devOps)

		case strings.HasPrefix(exp, "Системы совместной работы"):
			collaborationTools := strutil.SplitAnyCleanPrefix(
				exp,
				",.",
				"Системы совместной работы:",
			)
			r.SetCollaborationTools(collaborationTools)

		case strings.HasPrefix(exp, "Система контроля версий"):
			versionControls := strutil.SplitAnyCleanPrefix(exp, ",.", "Система контроля версий:")
			r.SetVersionControls(versionControls)

		default:
			slog.Warn("Unexpected experience content", slog.String("experience", exp))
		}
	}

	// doc.Content[5] is work history header, skip it

	for _, content := range d.Content[6:] {
		var wh WorkHistory

		for _, row := range content.Table.Rows {
			switch row[0] {
			case "Период":
				start, err := mydate.ExtractAndParseDate(row[1])
				if err != nil {
					slog.Warn("Failed to parse project start date", slog.String("cell", row[1]))
					continue
				}

				end, err := mydate.ExtractAndParseDate(row[2])
				if err != nil {
					slog.Warn("Failed to parse project end date", slog.String("cell", row[2]))
					continue
				}

				wh.Start = start
				wh.End = end

			case "Роли проекта", "Роль на проекте":
				wh.Role = row[1]

			case "Проект":
				wh.Project = row[1]

			case "Обязанности и достижения":
				responsibilities := strutil.SplitAnyClean(row[1], "\n")
				wh.Responsibilities = responsibilities

			case "Технологии":
				technologies := strutil.SplitAnyClean(row[1], ",")
				wh.Technologies = technologies

			default:
				slog.Warn("Unexpected row header in projects section", slog.String("cell", row[0]))
			}
		}

		r = r.AddWorkHistory(wh)
	}

	return r, nil
}
