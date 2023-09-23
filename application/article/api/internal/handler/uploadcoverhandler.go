package handler

import (
	"net/http"

	"beyond/application/article/api/internal/logic"
	"beyond/application/article/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadCoverHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUploadCoverLogic(r.Context(), svcCtx)
		resp, err := l.UploadCover(r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
