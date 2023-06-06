package api

type Resource struct {
	Name string `hcl:"name,label"`

	BugFix      string `hcl:"bug_fix"`
	Doc         string `hcl:"doc"`
	Enhancement string `hcl:"enhancement"`
	Feature     string `hcl:"feature"`
	NewResource bool   `hcl:"new_resource"`
}
