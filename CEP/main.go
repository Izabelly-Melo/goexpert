package main

import (
	"encoding/json"
	"net/http"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	http.HandleFunc("/", SearchCepHandler)
	http.ListenAndServe(":8080", nil)

}

func SearchCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cepData, err := SearchCep(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	/*
		result, err := json.Marshal(cepData)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(result)
	*/

	json.NewEncoder(w).Encode(cepData)
}

func SearchCep(cep string) (*ViaCEP, error) {
	res, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body := &ViaCEP{}
	err = json.NewDecoder(res.Body).Decode(body)
	if err != nil {
		return nil, err
	}

	/*
		var dado ViaCEP
		error := json.Unmarshal(body, &dado)
		if error != nil {
			return nil, error
		}
		return &dado, nil
	*/

	return body, nil
}
