package handler

import (
	"fmt"
	"net/http"

	"beyond/application/article/api/internal/logic"
	"beyond/application/article/api/internal/svc"
	"beyond/application/article/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ArticleDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ArticleDetailRequest
		if err := httpx.Parse(r, &req); err != nil {
			fmt.Printf("parse request error: %v\n", err)
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewArticleDetailLogic(r.Context(), svcCtx)
		resp, err := l.ArticleDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
