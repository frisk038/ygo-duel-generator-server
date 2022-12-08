package v1

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type stunBusiness interface {
	GetAnswerICE(ctx context.Context, roomID string) (*string, error)
	GetAnswer(ctx context.Context, roomID string) (*string, error)
	GetOfferICE(ctx context.Context, roomID string) (*string, error)
	GetOffer(ctx context.Context, roomID string) (*string, error)
	StoreAnswerICE(ctx context.Context, roomID string, ice []byte) error
	StoreAnswer(ctx context.Context, roomID string, answer []byte) error
	StoreOfferICE(ctx context.Context, roomID string, ice []byte) error
	StoreOffer(ctx context.Context, roomID string, offer []byte) error
}

type StunHandler struct {
	business stunBusiness
}

func NewStunHandler(b stunBusiness) *StunHandler {
	return &StunHandler{business: b}
}

func (h *StunHandler) GetOfferHandler(c *gin.Context) {
	of, err := h.business.GetOffer(c.Request.Context(), c.Param("roomid"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, of)
}

func (h *StunHandler) GetOfferICEHandler(c *gin.Context) {
	of, err := h.business.GetOfferICE(c.Request.Context(), c.Param("roomid"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, of)
}

func (h *StunHandler) GetAnswerHandler(c *gin.Context) {
	of, err := h.business.GetAnswer(c.Request.Context(), c.Param("roomid"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, of)
}

func (h *StunHandler) GetAnswerICEHandler(c *gin.Context) {
	of, err := h.business.GetAnswerICE(c.Request.Context(), c.Param("roomid"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, of)
}

func (h *StunHandler) PostOfferHandler(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	err = h.business.StoreOffer(c.Request.Context(), c.Param("roomid"), body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *StunHandler) PostOfferICEHandler(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	err = h.business.StoreOfferICE(c.Request.Context(), c.Param("roomid"), body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *StunHandler) PostAnswerHandler(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	fmt.Println("body", string(body))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	err = h.business.StoreAnswer(c.Request.Context(), c.Param("roomid"), body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *StunHandler) PostAnswerICEHandler(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	err = h.business.StoreAnswerICE(c.Request.Context(), c.Param("roomid"), body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
