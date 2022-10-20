package Model

type UploadFileUtilityInputDTO struct {
	FileBytes    []byte
	FileName     string
	FileLength   int64
	RelationType int
	RelationId   int
}
