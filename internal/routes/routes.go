package routes

/*
import (
	"github.com/dundudun/rest_test_back/internal/handlers"
	"github.com/gin-gonic/gin"
)

func ApiGroup(r *gin.Engine) *gin.RouterGroup {
	return r.Group("/api")
}

func OrganizationGroup(r *gin.RouterGroup) *gin.RouterGroup {
	return r.Group("/organizations")
}

func WasteStorageGroup(r *gin.RouterGroup) *gin.RouterGroup {
	return r.Group("/waste_storages")
}

func CreateOrganization(r *gin.RouterGroup, handler *handlers.Handler) *gin.RouterGroup {
	r.POST("", handler.CreateOrganization)
	return r
}

/*
router := server.Group("/api")

organization := router.Group("/organizations")
{
	organization.POST("", handler.CreateOrganization)
	organization.GET("", handler.ListOrganizations)
	organization.GET("/:id", handler.GetOrganization)
	organization.PUT("/:id", handler.ChangeOrganization)
	organization.PATCH("/:id", handler.PartlyChangeOrganization)
	organization.DELETE("/:id", handler.DeleteOrganization)
	organization.POST("/:id/produce", handler.ProduceWaste)
}

storage := router.Group("/waste_storages")
{
	storage.POST("", handler.CreateWasteStorage)
	storage.GET("", handler.ListWasteStorages)
	storage.GET("/:id", handler.GetWasteStorage)
	storage.PUT("/:id", handler.ChangeWasteStorage)
	storage.PATCH("/:id", handler.PartlyChangeWasteStorage)
	storage.DELETE("/:id", handler.DeleteWasteStorage)
}
*/
