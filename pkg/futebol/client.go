package futebol

import (
	"encoding/json"
	"fmt"
)

type EndpointID int

const (
	CAMPEONATOS EndpointID = iota
)

var Endpoints = map[EndpointID][]string{CAMPEONATOS: {"v1", "campeonatos"}}

type Client interface {
	Campeonatos() ([]CampeonatoData, error)
}

type futebolClient struct {
	client HTTPClient
}

func New(hc HTTPClient) Client {
	return &futebolClient{client: hc}
}

func (fc futebolClient) Campeonatos() ([]CampeonatoData, error) {
	camp := []CampeonatoData{}

	data, err := fc.client.Get("v1", "campeonatos")
	if err != nil {
		return camp, fmt.Errorf("error getting campeonatos data: %w", err)
	}

	// data, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return nil, fmt.Errorf("error reading response body: %w", err)
	// }

	err = json.Unmarshal(data, &camp)
	if err != nil {
		return camp, fmt.Errorf("error unmarshling data '%s': %w", string(data), err)
	}

	return camp, nil
}

type CampeonatoData struct {
	CampeonatoID int        `json:"campeonato_id"`
	Nome         string     `json:"nome"`
	Slug         string     `json:"slug"`
	NomePopular  string     `json:"nome_popular"`
	EdicaoAtual  EdicaoData `json:"edicao_atual"`
	FaseAtual    FaseData   `json:"fase_atual"`
	Status       string     `json:"status"`
	Tipo         string     `json:"tipo"`
	Logo         string     `json:"logo"`
	Regiao       string     `json:"regiao"`
	Link         string     `json:"_link"`
}

type EdicaoData struct {
	EdicaoID    int    `json:"edicao_id"`
	Temporada   string `json:"temporada"`
	Nome        string `json:"nome"`
	Slug        string `json:"slug"`
	NomePopular string `json:"nome_popular"`
}

type FaseData struct {
	FaseID int    `json:"fase_id"`
	Nome   string `json:"nome"`
	Slug   string `json:"slug"`
	Tipo   string `json:"tipo"`
	Link   string `json:"_link"`
}
