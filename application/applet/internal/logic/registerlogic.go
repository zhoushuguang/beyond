package logic

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"beyond/application/applet/internal/code"
	"beyond/application/applet/internal/svc"
	"beyond/application/applet/internal/types"
	"beyond/application/user/rpc/user"
	"beyond/pkg/encrypt"
	"beyond/pkg/jwt"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	prefixActivation = "biz#activation#%s"
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
	if len(req.Mobile) == 0 {
		return nil, code.RegisterMobileEmpty
	}
	req.Password = strings.TrimSpace(req.Password)
	if len(req.Password) == 0 {
		req.Password = encrypt.EncPassword(req.Password)
	}
	req.VerificationCode = strings.TrimSpace(req.VerificationCode)
	if len(req.VerificationCode) == 0 {
		return nil, code.VerificationCodeEmpty
	}
	err := l.checkVerificationCode(l.ctx, req.Mobile, req.VerificationCode)
	if err != nil {
		return nil, err
	}
	mobile, err := encrypt.EncMobile(req.Mobile)
	if err != nil {
		logx.Errorf("EncMobile mobile: %s error: %v", req.Mobile, err)
		return nil, err
	}
	userRet, err := l.svcCtx.UserRPC.FindByMobile(l.ctx, &user.FindByMobileRequest{
		Mobile: mobile,
	})
	if err != nil {
		return nil, err
	}
	if userRet != nil && userRet.UserId > 0 {
		return nil, code.MobileHasRegistered
	}
	regRet, err := l.svcCtx.UserRPC.Register(l.ctx, &user.RegisterRequest{
		Username: req.Name,
		Mobile:   mobile,
	})
	if err != nil {
		fmt.Printf("register err %v\n", err)
		return nil, err
	}

	token, err := jwt.BuildTokens(jwt.TokenOptions{
		AccessSecret: l.svcCtx.Config.Auth.AccessSecret,
		AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
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
		return errors.New("verification code expired")
	}
	if cacheCode != code {
		return errors.New("verification code failed")
	}

	return nil
}
