package transport

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/Antoha2/sandbox/service"
	"github.com/gin-gonic/gin"
)

func (wImpl *webImpl) StartHTTP() error {
	router := gin.Default()
	router.GET("/users/", wImpl.getUserHandler)
	router.POST("/users/", wImpl.addUserHandler)
	router.DELETE("/users/:id", wImpl.delUserHandler)
	router.PATCH("/users/", wImpl.updateUserHandler)
	router.Run()
	return nil
}

//get
func (wImpl *webImpl) getUserHandler(c *gin.Context) {
	q := c.Request.URL.Query()
	age, err := strconv.Atoi(q.Get("age"))
	if err != nil {
		log.Println(err)
		return
	}
	user := service.User{
		Name:        q.Get("name"),
		SurName:     q.Get("surname"),
		Patronymic:  q.Get("patronymic"),
		Age:         age,
		Gender:      q.Get("gender"),
		Nationality: q.Get("nationality"),
	}

	err = wImpl.service.GetUsers(user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

//add
func (wImpl *webImpl) addUserHandler(c *gin.Context) {

	user := new(service.User)
	if err := c.BindJSON(&user); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	err := wImpl.service.AddUser(*user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, user)
}

//del
func (wImpl *webImpl) delUserHandler(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return
	}
	err = wImpl.service.DelUser(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, id)
}

//update
func (wImpl *webImpl) updateUserHandler(c *gin.Context) {
	user := new(service.User)
	if err := c.BindJSON(&user); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	err := wImpl.service.UpdateUser(*user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, user)
}

//декодеры JSON
func (wImpl *webImpl) Decoder(r *http.Request, user *service.User) error {

	body, err := ioutil.ReadAll(r.Body)
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
