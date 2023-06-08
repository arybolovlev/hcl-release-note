package api

import "testing"

func TestResourceValidator(t *testing.T) {
	successCases := map[string]Resource{
		"NewResourceWithDoc": {
			NewResource: true,
			Doc:         "this",
		},
	}

	for n, c := range successCases {
		t.Run(n, func(t *testing.T) {
			if err := c.Validate(); err != nil {
				t.Errorf("Unexpected validation error: %v", err)
			}
		})
	}

	errorCases := map[string]Resource{
		"NewResourceWithoutDoc": {
			NewResource: true,
		},
	}

	for n, c := range errorCases {
		t.Run(n, func(t *testing.T) {
			if err := c.Validate(); err == nil {
				t.Error("Error is expected")
			}
		})
	}
}
