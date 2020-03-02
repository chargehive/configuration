package object

type MetaData struct {
	ProjectID   string            `json:"projectId" yaml:"projectId" validate:"required"`
	Name        string            `json:"name" yaml:"name"`
	UUID        string            `json:"uuid" yaml:"uuid"`
	DisplayName string            `json:"displayName" yaml:"displayName"`
	Description string            `json:"description" yaml:"description"`
	Annotations map[string]string `json:"annotations" yaml:"annotations"`
	Labels      map[string]string `json:"labels" yaml:"labels"`
}
