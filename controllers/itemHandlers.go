package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rahulkherajani/backend-repo/models"
	"github.com/rahulkherajani/backend-repo/utils"
)

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var items []models.Item
	models.DB.Find(&items)

	json.NewEncoder(w).Encode(items)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var item models.Item

	res := models.DB.First(&item, id)

	if res.Error != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Item not found")
		return
	}

	json.NewEncoder(w).Encode(item)
}

type ItemInput struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input ItemInput
	body, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(body, &input)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Couldnot parse request body")
		return
	}

	validate := validator.New()
	err = validate.Struct(input)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Validation error")
		return
	}

	item := models.Item{
		ID:          uuid.New(),
		Title:       input.Title,
		Description: input.Description,
	}

	models.DB.Create(&item)

	json.NewEncoder(w).Encode(item)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var item models.Item

	res := models.DB.First(&item, id)

	if res.Error != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Item not found")
		return
	}

	var input ItemInput
	body, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(body, &input)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Couldnot parse request body")
		return
	}

	validate := validator.New()
	err = validate.Struct(input)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Validation error")
		return
	}

	item.Title = input.Title
	item.Description = input.Description

	models.DB.Save(&item)

	json.NewEncoder(w).Encode(item)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var item models.Item

	res := models.DB.First(&item, id)

	if res.Error != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Item not found")
		return
	}

	models.DB.Delete(&item)

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(item)
}
