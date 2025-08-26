package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T)    {
	cases := []struct {
		name      string
		header    http.Header
		output    string
		expectErr bool}{
		{
			name:      "Valid api key",
			header:    http.Header{"Authorization": []string{"ApiKey randomApiKey"}},
			output:    "randomApiKey",
			expectErr: false,
		},
		{
			name:      "invalid prefix",
			header:    http.Header{"Authorization": []string{"auth lsajdf9f7sd;k"}},
			output:    "lsajdf9f7sd;k",
			expectErr: true,
		},
		{
			name:      "valid api key 2",
			header:    http.Header{"Authorization": []string{"ApiKey als;dkjfjsalkdfjff97vc"}},
			output:    "als;dkjfjsalkdfjff97vc",
			expectErr: false,
		},
		{
			name:      "missing authorization header",
			header:    http.Header{"NotAuthorizationHeader": []string{"ApiKey sdf7890qwehjklo234hjkl"}},
			output:    "sdf7890qwehjklo234hjkl",
			expectErr: true,
		},
		{
			name:      "missing prefix",
			header:    http.Header{"Authorization": []string{"notApiKey"}},
			output:    "idk",
			expectErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T)                {
			got, err := GetAPIKey(c.header)
			
			
			
			
			
			
			if err != nil {
				if !c.expectErr {
					t.Errorf("unexpected error: %v\n", err)
				}
				return
			}

			if c.expectErr {
				t.Error("expected error but got none case")
				
				
				return
			}

			if got != c.output {
				t.Errorf("expect: %vgot: %v\n", c.output, got)
			}
		})
	}
}
