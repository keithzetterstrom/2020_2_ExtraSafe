package middlewares

import (
	"github.com/go-park-mail-ru/2020_2_ExtraSafe/internal/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"strconv"
)

type Middleware interface {
	CORS() echo.MiddlewareFunc
	CookieSession(next echo.HandlerFunc) echo.HandlerFunc
	AuthCookieSession(next echo.HandlerFunc) echo.HandlerFunc
	CheckBoardUserPermission(next echo.HandlerFunc) echo.HandlerFunc
	CheckCardUserPermission(next echo.HandlerFunc) echo.HandlerFunc
	CheckTaskUserPermission(next echo.HandlerFunc) echo.HandlerFunc
}

type middlew struct {
	sessionsService sessionsService
	errorWorker errorWorker
	authService authService
	authTransport authTransport
	boardStorage boardStorage
}

func NewMiddleware(sessionsService sessionsService, errorWorker errorWorker, authService authService,
	authTransport authTransport) Middleware {
	return middlew{
		sessionsService: sessionsService,
		errorWorker: errorWorker,
		authService: authService,
		authTransport: authTransport,
	}
}

func (m middlew) CORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://95.163.213.142:80",
			"http://95.163.213.142",
			"http://127.0.0.1:3033",
			"http://tabutask.ru"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	})
}

func (m middlew) CookieSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := m.sessionsService.CheckCookie(c)
		if err != nil {
			return m.errorWorker.TransportError(c)
		}
		c.Set("userId", userId)
		return next(c)
	}
}

func (m middlew) AuthCookieSession(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, err := m.sessionsService.CheckCookie(c)
		if err == nil {
			userInput := new(models.UserInput)
			userInput.ID = userId
			user, _ := m.authService.Auth(*userInput)
			response, err := m.authTransport.AuthWrite(user)
			if err != nil {
				return m.errorWorker.TransportError(c)
			}
			return c.JSON(http.StatusOK, response)
		}
		return next(c)
	}
}

func (m middlew) CheckBoardUserPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		bid := c.Param("boardID")
		boardID, err := strconv.ParseUint(bid, 10, 64)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		userID := c.Get("userId").(uint64)

		err = m.boardStorage.CheckBoardPermission(userID, boardID, false)
		if err != nil {
			return c.NoContent(http.StatusForbidden)
		}

		c.Set("boardID", bid)

		return next(c)
	}
}

func (m middlew) CheckBoardAdminPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		bid := c.Param("boardID")
		boardID, err := strconv.ParseUint(bid, 10, 64)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		userID := c.Get("userId").(uint64)

		err = m.boardStorage.CheckBoardPermission(userID, boardID, true)
		if err != nil {
			return c.NoContent(http.StatusForbidden)
		}

		c.Set("boardID", bid)

		return next(c)
	}
}

func (m middlew) CheckCardUserPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cid := c.Param("cardID")
		cardID, err := strconv.ParseUint(cid, 10, 64)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		userID := c.Get("userId").(uint64)

		err = m.boardStorage.CheckCardPermission(userID, cardID)
		if err != nil {
			return c.NoContent(http.StatusForbidden)
		}

		c.Set("boardID", cid)

		return next(c)
	}
}

func (m middlew) CheckTaskUserPermission(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tid := c.Param("taskID")
		taskID, err := strconv.ParseUint(tid, 10, 64)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		userID := c.Get("userId").(uint64)

		err = m.boardStorage.CheckCardPermission(userID, taskID)
		if err != nil {
			return c.NoContent(http.StatusForbidden)
		}

		c.Set("boardID", tid)

		return next(c)
	}
}