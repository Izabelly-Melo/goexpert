package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Izabelly-Melo/goexpert/api/internal/dto"
	"github.com/Izabelly-Melo/goexpert/api/internal/entity"
	"github.com/Izabelly-Melo/goexpert/api/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type UserHandler struct {
	UserDB *database.User
}

func NewUserHandler(db *database.User) *UserHandler {
	return &UserHandler{
		UserDB: db,
	}
}

type Error struct {
	Message string `json:"message"`
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetToken(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("jwtExpiresIn").(int)
	var user dto.GetTokenInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := h.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	if !u.CheckPassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		error := Error{Message: "invalid password"}
		json.NewEncoder(w).Encode(error)
		return
	}

	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"user_id": u.ID.String(),
		"exp":     time.Now().Add(time.Duration(jwtExpiresIn) * time.Second).Unix(),
	})
	acessToken := dto.GetTokenOutput{AcessToken: tokenString}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(acessToken)
}
