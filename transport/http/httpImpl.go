package transport

import (
	"context"
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
	router.Run()                                  // обработка ошибки
	return nil
}

// get user
func (w *webImpl) getUserHandler(c *gin.Context) {
	// зачем ты создаешь новый контекст?? есть же от джина его и используй
	ctx := context.Background()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err) // slog!!
		return           // надо 400 возвращать
	}
	user := new(service.User) // не надо так делать, зачем зря выделять память а потом сразу переписывать
	user, err = w.service.GetUser(ctx, id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

// get users
func (w *webImpl) getUsersHandler(c *gin.Context) {
	ctx := context.Background()
	// const op = "getUser"
	// w.log.With()
	// log := w.log.With(
	// 	slog.String("op", op),
	// 	slog.String("username", email),
	// )

	// log.Info("attempting to login user")
	// дефолтные значения на уровень сервиса и вынести к константы
	age := 0
	limit := 100
	offset := 0
	var err error
	q := c.Request.URL.Query()

	// по хорошему все стринги тоже нужно в константы
	// не надо два раза вызывать одну и туже функцию , например q.Get("age")
	if q.Get("age") != "" {
		age, err = strconv.Atoi(q.Get("age"))
		if err != nil {
			// юзер возможно просто не передал age в запросе а ты сразу за ошибку считаешь
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
	userQuery := &service.GetQueryFilter{

		Name:        q.Get("name"),
		SurName:     q.Get("surname"),
		Patronymic:  q.Get("patronymic"),
		Age:         age,
		Gender:      q.Get("gender"),
		Nationality: q.Get("nationality"),
		Offset:      offset,
		Limit:       limit,
	}
	users, err := w.service.GetUsers(ctx, userQuery)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, users)
}

// add
func (w *webImpl) addUserHandler(c *gin.Context) {
	ctx := context.Background()
	user := &service.User{}     // единообразие!! через new
	respUser := &service.User{} // это лишнее выделение памяти, эта строка вообще не нужна

	if err := c.BindJSON(&user); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error()) // это 400 а не 500
		return
	}
	respUser, err := w.service.AddUser(ctx, user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, respUser) // 201, а вот это хорошо
}

// del
func (w *webImpl) delUserHandler(c *gin.Context) {
	ctx := context.Background()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return // 400
	}

	err = w.service.DelUser(ctx, id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, id)
}

// update
func (w *webImpl) updateUserHandler(c *gin.Context) {
	ctx := context.Background()
	user := new(service.User)
	respUser := new(service.User) // не нужно
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		return
	}
	if err := c.BindJSON(&user); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error()) // 400
		return
	}
	user.Id = id
	respUser, err = w.service.UpdateUser(ctx, user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, respUser)
}

// не используется
// декодеры JSON
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
