package action

type ActionsFile struct {
	Actions []Action `yaml:"actions"`
}

type Action struct {
	Key         string            `yaml:"key"`
	Label       string            `yaml:"label"`
	Description string            `yaml:"description"`
	Scripts     map[string]string `yaml:"scripts"`
}
