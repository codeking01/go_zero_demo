package handler

import (
	"go_zero_demo/apps/common/response"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_zero_demo/apps/login/internal/logic"
	"go_zero_demo/apps/login/internal/svc"
	"go_zero_demo/apps/login/internal/types"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		response.Response(r, w, resp, err)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
