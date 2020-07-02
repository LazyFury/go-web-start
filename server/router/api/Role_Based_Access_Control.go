package api

import "EK-Server/util/middleware"

var (
	rbacAdmin  = middleware.AdminJWT
	rbacUser   = middleware.UserJWT
	rbacAuthor = middleware.UserJWT
)
