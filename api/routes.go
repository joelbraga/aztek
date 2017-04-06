package api

import (
	"github.com/gin-gonic/gin"
	"github.com/joelbraga/aztek/action"
)

func (api *ApiResource) GetAll(c *gin.Context, model interface{}) {
	if products, err := api.Repository.GetAll(model); err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, api.HttpResponse("ok", "GetAll", products))
	}
}

func (api *ApiResource) Get(c *gin.Context, model interface{}, preload []string) {
	id := c.Param("id")
	if product, err := api.Repository.GetById(id, model, preload); err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, api.HttpResponse("ok", "Get", product))
	}
}

func (api *ApiResource) GetByCode(c *gin.Context, model interface{}, preload []string) {
	code := c.Param("code")
	if product, err := api.Repository.GetByCode(code, model, preload); err != nil {
		c.AbortWithStatus(404)
	} else {
		c.JSON(200, api.HttpResponse("ok", "GetByCode", product))
	}
}

func (api *ApiResource) Update(c *gin.Context, model interface{}) {
	id := c.Param("id")
	if product, err := api.Repository.GetById(id, model, nil); err != nil {
		c.AbortWithStatus(404)
	} else {
		if bindErr := c.BindJSON(product); bindErr == nil {
			err = api.Repository.Update(id, product)
			if err != nil {
				c.AbortWithStatus(404)
			} else {
				if api.ActionEvent != nil {
					api.ActionEvent.AddEvent(action.UPDATE_MODEL, product)
				}
				c.JSON(200, api.HttpResponse("ok", "Updated", product))
			}
		}
	}
}

func (api *ApiResource) Delete(c *gin.Context, model interface{}) {
	id := c.Param("id")
	if err := api.Repository.Delete(id, model); err != nil {
		c.AbortWithStatus(404)
	} else {
		if api.ActionEvent != nil {
			api.ActionEvent.AddEvent(action.DELETE_MODEL, id)
		}
		c.JSON(200, api.HttpResponse("ok", "Deleted", "Deleted"))
	}
}
