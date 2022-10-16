package Model

type UploadFileServiceInputDTO struct {
	FileBytes  []byte
	FileName   string
	FileLength int64
	Extension  string
	Path       string
}
type UploadFileServiceOutputDTO struct {
	BaseUrl     string
	RelativeUrl string
}
