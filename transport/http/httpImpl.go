package transport

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/Antoha2/sandbox/service"
	"github.com/gin-gonic/gin"
)

func (w *webImpl) StartHTTP() error {
	router := gin.Default()
	router.GET("/users/:id", w.getUserHandler)    //get user
	router.GET("/users/", w.getUsersHandler)      //get userS
	router.POST("/users/", w.addUserHandler)      //add user
	router.DELETE("/users/:id", w.delUserHandler) //del user
	router.PUT("/users/:id", w.updateUserHandler) //update user
	router.Run()
	return nil
}

//get user
func (w *webImpl) getUserHandler(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return
	}
	user := new(service.User)
	*user, err = w.service.GetUser(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

//get users
func (w *webImpl) getUsersHandler(c *gin.Context) {
	// const op = "getUser"
	// w.log.With()
	// log := w.log.With(
	// 	slog.String("op", op),
	// 	slog.String("username", email),
	// )

	// log.Info("attempting to login user")
	age := 0
	limit := 100
	offset := 0
	var err error
	q := c.Request.URL.Query()

	if q.Get("age") != "" {
		age, err = strconv.Atoi(q.Get("age"))
		if err != nil {
			log.Println(err)
			//return
		}
	}
	if q.Get("offset") != "" {
		offset, err = strconv.Atoi(q.Get("offset"))
		if err != nil {
			log.Println(err)
			//return
		}
	}
	if q.Get("limit") != "" {
		limit, err = strconv.Atoi(q.Get("limit"))
		if err != nil {
			log.Println(err)
			//return
		}
	}
	userQuery := service.GetQueryFilter{

		Name:        q.Get("name"),
		SurName:     q.Get("surname"),
		Patronymic:  q.Get("patronymic"),
		Age:         age,
		Gender:      q.Get("gender"),
		Nationality: q.Get("nationality"),
		Offset:      offset,
		Limit:       limit,
	}
	users, err := w.service.GetUsers(userQuery)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, users)
}

//add
func (w *webImpl) addUserHandler(c *gin.Context) {

	var user, respUser service.User

	if err := c.BindJSON(&user); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	respUser, err := w.service.AddUser(user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, respUser)
}

//del
func (w *webImpl) delUserHandler(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return
	}

	err = w.service.DelUser(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, id)
}

//update
func (w *webImpl) updateUserHandler(c *gin.Context) {
	user := new(service.User)
	respUser := new(service.User)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return
	}
	if err := c.BindJSON(&user); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	user.Id = id
	*respUser, err = w.service.UpdateUser(*user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, respUser)
}

//декодеры JSON
func (w *webImpl) Decoder(r *http.Request, user *service.User) error {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, user)
	if err != nil {
		fmt.Println("can't unmarshal !!!!!: ", err.Error())
		return err
	}
	return nil
}
