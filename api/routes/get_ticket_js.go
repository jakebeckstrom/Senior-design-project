package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/csci4950tgt/api/models"
	"github.com/csci4950tgt/api/util"
	"github.com/gorilla/mux"
)

func GetTicketJS(w http.ResponseWriter, r *http.Request) {
	// Get variables from route handler
	vars := mux.Vars(r)                       // get dynamic variables from mux handler
	ticketId, err := strconv.Atoi(vars["id"]) // get integer "id" from vars

	if err != nil {
		util.WriteHttpErrorCode(w, http.StatusBadRequest, "Missing required parameter: id.")

		return
	}

	// TODO: Actually get js from ticket

	// Initialize Response
	msg := fmt.Sprintf("Not yet implemented... ticketId: %ds", ticketId)
	res := models.Response{
		Success: true,
		Message: &msg,
		// TODO: Send js in response
	}

	util.WriteHttpResponse(w, res)
}
