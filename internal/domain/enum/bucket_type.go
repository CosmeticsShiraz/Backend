package enum

type BucketType uint

const (
	VATTaxpayerCertificate BucketType = iota + 1
	OfficialNewspaperAD
	ProfilePic
	LogoPic
	NewsMedia
)

func (bt BucketType) String() string {
	switch bt {
	case VATTaxpayerCertificate:
		return "vatTaxpayerCertificate"
	case OfficialNewspaperAD:
		return "officialNewspaperAD"
	case ProfilePic:
		return "profilePic"
	case LogoPic:
		return "logoPic"
	case NewsMedia:
		return "newsMedia"
	}
	return ""
}

func GetAllBucketTypes() []BucketType {
	return []BucketType{
		VATTaxpayerCertificate,
		OfficialNewspaperAD,
		ProfilePic,
		LogoPic,
		NewsMedia,
	}
}
