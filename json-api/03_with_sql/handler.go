package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Proverb states a general truth or piece of advice.
type Proverb struct {
	ID          int    `json:"id"`
	Text        string `json:"text"`
	Philosopher string `json:"philosopher,omitempty"`
}

type handler struct {
	db *sql.DB
}

func newHandler(db *sql.DB) *handler {
	return &handler{db}
}

func (h *handler) createProverb(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer r.Body.Close()

	var p Proverb
	if err := json.Unmarshal(body, &p); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if p.Text == "" || p.Philosopher == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	stmt, err := h.db.Prepare(sqlInsert)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := stmt.Exec(p.Text, p.Philosopher)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/proverbs/%d", id))
	w.WriteHeader(http.StatusCreated)
}

func (h *handler) getProverbs(w http.ResponseWriter, r *http.Request) {
	rows, err := h.db.Query(sqlSelectAll)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var proverbs []Proverb
	for rows.Next() {
		p := Proverb{}
		if err := rows.Scan(&p.ID, &p.Text, &p.Philosopher); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		proverbs = append(proverbs, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&proverbs)
}

func (h *handler) getProverb(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	p, err := h.lookupProverb(id)
	if err != nil {
		if err == errProverbNotFound {
			http.NotFound(w, r)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&p)
}

func (h *handler) updateProverb(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var update Proverb
	if err := json.Unmarshal(body, &update); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if update.Text == "" || update.Philosopher == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	existing, err := h.lookupProverb(id)
	if err != nil {
		if err == errProverbNotFound {
			http.NotFound(w, r)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	stmt, err := h.db.Prepare(sqlUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = stmt.Exec(update.Text, update.Philosopher, existing.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) deleteProverb(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	existing, err := h.lookupProverb(id)
	if err != nil {
		if err == errProverbNotFound {
			http.NotFound(w, r)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	stmt, err := h.db.Prepare(sqlDelete)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = stmt.Exec(existing.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

var errProverbNotFound = errors.New("proverb not found")

func (h *handler) lookupProverb(id int) (*Proverb, error) {
	var p Proverb
	row := h.db.QueryRow(sqlSelectOne, id)
	err := row.Scan(&p.ID, &p.Text, &p.Philosopher)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errProverbNotFound
		}
		return nil, err
	}
	return &p, nil
}
