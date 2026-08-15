package compose

import "strings"

type VolumeMount struct {
	Source   string
	Target   string
	ReadOnly bool
}

func ParseVolumeMount(volume string) VolumeMount {
	parts := strings.Split(volume, ":")

	mount := VolumeMount{}

	if len(parts) >= 1 {
		mount.Source = parts[0]
	}

	if len(parts) >= 2 {
		mount.Target = parts[1]
	}

	if len(parts) >= 3 {
		mount.ReadOnly = parts[len(parts)-1] == "ro"
	}

	return mount
}
