package profileHandler

import (
	"github.com/labstack/echo"
	"net/http"
)

type Handler interface {
	Profile(c echo.Context) error
	Accounts(c echo.Context) error
	ProfileChange(c echo.Context) error
	AccountsChange(c echo.Context) error
	PasswordChange(c echo.Context) error
}

type handler struct {
	profileService   profileService
	profileTransport profileTransport
	//errorWorker     errorWorker
}

func NewHandler(profileService profileService, profileTransport profileTransport) *handler {
	return &handler{
		profileService:   profileService,
		profileTransport: profileTransport,
		//errorWorker:     errorWorker,
	}
}

func (h *handler) Profile(c echo.Context) error {
	userInput, err := h.profileTransport.ProfileRead(c)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	user, err := h.profileService.Profile(userInput)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	response, err := h.profileTransport.ProfileWrite(user)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, response)
}

func (h *handler) Accounts(c echo.Context) error {
	userInput, err := h.profileTransport.ProfileRead(c)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	user, err := h.profileService.Accounts(userInput)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	response, err := h.profileTransport.AccountsWrite(user)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, response)
}

func (h *handler) ProfileChange(c echo.Context) error {
	userInput, err := h.profileTransport.ProfileChangeRead(c)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	user, err := h.profileService.ProfileChange(userInput)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	response, err := h.profileTransport.ProfileWrite(user)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, response)
}

func (h *handler) AccountsChange(c echo.Context) error {
	//cc := c.(*models.CustomContext)

	userInput, err := h.profileTransport.AccountsChangeRead(c)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	user, err := h.profileService.AccountsChange(userInput)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	response, err := h.profileTransport.AccountsWrite(user)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, response)
}

func (h *handler) PasswordChange(c echo.Context) error {
	userInput, err := h.profileTransport.PasswordChangeRead(c)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	user, err := h.profileService.PasswordChange(userInput)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	response, err := h.profileTransport.ProfileWrite(user)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, response)
}

/*func (h *handler) Photo(ctx *fasthttp.RequestCtx) {
	request, err := h.uploadTransport.PhotoDecode(ctx)
	if err != nil {
		err = h.errorWorker.ServeJSONError(ctx, err)
		if err != nil {
			h.errorWorker.ServeFatalError(ctx)
		}
		return
	}

	response, err := h.uploadService.Photo(request)
	if err != nil {
		err = h.errorWorker.ServeJSONError(ctx, err)
		if err != nil {
			h.errorWorker.ServeFatalError(ctx)
		}
		return
	}

	err = h.uploadTransport.PhotoEncode(response, ctx)
	if err != nil {
		err = h.errorWorker.ServeJSONError(ctx, err)
		if err != nil {
			h.errorWorker.ServeFatalError(ctx)
		}
		return
	}
}*/