package http

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"main/internal/user/domain/model"
	"net/http"
)

//func isValidUUID(id string) bool {
//	_, err := uuid.Parse(id)
//	return err == nil
//}

func (u UserAPI) Create(w http.ResponseWriter, r *http.Request) {

	k := r.Header.Get("User-Agent")
	if k == "" {
		return
		//	some smart checking
	}

	defer func() {
		_ = r.Body.Close()
	}()

	body, err := io.ReadAll(r.Body)

	if err != nil {
		// logger err read body
		return
	}

	user := model.User{}
	if err := json.Unmarshal(body, &user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// switch case on http status
		return
	}

	if err := u.useCase.Create(r.Context(), user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (u UserAPI) Update(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	//mid ??? mb don't used
	_, err := uuid.Parse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	defer func() {
		_ = r.Body.Close()
	}()

	user := model.User{}

	body, err := io.ReadAll(r.Body)
	err = json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = u.useCase.Update(r.Context(), user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)

}

func (u UserAPI) Read(w http.ResponseWriter, r *http.Request) {

	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := u.useCase.Read(r.Context(), model.ID(id))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		u.log.Err(err).Msg("failed to write response")
		return
	}

}

func (u UserAPI) Delete(w http.ResponseWriter, r *http.Request) {

	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = u.useCase.Delete(r.Context(), model.ID(id))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
