package transport

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Antoha2/sandbox/internal/service"
	"github.com/Antoha2/sandbox/pkg/logger/sl"
	"github.com/gin-gonic/gin"
)

func (a *apiImpl) StartHTTP() error {
	// должна быть возможность запускать сервер на определенном порту, который мы получаем из конфига
	router := gin.Default()
	router.GET("/users/:id", a.getUserHandler)    //get user
	router.GET("/users/", a.getUsersHandler)      //get userS
	router.POST("/users/", a.addUserHandler)      //add user
	router.DELETE("/users/:id", a.delUserHandler) //del user
	router.PUT("/users/:id", a.updateUserHandler) //update user
	err := router.Run()
	if err != nil {
		return err
	}
	return nil
}

// get user
func (a *apiImpl) getUserHandler(c *gin.Context) {
	const op = "getUser"

	log := a.log.With(
		slog.String("op", op),
	)
	log.Info("attempting to get user")

	id, err := strconv.Atoi(c.Param(ID))
	if err != nil {
		a.log.Error("id not match type", sl.Err(err))
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := a.service.GetUser(c, id)
	if err != nil {
		a.log.Error("occurred error for GetUser", sl.Err(err))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Info("request getUser completed successfully") // добавь в  логирование юзера которого возвращаем в запросе
	c.JSON(http.StatusOK, user)
}

// get users
func (a *apiImpl) getUsersHandler(c *gin.Context) {

	const op = "getUsers"

	log := a.log.With(
		slog.String("op", op),
	)
	log.Info("attempting to get users")

	var err error
	age := service.DefaultPropertyAge
	limit := service.DefaultPropertyLimit
	offset := service.DefaultPropertyOffset

	q := c.Request.URL.Query()

	qAge := q.Get(AGE)
	if qAge != "" {
		age, err = strconv.Atoi(qAge)
		if err != nil {
			a.log.Error("Age not match type", sl.Err(err))
			c.JSON(http.StatusBadRequest, sl.Err(err))
			return
		}
	}

	qOffset := q.Get(OFFSET)
	if qOffset != "" {
		offset, err = strconv.Atoi(qOffset)
		if err != nil {
			a.log.Error("offset not match type", sl.Err(err))
			c.JSON(http.StatusBadRequest, sl.Err(err))
			return
		}
	}

	qLimit := q.Get(LIMIT)
	if qLimit != "" {
		limit, err = strconv.Atoi(qLimit)
		if err != nil {
			a.log.Error("limit not match type", sl.Err(err))
			c.JSON(http.StatusBadRequest, sl.Err(err))
			return
		}
	}

	userQuery := &service.QueryUsersFilter{
		// все строки в константы
		Name:        q.Get("name"),
		SurName:     q.Get("surname"),
		Patronymic:  q.Get("patronymic"),
		Age:         age,
		Gender:      q.Get("gender"),
		Nationality: q.Get("nationality"),
		Offset:      offset,
		Limit:       limit,
	}

	users, err := a.service.GetUsers(c, userQuery)
	if err != nil {
		a.log.Error("runtime error GetUsers", sl.Err(err))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Info("request getUsers completed successfully") // add more info
	c.JSON(http.StatusOK, users)
}

// add
func (a *apiImpl) addUserHandler(c *gin.Context) {
	const op = "addUsers"

	log := a.log.With(
		slog.String("op", op),
	)
	log.Info("attempting to add user")

	user := service.User{}
	if err := c.BindJSON(&user); err != nil {
		log.Error("cant unmarshall", sl.Err(err))
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	respUser, err := a.service.AddUser(c, &user)
	if err != nil {
		a.log.Error("runtime error", sl.Err(err))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Info("request completed successfully") // info
	c.JSON(http.StatusCreated, respUser)
}

// del
func (a *apiImpl) delUserHandler(c *gin.Context) {
	const op = "delUsers"

	log := a.log.With(
		slog.String("op", op),
	)
	log.Info("attempting to del user")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		a.log.Debug("id not match type", sl.Err(err))
		c.JSON(http.StatusBadRequest, sl.Err(err))
		return
	}

	user, err := a.service.DelUser(c, id)
	if err != nil {
		a.log.Debug("runtime error delUser", sl.Err(err))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Info("request delUser completed successfully")
	c.JSON(http.StatusOK, user)
}

// update
func (a *apiImpl) updateUserHandler(c *gin.Context) {
	const op = "updateUsers"

	log := a.log.With(
		slog.String("op", op),
	)

	user := service.User{}
	id, err := strconv.Atoi(c.Param(ID))
	if err != nil {
		a.log.Debug("id not match type", sl.Err(err))
		c.JSON(http.StatusBadRequest, sl.Err(err))
		return
	}

	if err := c.BindJSON(&user); err != nil {
		a.log.Debug("cant unmarshall updateUser", sl.Err(err))
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user.Id = id
	respUser, err := a.service.UpdateUser(c, &user)
	if err != nil {
		a.log.Debug("runtime error updateUser", sl.Err(err))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	log.Info("request updateUser completed successfully")
	c.JSON(http.StatusCreated, respUser)
}
