package routes

import (
	"echo_blogs/controllers/admin"
	"echo_blogs/controllers/admin/auth"
	"echo_blogs/middleware"
	"echo_blogs/validators"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func Admin(e *echo.Echo, db *gorm.DB) {
	authController := auth.NewAuthController{DB: db}
	dashboardController := admin.NewDashboardController{DB: db}
	categoryController := admin.NewCategoryController{DB: db}
	userController := admin.NewUserController{DB: db}
	blogRepository := admin.NewBlogController{DB: db}

	authMiddleware := middleware.NewCheckAuthMiddleware{DB: db}
	guestMiddleware := middleware.NewGuestMiddleware{DB: db}
	sessionMiddleware := middleware.NewSessionMiddleware{DB: db}

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.Use(sessionMiddleware.Handle)

	e.Validator = &validators.NewCustomValidator{Validator: validator.New()}

	g := e.Group("/admin")

	g.GET("/login", authController.ShowLogin, guestMiddleware.Handle).Name = "admin.login"
	g.POST("/login/user", authController.Login, guestMiddleware.Handle).Name = "admin.post.login"
	g.GET("/logout", authController.Logout, authMiddleware.Handle).Name = "admin.logout"

	g.GET("/register", authController.ShowRegister, guestMiddleware.Handle).Name = "admin.register"
	g.POST("/register", authController.Register, guestMiddleware.Handle).Name = "admin.post.register"

	g.GET("/dashboard", dashboardController.Index, authMiddleware.Handle).Name = "admin.dashboard"

	g.GET("/categories", categoryController.Index, authMiddleware.Handle).Name = "admin.categories"
	g.GET("/categories/list", categoryController.List, authMiddleware.Handle).Name = "admin.categories.list"
	g.GET("/categories/create", categoryController.Create, authMiddleware.Handle).Name = "admin.categories.create"
	g.POST("/categories/save", categoryController.Save, authMiddleware.Handle).Name = "admin.categories.save"
	g.GET("/categories/:id/edit", categoryController.Edit, authMiddleware.Handle).Name = "admin.categories.edit"
	g.POST("/categories/:id/update", categoryController.Update, authMiddleware.Handle).Name = "admin.categories.update"
	g.DELETE("/categories/:id/delete", categoryController.Delete, authMiddleware.Handle).Name = "admin.categories.delete"

	g.GET("/users", userController.Index, authMiddleware.Handle).Name = "admin.users"
	g.GET("/users/list", userController.List, authMiddleware.Handle).Name = "admin.users.list"
	g.GET("/users/create", userController.Create, authMiddleware.Handle).Name = "admin.users.create"
	g.POST("/users/save", userController.Save, authMiddleware.Handle).Name = "admin.users.save"
	g.GET("/users/:id/edit", userController.Edit, authMiddleware.Handle).Name = "admin.users.edit"
	g.POST("/users/:id/update", userController.Update, authMiddleware.Handle).Name = "admin.users.update"
	g.DELETE("/users/:id/delete", userController.Delete, authMiddleware.Handle).Name = "admin.users.delete"

	g.GET("/blogs", blogRepository.Index, authMiddleware.Handle).Name = "admin.blogs"
	g.GET("/blogs/list", blogRepository.List, authMiddleware.Handle).Name = "admin.blogs.list"
	g.GET("/blogs/create", blogRepository.Create, authMiddleware.Handle).Name = "admin.blogs.create"
	g.POST("/blogs/save", blogRepository.Save, authMiddleware.Handle).Name = "admin.blogs.save"
	g.GET("/blogs/:id/edit", blogRepository.Edit, authMiddleware.Handle).Name = "admin.blogs.edit"
	g.POST("/blogs/:id/update", blogRepository.Update, authMiddleware.Handle).Name = "admin.blogs.update"
	g.DELETE("/blogs/:id/delete", blogRepository.Delete, authMiddleware.Handle).Name = "admin.blogs.delete"

	g.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusFound, c.Echo().Reverse("admin.dashboard"))
	})

}
