package handler

import (
	"Lab1/intermal/app/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Repository *repository.Repository
}

func NewHandler(r *repository.Repository) *Handler {
	return &Handler{
		Repository: r,
	}
}

func (h *Handler) GetOrders(ctx *gin.Context) {
	var orders []repository.Order
	var err error

	searchQuery := ctx.Query("query") // получаем значение из поля поиска
	if searchQuery == "" {            // если поле поиска пусто, то просто получаем из репозитория все записи
		orders, err = h.Repository.GetOrders()
		if err != nil {
			logrus.Error(err)
		}
	} else {
		orders, err = h.Repository.GetOrdersByTitle(searchQuery) // в ином случае ищем заказ по заголовку
		if err != nil {
			logrus.Error(err)
		}
	}

	cart, err := h.Repository.GetCart()
	if err != nil {
		logrus.Error(err)
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"orders": orders,
		"query":  searchQuery,
		"cart":   cart,
	})
}

func (h *Handler) GetOrder(ctx *gin.Context) {
	idStr := ctx.Param("id") // получаем id заказа из урла (то есть из /order/:id)
	// через двоеточие мы указываем параметры, которые потом сможем считать через функцию выше
	id, err := strconv.Atoi(idStr) // так как функция выше возвращает нам строку, нужно ее преобразовать в int
	if err != nil {
		logrus.Error(err)
	}

	// ДОБАВИТЬ ЭТУ СТРОКУ - получение корзины
	cart, err := h.Repository.GetCart()
	if err != nil {
		logrus.Error(err)
	}

	order, err := h.Repository.GetOrder(id)
	if err != nil {
		logrus.Error(err)
	}

	ctx.HTML(http.StatusOK, "order.html", gin.H{
		"order": order,
		"cart":  cart, // Добавляем корзину
	})
}

func (h *Handler) GetCart(ctx *gin.Context) {
	cart, err := h.Repository.GetCart()
	if err != nil {
		logrus.Error(err)
		ctx.HTML(http.StatusInternalServerError, "cart.html", gin.H{
			"error": "Ошибка загрузки корзины",
		})
		return
	}

	ctx.HTML(http.StatusOK, "cart.html", gin.H{
		"cart": cart,
	})
}
