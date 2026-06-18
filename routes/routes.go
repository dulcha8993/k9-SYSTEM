package routes

import (
	"github.com/gin-gonic/gin"

	"k9-system/middleware"
)

import (
	authHandler "k9-system/handlers/auth"
	userHandler "k9-system/handlers/user"
	k9ProfileHandler "k9-system/handlers/k9_profile"
	k9TrainerHandler "k9-system/handlers/k9_trainer"
	k9TrainingRecordHandler "k9-system/handlers/training_record"
	k9HealthRecordHandler "k9-system/handlers/k9_health"
	criminalCaseHandler "k9-system/handlers/officer_intelligence"
	dashboardHandler "k9-system/handlers/dashboard"
)

func SetupRoutes(router *gin.Engine) {

	api := router.Group("/api")

	// PUBLIC ROUTES
	{
		api.POST("/auth/login", authHandler.Login)
	}

	// PROTECTED ROUTES

	// user APIs
	{
		api.POST( "/users",

			middleware.AuthMiddleware(),

			middleware.PermissionMiddleware(
				"system_admin",
				"create",
			),

		userHandler.CreateUser,
)
	}

	{
		api.PUT( "/users/:id",

			middleware.AuthMiddleware(),

			middleware.PermissionMiddleware(
				"system_admin",
				"update",
			),

		userHandler.UpdateUser,
)
	}

	{
		api.GET(
		"/users",

		middleware.AuthMiddleware(),

		middleware.PermissionMiddleware(
			"system_admin",
			"view",
		),

		userHandler.GetUsers,
	)
}

{	api.GET(
	"/users/:id",

	middleware.AuthMiddleware(),

	middleware.PermissionMiddleware(
		"system_admin",
		"view",
	),

	userHandler.GetUserByID,
)}


// user roles

	api.GET(
		"/user-roles",

		middleware.AuthMiddleware(),

		middleware.PermissionMiddleware(
			"system_admin",
			"view",
		),

		userHandler.GetUserRoles,
	)

// k9 profile APIs
	{
		api.POST(
			"/k9-profiles",

			middleware.AuthMiddleware(),

			middleware.PermissionMiddleware(
				"K9_profile_management",
				"create",
			),

			k9ProfileHandler.CreateK9Profile,
		)
	}

	{
		api.PUT(
			"/k9-profiles/:id",

			middleware.AuthMiddleware(),

			middleware.PermissionMiddleware(
				"K9_profile_management",
				"update",
			),

			k9ProfileHandler.UpdateK9Profile,
		)
	}

	{
		api.GET(
			"/k9-profiles",

			middleware.AuthMiddleware(),

			middleware.PermissionMiddleware(
				"K9_profile_management",
				"view",
			),

			k9ProfileHandler.GetK9Profiles,
		)
	}

	{
		api.GET(
			"/k9-profiles/:id",

			middleware.AuthMiddleware(),

			middleware.PermissionMiddleware(
				"K9_profile_management",
				"view",
			),

			k9ProfileHandler.GetK9ProfileByID,
		)
	}

// k9 trainer APIs

	{
		api.POST(
			"/k9-trainers",

			middleware.AuthMiddleware(),

			middleware.PermissionMiddleware(
				"k9_handler",
				"create",
			),

			k9TrainerHandler.CreateK9Trainer,
		)
	}

	{
		api.PUT(
			"/k9-trainers/:id",

			middleware.AuthMiddleware(),

			middleware.PermissionMiddleware(
				"k9_handler",
				"update",
			),

			k9TrainerHandler.UpdateK9Trainer,
		)
	}

	{
		api.GET(
			"/k9-trainers",

			middleware.AuthMiddleware(),

			middleware.PermissionMiddleware(
				"k9_handler",
				"view",
			),

			k9TrainerHandler.GetK9Trainers,
		)
	}

	
	{
		api.POST(
			"/k9-trainers/assign",

			middleware.AuthMiddleware(),

			middleware.PermissionMiddleware(
				"k9_handler",
				"create",
			),

			k9TrainerHandler.AssignK9TrainerToK9,
		)
	}

// K9 Training Records

	{
		api.POST(
			"/k9-training-record",

			middleware.AuthMiddleware(),

			middleware.PermissionMiddleware(
				"k9_training",
				"create",
			),

			k9TrainingRecordHandler.CreateTraningRecord,
		)
	}

	{
		api.PUT(
			"/k9-training-record/:id",

			middleware.AuthMiddleware(),

			middleware.PermissionMiddleware(
				"k9_training",
				"update",
			),

			k9TrainingRecordHandler.UpdateTrainingRecord,
		)
	}

	{
		api.GET(
			"/k9-training-record",

			middleware.AuthMiddleware(),

			middleware.PermissionMiddleware(
				"k9_training",
				"view",
			),

			k9TrainingRecordHandler.GetTrainingRecords,
		)
	}

//K9 Health Records

  {
	api.POST(
		"/k9-health-record",

		middleware.AuthMiddleware(),

		middleware.PermissionMiddleware(
			"k9_veterianary",
			"create",
		),

		k9HealthRecordHandler.CretateHealthRecord,
	)
  }

  	{
		api.PUT(
			"/k9-health-record/:id",

			middleware.AuthMiddleware(),

			middleware.PermissionMiddleware(
				"k9_veterianary",
				"update",
			),

			k9HealthRecordHandler.UpdateHealthRecord,
		)
	}

	{
		api.GET(
			"/k9-health-record",

			middleware.AuthMiddleware(),

			middleware.PermissionMiddleware(
				"k9_veterianary",
				"view",
			),

			k9HealthRecordHandler.GetHealthRecords,
		)
	}

	{
		api.GET(
			"/k9-health-record/:id",

			middleware.AuthMiddleware(),

			middleware.PermissionMiddleware(
				"k9_veterianary",
				"view",
			),

			k9HealthRecordHandler.GetHealthRecordByID,
		)
	}

// criminal case APIs

 {
	api.POST(
		"/criminal-cases",

		middleware.AuthMiddleware(),

		middleware.PermissionMiddleware(
			"officer_intelligence",
			"create",
		),

		criminalCaseHandler.CreateCriminalCase,
	)
 }

{
	api.PUT(
		"/criminal-cases/:id",

		middleware.AuthMiddleware(),

		middleware.PermissionMiddleware(
			"officer_intelligence",
			"update",
		),

		criminalCaseHandler.UpdateCriminalCase,
	)
}
{
		api.GET(
		"/criminal-cases",

		middleware.AuthMiddleware(),

		middleware.PermissionMiddleware(
			"officer_intelligence",
			"view",
		),

		criminalCaseHandler.GetCriminalCases,
	)
}

{	
	api.GET(
		"/criminal-cases/:id",

		middleware.AuthMiddleware(),

		middleware.PermissionMiddleware(
			"officer_intelligence",
			"view",
		),

		criminalCaseHandler.GetCriminalCaseByID,
	)
}

// dashbaord api

{
	api.GET(
		"/dashboard",

		middleware.AuthMiddleware(),

		// middleware.PermissionMiddleware(
		// 	"dashboard",
		// 	"view",
		// ),

		dashboardHandler.GetDashboard,
	)
}
 }
