package api

import (
	"bar/autogen"
	"bar/internal/models"

	"github.com/labstack/echo/v4"
)

func (s *Server) SetCookie(c echo.Context, account *models.Account) {
	sess := s.getUserSess(c)
	sess.Options.MaxAge = 60 * 60 * 24 * 7 // 1 week
	sess.Options.HttpOnly = true
	sess.Options.Secure = true
	sess.Values["account_id"] = account.Account.Id.String()
	sess.Save(c.Request(), c.Response())

	if account.IsAdmin() {
		sess := s.getAdminSess(c)
		sess.Options.MaxAge = 60 * 60 * 24 * 7 // 1 week
		sess.Options.HttpOnly = true
		sess.Options.Secure = true
		sess.Values["admin_account_id"] = account.Account.Id.String()
		if account.Role == autogen.AccountSuperAdmin {
			sess.Values["super_admin"] = true
		}
		sess.Save(c.Request(), c.Response())
	}
}

func (s *Server) RemoveCookie(c echo.Context) {
	sess := s.getUserSess(c)
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	sess = s.getAdminSess(c)
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())
}