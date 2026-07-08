package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	dto "github.com/Izabelly-Melo/goexpert/api/internal/dto"
	"github.com/Izabelly-Melo/goexpert/api/internal/entity"
	"github.com/Izabelly-Melo/goexpert/api/internal/infra/database"
	entityPkg "github.com/Izabelly-Melo/goexpert/api/pkg/entity"
	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	ProductDB *database.Product
}

func NewProductHandler(db *database.Product) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

// CreateProduct godoc
// @Summary      Create a product
// @Description  Create a new product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        request body dto.CreateProductInput true "product request"
// @Success      201 {object} entity.Product
// @Failure      400 {object} Error
// @Failure      500
// @Router       /products [post]
// @Security     ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := struct {
			Menssage string `json:"message"`
		}{
			Menssage: err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return
	}

	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := struct {
			Menssage string `json:"message"`
		}{
			Menssage: err.Error(),
		}
		json.NewEncoder(w).Encode(msg)
		return
	}

	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

// GetProduct godoc
// @Summary      Get a product
// @Description  Get a product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id path string true "product ID"
// @Success      200 {object} entity.Product
// @Failure      400
// @Router       /products/{id} [get]
// @Security     ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// UpdateProduct godoc
// @Summary      Update a product
// @Description  Update an existing product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id path string true "product ID"
// @Param        request body dto.CreateProductInput true "product request"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /products/{id} [put]
// @Security     ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		msg := struct {
			Menssage string `json:"message"`
		}{
			Menssage: "ID is required",
		}
		json.NewEncoder(w).Encode(msg)
		return
	}

	existing, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var product entity.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product.ID, err = entityPkg.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product.CreatedAt = existing.CreatedAt

	err = h.ProductDB.Update(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// DeleteProduct godoc
// @Summary      Delete a product
// @Description  Delete a product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id path string true "product ID"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /products/{id} [delete]
// @Security     ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err := h.ProductDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = h.ProductDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetProducts godoc
// @Summary      List products
// @Description  Get all products, paginated
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        page  query string false "page"
// @Param        limit query string false "limit"
// @Param        sort  query string false "sort" Enums(asc, desc)
// @Success      200 {array} entity.Product
// @Failure      500
// @Router       /products [get]
// @Security     ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}
	sort := r.URL.Query().Get("sort")

	products, err := h.ProductDB.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
