package app

import (
	"encoding/json"
	_ "errors"
	"io/ioutil"
	"net/http"

	"github.com/kosha/twilio-connector/pkg/httpclient"
	"github.com/kosha/twilio-connector/pkg/models"
)

// listConnectorSpecification godoc
// @Summary Get connector specification details
// @Description Retrieve necessary environment variables
// @Tags specification
// @Produce json
// @Success 200 {object} object
// @Router /api/v1/specification/list [get]
func (a *App) listConnectorSpecification(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	respondWithJSON(w, http.StatusOK, map[string]string{

		"USERNAME": "twilio Username",
		"PASSWORD": "twilio Password",
	})
}

// testConnectorSpecification godoc
// @Summary Test auth against the specification
// @Description Check if domain account can be verified
// @Tags specification
// @Accept  json
// @Produce  json
// @Param text body models.Specification false "Enter auth and domain name properties"
// @Success 200
// @Router /api/v1/specification/test [get]
func (a *App) testConnectorSpecification(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	if (*r).Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	var s models.Specification
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&s); err != nil {
		a.Log.Errorf("Error parsing json payload", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	account := httpclient.GetAccount(s.Username, s.Password, r.URL.Query())
	if account != "" {
		respondWithXML(w, http.StatusOK, account)
	} else {
		respondWithError(w, http.StatusBadRequest, "Account not verified")
	}

}

// sendWhatsAppMessage godoc
// @Summary Send a Message with the Twilio API for WhatsApp
// @Tags communication
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Success 200
// @Router /api/v1/message [post]
func (a *App) sendWhatsAppMessage(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(r.Body)
	} else {
		respondWithError(w, http.StatusBadRequest, "Empty Body")
	}

	res, err := httpclient.SendWhatsAppMessage(a.Cfg.GetUsername(), a.Cfg.GetPassword(), bodyBytes, r.URL.Query())
	if err != nil {
		a.Log.Errorf("Error in sendWhatsAppMessage", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	respondWithJSON(w, http.StatusOK, res)

}

// makePhoneCall godoc
// @Summary Make a phone call
// @Tags communication
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Success 200
// @Router /api/v1/call [post]
func (a *App) makePhoneCall(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(r.Body)
	} else {
		respondWithError(w, http.StatusBadRequest, "Empty Body")
	}

	res, err := httpclient.MakePhoneCall(a.Cfg.GetUsername(), a.Cfg.GetPassword(), bodyBytes, r.URL.Query())
	if err != nil {
		a.Log.Errorf("Error in makePhoneCall", err)
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	respondWithJSON(w, http.StatusOK, res)

}
