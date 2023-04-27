package handlers

import (
	productsdto "BE/dto/products"
	dto "BE/dto/result"
	"BE/models"
	"BE/repositories"
	"net/http"
	"strconv"
	"context"
	"os"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var ctx = context.Background()

type handlerProduct struct {
	ProductRepository repositories.ProductRepository
}

func HandlerProduct(ProductRepository repositories.ProductRepository) *handlerProduct {
	return &handlerProduct{ProductRepository}
}

func (h *handlerProduct) GetAllProduct(c echo.Context) error {
	products, err := h.ProductRepository.GetAllProduct()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: products})
}

func (h *handlerProduct) GetOneProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var product models.Product
	product, err := h.ProductRepository.GetOneProduct(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: product})
}

func (h *handlerProduct) CreateProduct(c echo.Context) error {
	request := new(productsdto.CreateProductRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	// get from middleware
	dataFile := c.Get("dataFile").(string)
	ImageCloud := ""
	if dataFile != "" {
		// Configuration
		cld, _ := cloudinary.NewFromParams(os.Getenv("CLOUD_NAME"), os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
		// Upload file to Cloudinary ...
		resp, err := cld.Upload.Upload(ctx, dataFile, uploader.UploadParams{Folder: os.Getenv("CLOUD_FOLDER")});
		if err != nil {
		fmt.Println(err.Error())
		}
		ImageCloud =  resp.SecureURL

	}else{
		ImageCloud =  ""
	}

	product := models.Product{
		Name:   		request.Name,
		Stock:  		request.Stock,
		Price:  		request.Price,
		Description:	request.Description,
		Image:  		ImageCloud,
	}

	product, err = h.ProductRepository.CreateProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	product, _ = h.ProductRepository.GetOneProduct(product.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: product})
}

func (h *handlerProduct) UpdateProduct(c echo.Context) error {
	request := new(productsdto.UpdateProductRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := h.ProductRepository.GetOneProduct(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Name != "" {
		product.Name = request.Name
	}

	if request.Stock != 0 {
		product.Stock = request.Stock
	}

	if request.Price != 0 {
		product.Price = request.Price
	}

	if request.Description != "" {
		product.Description = request.Description
	}

	// get from middleware
	dataFile := c.Get("dataFile").(string)
	if dataFile != "" {
		// Configuration
		cld, _ := cloudinary.NewFromParams(os.Getenv("CLOUD_NAME"), os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
		// Upload file to Cloudinary ...
		resp, err := cld.Upload.Upload(ctx, dataFile, uploader.UploadParams{Folder: os.Getenv("CLOUD_FOLDER")});
		if err != nil {
		fmt.Println(err.Error())
		}
		product.Image =  resp.SecureURL

	}
	
	data, err := h.ProductRepository.UpdateProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

func (h *handlerProduct) DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := h.ProductRepository.GetOneProduct(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.ProductRepository.DeleteProduct(product, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}
