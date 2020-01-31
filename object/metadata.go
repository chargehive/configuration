package object

type MetaData struct {
	ProjectID   string            `json:"projectId" yaml:"projectId" validate:"required"`
	Name        string            `json:"name" yaml:"name"`
	UUID        string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	DisplayName string            `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	Description string            `json:"description,omitempty" yaml:"description,omitempty"`
	Annotations map[string]string `json:"annotaions,omitempty" yaml:"annotaions,omitempty"`
	Labels      map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
}
