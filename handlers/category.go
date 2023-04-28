package handlers

import (
	categorydto "BE/dto/category"
	dto "BE/dto/result"
	"BE/models"
	"BE/repositories"
	"net/http"
	"strconv"
	"os"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/golang-jwt/jwt/v4"
)

type handlerCategory struct {
	CategoryRepository repositories.CategoryRepository
	CartRepository     repositories.CartRepository
}

func HandlerCategory(CategoryRepository repositories.CategoryRepository, CartRepository repositories.CartRepository) *handlerCategory {
	return &handlerCategory{
		CategoryRepository: CategoryRepository,
		CartRepository:     CartRepository,
	}
}

func (h *handlerCategory) GetAllCategory(c echo.Context) error {
	Category, err := h.CategoryRepository.GetAllCategory()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: Category})
}

func (h *handlerCategory) GetOneCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var Category models.Category
	Category, err := h.CategoryRepository.GetOneCategory(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: Category})
}

func (h *handlerCategory) CreateCategory(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := int(userLogin.(jwt.MapClaims)["id"].(float64))

	// Check if email already confirmed
	_, error := h.CartRepository.GetUserEmailStats(userId)
	if error != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Email has not been confirmed"})
	}

	request := new(categorydto.CreateCategoryRequest)
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

	Category := models.Category{
		Name:  request.Name,
		Image: ImageCloud,
	}

	Category, err = h.CategoryRepository.CreateCategory(Category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	Category, _ = h.CategoryRepository.GetOneCategory(Category.ID)

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: Category})
}

func (h *handlerCategory) UpdateCategory(c echo.Context) error {
	request := new(categorydto.UpdateCategoryRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	
	id, _ := strconv.Atoi(c.Param("id"))

	Category, err := h.CategoryRepository.GetOneCategory(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Name != "" {
		Category.Name = request.Name
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
		Category.Image =  resp.SecureURL

	}
	
	data, err := h.CategoryRepository.UpdateCategory(Category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}

func (h *handlerCategory) DeleteCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	Category, err := h.CategoryRepository.GetOneCategory(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.CategoryRepository.DeleteCategory(Category, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: data})
}
