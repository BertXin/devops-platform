package service

import (
	"github.com/gin-gonic/gin"
	"devops-platform/internal/common/web/internal/domain"
	"strings"
)

var ignoreUrls []string

func AddIgnoreUrls(urls ...string) {
	if len(urls) == 0 {
		return
	}
	if len(ignoreUrls) == 0 {
		ignoreUrls = urls
		return
	}
	ignoreUrls = append(ignoreUrls, urls...)
}

func Verify(ctx *gin.Context) {

	requestURI := ctx.Request.RequestURI
	for _, url := range ignoreUrls {
		if strings.HasPrefix(requestURI, url) {
			ctx.Next()
			return
		}
	}

	if !Authenticated(ctx) {
		AbortErr(ctx, domain.NewUnauthorizedError("对不起，认证不通过，请登录"))
		return
	}

	ctx.Next()

}
