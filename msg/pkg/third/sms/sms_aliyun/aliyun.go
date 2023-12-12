package sms_aliyun

import (
	"blog.hideyoshi.top/msg/pkg/third/sms"
	"encoding/json"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"strings"
)

type AliYun struct {
	sms.SmsServer
}

func (a AliYun) createClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Endpoint:        tea.String("dysmsapi.aliyuncs.com"),
	}
	return dysmsapi20170525.NewClient(config)
}

type AliYunBatchArgs struct {
	PhoneNumberJson   *string
	SignNameJson      *string
	TemplateCode      *string
	TemplateParamJson *string
	OutId             *string
}

func (a AliYun) SendBatchSms(args *AliYunBatchArgs) (_err error) {
	client, _err := a.createClient(a.AccessKeyId, a.AccessKeySecret)
	if _err != nil {
		return _err
	}
	sendBatchSmsRequest := &dysmsapi20170525.SendBatchSmsRequest{
		PhoneNumberJson:   args.PhoneNumberJson,
		SignNameJson:      args.SignNameJson,
		TemplateCode:      args.TemplateCode,
		TemplateParamJson: args.TemplateParamJson,
		OutId:             args.OutId,
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		_, _err = client.SendBatchSmsWithOptions(sendBatchSmsRequest, runtime)
		if _err != nil {
			return _err
		}

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 错误 message
		fmt.Println(tea.StringValue(error.Message))
		// 诊断地址
		var data interface{}
		d := json.NewDecoder(strings.NewReader(tea.StringValue(error.Data)))
		d.Decode(&data)
		if m, ok := data.(map[string]interface{}); ok {
			recommend, _ := m["Recommend"]
			fmt.Println(recommend)
		}
		_, _err = util.AssertAsString(error.Message)
		if _err != nil {
			return _err
		}
	}
	return _err
}
