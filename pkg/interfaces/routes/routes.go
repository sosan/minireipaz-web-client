package routes

import (
	"minireipaz/pkg/common"
	"minireipaz/pkg/interfaces/controllers"
	"minireipaz/pkg/interfaces/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(app *gin.Engine, workflowController *controllers.WorkflowController, userController *controllers.UserController, dashboardController *controllers.DashboardController, credentialController *controllers.CredentialsController) {
	app.NoRoute(ErrRouter)
	route := app.Group("/api")
	{
		route.GET("/ping", common.Ping)
		route.POST("/workflows", middlewares.ValidateWorkflow(), workflowController.CreateWorkflow)
		route.GET("/workflows/:iduser", middlewares.ValidateUserID(), workflowController.GetAllWorkflows)
		route.GET("/workflows/:iduser/:idworkflow", middlewares.ValidateGetWorkflow(), workflowController.GetWorkflowByID)
		route.PUT("/workflows/:id", middlewares.ValidateUpdateWorkflow(), workflowController.UpdateWorkflow)
		// route.DELETE("/workflows/:id", controllers.DeleteWorkflow)
		route.POST("/users", middlewares.ValidateUser(), userController.SyncUser)
		route.GET("/dashboard/:iduser", middlewares.ValidateUserID(), dashboardController.GetUserDashboardByID)

		route.POST("/credentials", middlewares.ValidateCredential(), credentialController.CreateCredentials)
		route.GET("/credentials/:iduser", middlewares.ValidateUserID(), credentialController.GetAllCredentials)
		// route.GET("/credentials/:iduser/:idcredential", middlewares.ValidateGetCredential(), credentialController.GetCredentialsByID)
	}
	creds := app.Group("/oauth2-credentials")
	{
		creds.POST("/save", middlewares.ValidateCredentialExchange(), credentialController.CallbackCredentials)
	}
}

func ErrRouter(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"errors": "this page could not be found",
	})
}
