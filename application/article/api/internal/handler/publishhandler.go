package handler

import (
	"net/http"

	"beyond/application/article/api/internal/logic"
	"beyond/application/article/api/internal/svc"
	"beyond/application/article/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PublishHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewPublishLogic(r.Context(), svcCtx)
		resp, err := l.Publish(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
