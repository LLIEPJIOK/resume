package docs

type contentType string

const (
	TypeParagraph contentType = "paragraph"
	TypeTable     contentType = "table"
)

func (t contentType) String() string {
	return string(t)
}

func (t contentType) Paragraph() bool {
	return t == TypeParagraph
}

func (t contentType) Table() bool {
	return t == TypeTable
}

const expectedContentLen = 7

func expectedContentType(idx int) contentType {
	mp := map[int]contentType{
		0: TypeParagraph, // name
		1: TypeParagraph, // position
		2: TypeTable,     // skills + projects + education + language skills
		3: TypeParagraph, // work experience header
		4: TypeTable,     // work experience details
		5: TypeParagraph, // work history header
		6: TypeTable,     // work history details
	}

	ct, ok := mp[idx]
	if !ok {
		return TypeTable // default for any extra work history entries
	}

	return ct
}
