package Model

type UploadFileUtilityInputDTO struct {
	FileBytes    []byte
	FileName     string
	FileLength   int64
	Extension    string
	RelationType int
	RelationId   int
}
