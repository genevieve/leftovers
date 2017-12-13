package fakes

type BucketManager struct {
	IsInRegionCall struct {
		CallCount int
		Receives  struct {
			Bucket string
			Region string
		}
		Returns struct {
			Output bool
		}
	}
}

func (b *BucketManager) IsInRegion(bucket, region string) bool {
	b.IsInRegionCall.CallCount++
	b.IsInRegionCall.Receives.Bucket = bucket
	b.IsInRegionCall.Receives.Region = region

	return b.IsInRegionCall.Returns.Output
}
