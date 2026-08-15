package compose

import "testing"

func TestParseVolumeMount(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		source   string
		target   string
		readOnly bool
	}{
		{
			name:     "read write bind mount",
			input:    "./data:/app/data",
			source:   "./data",
			target:   "/app/data",
			readOnly: false,
		},
		{
			name:     "read only bind mount",
			input:    "./data:/app/data:ro",
			source:   "./data",
			target:   "/app/data",
			readOnly: true,
		},
		{
			name:     "docker socket read write",
			input:    "/var/run/docker.sock:/var/run/docker.sock",
			source:   "/var/run/docker.sock",
			target:   "/var/run/docker.sock",
			readOnly: false,
		},
		{
			name:     "docker socket read only",
			input:    "/var/run/docker.sock:/var/run/docker.sock:ro",
			source:   "/var/run/docker.sock",
			target:   "/var/run/docker.sock",
			readOnly: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mount := ParseVolumeMount(tt.input)

			if mount.Source != tt.source {
				t.Errorf("expected source %q, got %q", tt.source, mount.Source)
			}

			if mount.Target != tt.target {
				t.Errorf("expected target %q, got %q", tt.target, mount.Target)
			}

			if mount.ReadOnly != tt.readOnly {
				t.Errorf("expected readOnly %v, got %v", tt.readOnly, mount.ReadOnly)
			}
		})
	}
}
