package aliyun

import (
	"net/http"

	"common/oss/api/internal/logic/aliyun"
	"common/oss/api/internal/svc"
	"common/oss/api/internal/types"

	"github.com/ginvcom/util"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func StsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := aliyun.NewStsLogic(r.Context(), svcCtx)
		resp, err := l.Sts(&req)
		util.Response(w, resp, err)
	}
}
