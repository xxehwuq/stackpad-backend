package transport

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (t *Transport) setUserId(ctx *gin.Context) {
	token := getTokenFromHeader(ctx)
	id, err := t.pkg.JwtManager.GetClaimFromToken(token, "sub")
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	ctx.Set("userId", id)
}

func (t *Transport) getUserId(ctx *gin.Context) string {
	id, _ := ctx.Get("userId")
	return fmt.Sprint(id)
}

func getTokenFromHeader(ctx *gin.Context) string {
	header := ctx.GetHeader("Authorization")
	headerParts := strings.Split(header, " ")
	token := headerParts[1]

	return token
}
