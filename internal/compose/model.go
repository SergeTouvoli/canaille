package compose

type ComposeFile struct {
	Services map[string]Service `yaml:"services"`
}

type Service struct {
	Image         string      `yaml:"image"`
	ContainerName string      `yaml:"container_name"`
	Build         Build       `yaml:"build"`
	Ports         []string    `yaml:"ports"`
	Restart       string      `yaml:"restart"`
	Command       string      `yaml:"command"`
	Networks      []string    `yaml:"networks"`
	Healthcheck   Healthcheck `yaml:"healthcheck"`
	Privileged    bool        `yaml:"privileged"`
	Volumes       []string    `yaml:"volumes"`
}

type Build struct {
	Context    string `yaml:"context"`
	Dockerfile string `yaml:"dockerfile"`
}

type Healthcheck struct {
	Test        []string `yaml:"test"`
	Interval    string   `yaml:"interval"`
	Timeout     string   `yaml:"timeout"`
	Retries     int      `yaml:"retries"`
	StartPeriod string   `yaml:"start_period"`
}
