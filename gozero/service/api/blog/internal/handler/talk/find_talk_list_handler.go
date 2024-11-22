package talk

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/ve-weiyi/ve-blog-golang/gozero/internal/responsex"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/blog/internal/logic/talk"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/blog/internal/types"
)

// 分页获取说说列表
func FindTalkListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TalkQueryReq
		if err := httpx.Parse(r, &req); err != nil {
			responsex.Response(r, w, nil, err)
			return
		}

		l := talk.NewFindTalkListLogic(r.Context(), svcCtx)
		resp, err := l.FindTalkList(&req)
		responsex.Response(r, w, resp, err)
	}
}