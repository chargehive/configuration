package object

type MetaData struct {
	ProjectID   string            `json:"projectId" yaml:"projectId" validate:"-"`
	Name        string            `json:"name" yaml:"name" validate:"required,lowercase"`
	UUID        string            `json:"uuid,omitempty" yaml:"uuid,omitempty" validate:"-"`
	DisplayName string            `json:"displayName" yaml:"displayName" validate:"-"`
	Description string            `json:"description" yaml:"description" validate:"-"`
	Annotations map[string]string `json:"annotations" yaml:"annotations" validate:"-"`
	Labels      map[string]string `json:"labels" yaml:"labels" validate:"-"`
	Disabled    bool              `json:"disabled" yaml:"disabled" validate:"-"`
}
