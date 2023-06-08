package api

import "fmt"

func (r *Resource) Validate() error {
	if r.NewResource {
		if r.Doc == "" {
			return fmt.Errorf("A new resource must have documentation")
		}
	}
	return nil
}
