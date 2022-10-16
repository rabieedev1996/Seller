package UploadFileCommand

type UploadFileCommand struct {
	RelationType int
	FileBytes    []byte
	FileName     string
	FileLength   int64
}
