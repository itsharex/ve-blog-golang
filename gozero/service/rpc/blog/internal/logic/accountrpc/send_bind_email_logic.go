package accountrpclogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/gozero/global/constant"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/common/rediskey"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/biz/apierr"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/mail"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/tempx"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/valid"
)

type SendBindEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendBindEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendBindEmailLogic {
	return &SendBindEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发送绑定邮箱邮件
func (l *SendBindEmailLogic) SendBindEmail(in *accountrpc.UserEmailReq) (*accountrpc.EmptyResp, error) {
	// 校验邮箱格式
	if !valid.IsEmailValid(in.Username) {
		return nil, apierr.NewApiError(apierr.CodeInvalidParam, "邮箱格式不正确")
	}

	// 验证用户是否存在
	exist, err := l.svcCtx.TUserModel.FindOneByUsername(l.ctx, in.Username)
	if exist != nil {
		return nil, apierr.NewApiError(apierr.CodeUserAlreadyExist, "用户已存在")
	}

	// 发送验证码邮件
	key := rediskey.GetCaptchaKey(constant.BindEmail, in.Username)
	code, _ := l.svcCtx.CaptchaHolder.GetCodeCaptcha(key)
	data := mail.CaptchaEmail{
		Username: in.Username,
		Code:     code,
	}

	// 组装邮件内容
	content, err := tempx.TempParseString(mail.TempBind, data)
	if err != nil {
		return nil, err
	}

	msg := &mail.EmailMessage{
		To:      []string{in.Username},
		Subject: "绑定邮件提醒",
		Content: content,
		CC:      false,
	}
	// 发送邮件
	err = l.svcCtx.EmailDeliver.DeliveryEmail(msg)
	if err != nil {
		return nil, err
	}

	return &accountrpc.EmptyResp{}, nil
}
