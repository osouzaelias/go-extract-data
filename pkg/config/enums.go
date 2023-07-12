package config

type BucketTypeName string

const (
	BucketTypeSor  BucketTypeName = "SOR"
	BucketTypeSot  BucketTypeName = "SOT"
	BucketTypeSpec BucketTypeName = "SPEC"
)

func (BucketTypeName) BucketTypeNames() []BucketTypeName {
	return []BucketTypeName{"SOR", "SOT", "SPEC"}
}
