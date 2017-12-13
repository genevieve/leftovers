package fakes

type BucketManager struct {
	IsInRegionCall struct {
		CallCount int
		Receives  struct {
			Bucket string
		}
		Returns struct {
			Output bool
		}
	}
}

func (b *BucketManager) IsInRegion(bucket string) bool {
	b.IsInRegionCall.CallCount++
	b.IsInRegionCall.Receives.Bucket = bucket

	return b.IsInRegionCall.Returns.Output
}
