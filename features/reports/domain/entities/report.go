package entities

import "fmt"

// Report generated and structured toggl report
type Report struct {
	Tasks []*Task
	ByProject map[string]*Project
	ByTag map[string][]*Task
}

func (r Report) String() string {
	str := "\n{\n\tTasks: ["

	for _, t := range r.Tasks {
		str += "\n\t\t{" +
			"\n\t\t\tName: " + t.Name +
			"\n\t\t\tStart: " + t.Start.String() +
			"\n\t\t\tEnd: " + t.End.String() + 
			"\n\t\t\tDuration: " + fmt.Sprint(t.Duration) +
			"\n\t\t\tTags: ["

		for i, tags := range t.Tags {
			str += tags

			if i < len(t.Tags) - 1 {
				str += ", "
			}
		}

		str += "]\n\t\t},"
	}
	str += "\n\t],\n\tByProject: {"
	for key, p := range r.ByProject {
		str += fmt.Sprintf("\n\t\t%v: {", key) +
			"\n\t\t\tName: " + p.Name +
			"\n\t\t\tColor: " + p.Color +
			"\n\t\t\tTasks: ["

			for i, t := range p.Tasks {
				str += t.Name

				if i < len(t.Tags) - 1 {
				str += ", "
			}
			}
			str += "]\n\t\t},"
	}
	str += "\n\t},\n\tByTag: {"
	for key, tag := range r.ByTag {
		str += fmt.Sprintf("\n\t\t%v: [", key)

		for i, t := range tag {
			str += t.Name

			if i < len(t.Tags) - 1 {
				str += ", "
			}
		}
		str += "]"
	}

	str += "\n\t},\n}\n"
	return str
}