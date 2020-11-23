//Routes/Routes.go
package Routes

import (
	"gorm-test/Controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/user-api")
	{
		grp1.GET("user", Controllers.GetUsers)
		grp1.POST("user", Controllers.CreateUser)
		grp1.GET("user/:id", Controllers.GetUserByID)
		grp1.PUT("user/:id", Controllers.UpdateUser)
		grp1.DELETE("user/:id", Controllers.DeleteUser)
	}
	grp2 := r.Group("/deal-api")
	{
		grp2.GET("deal", Controllers.GetDeals)
		grp2.POST("deal", Controllers.CreateDeal)
		grp2.GET("deal/:id", Controllers.GetDealByID)
		grp2.PUT("deal/:id", Controllers.UpdateDeal)
		grp2.DELETE("deal/:id", Controllers.DeleteDeal)
	}
	return r
}
