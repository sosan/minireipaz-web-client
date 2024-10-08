package routes

import (
	"minireipaz/pkg/common"
	"minireipaz/pkg/interfaces/controllers"
	"minireipaz/pkg/interfaces/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(app *gin.Engine, workflowController *controllers.WorkflowController, userController *controllers.UserController, dashboardController *controllers.DashboardController) {
	app.NoRoute(ErrRouter)
	route := app.Group("/api")
	{
		route.GET("/ping", common.Ping)
		route.POST("/workflows", middlewares.ValidateWorkflow(), workflowController.CreateWorkflow)
		// route.GET("/workflows", controllers.GetAllWorkflows)
		// route.GET("/workflows/:id", controllers.GetWorkflowByID)
		// route.PUT("/workflows/:id", controllers.UpdateWorkflow)
		// route.DELETE("/workflows/:id", controllers.DeleteWorkflow)
		route.POST("/users", middlewares.ValidateUser(), userController.SyncUser)
		route.GET("/dashboard/:iduser", middlewares.ValidateSub(), dashboardController.GetUserDashboardByID)
	}
}

func ErrRouter(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"errors": "this page could not be found",
	})
}
