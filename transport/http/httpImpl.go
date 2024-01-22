package transport

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Antoha2/sandbox/internal/service"
	"github.com/Antoha2/sandbox/pkg/logger/sl"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (a *apiImpl) StartHTTP() error {
	router := gin.Default()
	router.GET("/users/:id", a.getUserHandler)    //get user
	router.GET("/users/", a.getUsersHandler)      //get userS
	router.POST("/users/", a.addUserHandler)      //add user
	router.DELETE("/users/:id", a.delUserHandler) //del user
	router.PUT("/users/:id", a.updateUserHandler) //update user

	err := router.Run(fmt.Sprintf(":%s", a.cfg.HTTP.HostPort))
	if err != nil {
		a.log.Debug("ocurred error StartHTTP", sl.Err(err))
		return errors.Wrap(err, "ocurred error StartHTTP")
	}
	return nil
}

func (a *apiImpl) Stop() {
	if err := a.server.Shutdown(context.TODO()); err != nil {
		panic(errors.Wrap(err, "ocurred error Stop"))
	}
}

//get user
func (a *apiImpl) getUserHandler(c *gin.Context) {

	const op = "getUser"
	log := a.log.With(slog.String("op", op))

	id, err := strconv.Atoi(c.Param(ID))
	if err != nil {

		a.log.Debug("id not match type", sl.Err(err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	log.Info("run get User by ID", sl.Atr("id", id))

	user, err := a.service.GetUser(c, id)
	if err != nil {
		a.log.Debug("ocurred error Get User", sl.Err(err))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Info("get User successfully", sl.Atr("respUser", user))

	c.JSON(http.StatusOK, user)
}

//get users
func (a *apiImpl) getUsersHandler(c *gin.Context) {
	const op = "getUsers"
	log := a.log.With(slog.String("op", op))

	var err error
	age := service.DefaultPropertyAge
	limit := service.DefaultPropertyLimit
	offset := service.DefaultPropertyOffset

	q := c.Request.URL.Query()

	qAge := q.Get(AGE)
	if qAge != "" {
		age, err = strconv.Atoi(qAge)
		if err != nil {
			a.log.Debug("Age not match type", sl.Err(err))
			c.JSON(http.StatusBadRequest, sl.Err(err))
			return
		}
	}

	rOffset := q.Get(OFFSET)
	if rOffset != "" {
		offset, err = strconv.Atoi(rOffset)
		if err != nil {
			a.log.Debug("offset not match type", sl.Err(err))
			c.JSON(http.StatusBadRequest, sl.Err(err))
			return
		}
	}

	rLimit := q.Get(LIMIT)
	if rLimit != "" {
		limit, err = strconv.Atoi(rLimit)
		if err != nil {
			a.log.Debug("limit not match type", sl.Err(err))
			c.JSON(http.StatusBadRequest, sl.Err(err))
			return
		}
	}
	userQuery := &service.QueryUsersFilter{
		Name:        q.Get("name"),
		SurName:     q.Get("surname"),
		Patronymic:  q.Get("patronymic"),
		Age:         age,
		Gender:      q.Get("gender"),
		Nationality: q.Get("nationality"),
		Offset:      offset,
		Limit:       limit,
	}

	log.Info("run get Users", sl.Atr("filter", userQuery))

	users, err := a.service.GetUsers(c, userQuery)
	if err != nil {
		a.log.Debug("ocurred error Get Users", sl.Err(err))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Info("get Users successfully", sl.Atr("respUsers", users))

	c.JSON(http.StatusOK, users)
}

//add
func (a *apiImpl) addUserHandler(c *gin.Context) {

	const op = "addUsers"
	log := a.log.With(slog.String("op", op))

	user := &service.User{}
	if err := c.BindJSON(&user); err != nil {
		log.Error("cant unmarshall", sl.Err(err))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Info("run add User", sl.Atr("User", user))

	respUser, err := a.service.AddUser(c, user)
	if err != nil {
		a.log.Error("ocurred error for run add User", sl.Err(err))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Info("add User successfully", sl.Atr("respUser", respUser))

	c.JSON(http.StatusCreated, respUser)
}

//del
func (a *apiImpl) delUserHandler(c *gin.Context) {

	const op = "delUsers"
	log := a.log.With(slog.String("op", op))

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		a.log.Debug("id not match type", sl.Err(err))
		c.JSON(http.StatusBadRequest, sl.Err(err))
		return
	}

	log.Info("run del User by ID", sl.Atr("id", id))

	user, err := a.service.DeleteUser(c, id)
	if err != nil {
		a.log.Debug("ocurred error del User", sl.Err(err))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Info("del User successfully", sl.Atr("respUser", user))

	c.JSON(http.StatusOK, user)
}

//update
func (a *apiImpl) updateUserHandler(c *gin.Context) {
	const op = "updateUsers"
	log := a.log.With(slog.String("op", op))

	user := &service.User{}
	respUser := &service.User{}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		a.log.Debug("id not match type", sl.Err(err))
		c.JSON(http.StatusBadRequest, sl.Err(err))
		return
	}
	if err := c.BindJSON(&user); err != nil {
		a.log.Debug("cant unmarshall update User", sl.Err(err))
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user.Id = id

	log.Info("run update user", sl.Atr("User", user))

	respUser, err = a.service.UpdateUser(c, user)
	if err != nil {
		a.log.Debug("ocurred error update User", sl.Err(err))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Info("update User successfully", sl.Atr("respUser", respUser))

	c.JSON(http.StatusCreated, respUser)
}
