package enum

type BucketType uint

const (
	ProfilePic = iota + 1
	NewsMedia
)

func (bt BucketType) String() string {
	switch bt {
	case ProfilePic:
		return "profilePic"
	case NewsMedia:
		return "newsMedia"
	}
	return ""
}

func GetAllBucketTypes() []BucketType {
	return []BucketType{
		ProfilePic,
		NewsMedia,
	}
}
