package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop-api/goods-web/models"
)

// IsAdminAuth -> 用户鉴权
func IsAdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, _ := ctx.Get("claims")
		currentUser := claims.(*models.CustomClaims)
		// TODO 此处有bug：获取的Role->AuthorityId值不正确
		//fmt.Println(currentUser)
		if currentUser.AuthorityId == 0 {
			ctx.JSON(http.StatusForbidden, gin.H{
				"msg": "无权限",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
