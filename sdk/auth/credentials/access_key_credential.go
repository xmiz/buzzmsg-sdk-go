package credentials

type AccessKeyCredential struct {
	AccessKeyId     string
	AccessKeySecret string
}

func NewAccessKeyCredential(accessKeyId, accessKeySecret string) *AccessKeyCredential {
	return &AccessKeyCredential{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}
}

func (baseCred *AccessKeyCredential) ToAccessKeyCredential() *AccessKeyCredential {
	return &AccessKeyCredential{
		AccessKeyId:     baseCred.AccessKeyId,
		AccessKeySecret: baseCred.AccessKeySecret,
	}
}
