package object

type MetaData struct {
	ProjectID   string            `json:"projectId" yaml:"projectId" validate:"required"`
	Name        string            `json:"name" yaml:"name" validate:"required"`
	UUID        string            `json:"uuid" yaml:"uuid" validate:"-"`
	DisplayName string            `json:"displayName" yaml:"displayName" validate:"-"`
	Description string            `json:"description" yaml:"description" validate:"-"`
	Annotations map[string]string `json:"annotations" yaml:"annotations" validate:"-"`
	Labels      map[string]string `json:"labels" yaml:"labels" validate:"-"`
	Enabled     bool              `json:"enabled" yaml:"enabled" validate:"-"`
}
