package sms

type SmsServer struct {
	AccessKeyId     *string
	AccessKeySecret *string
}

func (s SmsServer) SetAccessKeyId(key string) {
	s.AccessKeyId = &key
}

func (s SmsServer) SetAccessKeySecret(secret string) {
	s.AccessKeySecret = &secret
}

type SmsInterface interface {
	SendBatchSms(args interface{}) (_err error)
}
