package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/JosacabDev/go_api/model"
)

type dbHandler interface {
	Create(person *model.Person) error
	Update(ID int, person *model.Person) error
	Delete(ID int) error
	GetByID(ID int) (model.Person, error)
	GetAll() (model.Peopel, error)
}

type person struct {
	dbHandler dbHandler
}

func newPerson(dbHandler dbHandler) person {
	return person{dbHandler}
}

func (p *person) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type:" "error", "message:" "Invalid Method"}`))
		return
	}

	data := model.Person{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type:" "error", "message:" "Invalid Data Struct"}`))
		return
	}

	err = p.dbHandler.Create(&data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message_type:" "error", "message:" "Can not Created a Person"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message_type:" "message", "message:" "Person Created Successfully"}`))

}

func (p *person) update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type:" "error", "message:" "Invalid Method"}`))
		return
	}

	ID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type:" "error", "message:" "ID should be Type Int"}`))
		return
	}

	data := model.Person{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type:" "error", "message:" "Invalid Data Struct"}`))
		return
	}

	err = p.dbHandler.Update(ID, &data)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message_type:" "error", "message:" "Can not Get any Data"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message_type:" "message", "message:" "Person Update Successfully"}`))
}

func (p *person) getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type:" "error", "message:" "Invalid Method"}`))
		return
	}

	resp, err := p.dbHandler.GetAll()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message_type:" "error", "message:" "Can not Get any Data"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message_type:" "error", "message:" "Can not convert to JSON"}`))
		return
	}
}
