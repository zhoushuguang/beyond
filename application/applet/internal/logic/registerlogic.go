package logic

import (
	"context"
	"errors"
	"strings"

	"beyond/application/applet/internal/svc"
	"beyond/application/applet/internal/types"
	"beyond/application/user/rpc/user"
	"beyond/pkg/encrypt"
	"beyond/pkg/jwt"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	prefixActivation = "biz#activation#"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (*types.RegisterResponse, error) {
	req.Name = strings.TrimSpace(req.Name)
	req.Mobile = strings.TrimSpace(req.Mobile)
	req.Password = strings.TrimSpace(req.Password)
	if len(req.Password) == 0 {
		req.Password = encrypt.EncPassword(req.Password)
	}
	req.VerificationCode = strings.TrimSpace(req.VerificationCode)
	if len(req.VerificationCode) == 0 {
		return nil, errors.New("验证码不能为空")
	}
	err := l.checkVerificationCode(l.ctx, req.Mobile, req.VerificationCode)
	if err != nil {
		return nil, err
	}

	userRet, err := l.svcCtx.UserRPC.FindByMobile(l.ctx, &user.FindByMobileRequest{
		Mobile: req.Mobile,
	})
	if err != nil {
		return nil, err
	}
	if userRet != nil && userRet.UserId > 0 {
		return nil, errors.New("该手机号已注册")
	}
	regRet, err := l.svcCtx.UserRPC.Register(l.ctx, &user.RegisterRequest{
		Username: req.Name,
		Mobile:   req.Mobile,
	})
	if err != nil {
		return nil, err
	}

	token, err := jwt.BuildTokens(jwt.TokenOptions{
		AccessSecret:  l.svcCtx.Config.Auth.AccessSecret,
		AccessExpire:  l.svcCtx.Config.Auth.AccessExpire,
		RefreshSecret: l.svcCtx.Config.Auth.RefreshSecret,
		RefreshExpire: l.svcCtx.Config.Auth.RefreshExpire,
		RefreshAfter:  l.svcCtx.Config.Auth.RefreshAfter,
		Fields: map[string]interface{}{
			"userId": regRet.UserId,
		},
	})
	if err != nil {
		return nil, err
	}

	return &types.RegisterResponse{
		UserId: regRet.UserId,
		Token:  token,
	}, nil
}

func (l *RegisterLogic) checkVerificationCode(ctx context.Context, mobile, code string) error {
	cacheCode, err := getActivationCache(mobile, l.svcCtx.BizRedis)
	if err != nil {
		return err
	}
	if cacheCode == "" {
		return errors.New("验证码已过期")
	}
	if cacheCode != code {
		return errors.New("验证码错误")
	}

	return nil
}
