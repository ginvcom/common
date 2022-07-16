package aliyun

import (
	"context"
	"errors"

	"common/oss/api/internal/svc"
	"common/oss/api/internal/types"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"github.com/zeromicro/go-zero/core/logx"
)

type StsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StsLogic {
	return &StsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StsLogic) Sts(req *types.StsReq) (resp *types.StsReply, err error) {
	// todo: add your logic here and delete this line

	//构建一个阿里云客户端, 用于发起请求。
	//构建阿里云客户端时，需要设置AccessKey ID和AccessKey Secret。
	bucketKey := req.Bucket
	logx.Info(bucketKey)
	logx.Info(l.svcCtx.Config)
	ossConf, ok := l.svcCtx.Config.AliyunOSS[bucketKey]
	if !ok {
		err = errors.New("bucket" + bucketKey + "不存在")
		return
	}
	logx.Info(ossConf.Domain)
	client, err := sts.NewClientWithAccessKey(ossConf.Region, ossConf.AccessKey, ossConf.SecretKey)

	if err != nil {
		return
	}

	//构建请求对象。
	request := sts.CreateAssumeRoleRequest()
	request.Scheme = "https"

	//设置参数。关于参数含义和设置方法，请参见API参考。
	request.RoleArn = "acs:ram::xxxx:role/ginv"
	request.RoleSessionName = "ginv"

	//发起请求，并得到响应。
	response, err := client.AssumeRole(request)
	if err != nil {
		return
	}

	resp = &types.StsReply{
		Region:        "oss-" + ossConf.Region,
		Endpoint:      ossConf.Domain,
		AccessKey:     response.Credentials.AccessKeyId,
		Bucket:        ossConf.Bucket,
		SecretKey:     response.Credentials.AccessKeySecret,
		SecurityToken: response.Credentials.SecurityToken,
	}

	return
}
