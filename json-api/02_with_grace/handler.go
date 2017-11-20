package main

import (
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
	proverbs []Proverb
}

func newHandler(proverbs []Proverb) *handler {
	return &handler{proverbs: proverbs}
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

	p.ID = len(h.proverbs) + 1
	h.proverbs = append(h.proverbs, p)

	w.Header().Set("Location", fmt.Sprintf("/proverbs/%d", p.ID))
	w.WriteHeader(http.StatusCreated)
}

func (h *handler) getProverbs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.proverbs)
}

func (h *handler) getProverb(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, p := range h.proverbs {
		if id == p.ID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.NotFound(w, r)
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
	}
	defer r.Body.Close()

	var update Proverb
	if err := json.Unmarshal(body, &update); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if update.Text == "" || update.Philosopher == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := h.lookupProverb(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	p.Text = update.Text
	p.Philosopher = update.Philosopher
	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) deleteProverb(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for i, p := range h.proverbs {
		if id == p.ID {
			h.proverbs = append(h.proverbs[:i], h.proverbs[i+1:]...)
			w.WriteHeader(http.StatusAccepted)
			return
		}
	}

	http.NotFound(w, r)
}

var errProverbNotFound = errors.New("proverb not found")

func (h *handler) lookupProverb(id int) (*Proverb, error) {
	for i, p := range h.proverbs {
		if id == p.ID {
			return &h.proverbs[i], nil
		}
	}
	return nil, errProverbNotFound
}
