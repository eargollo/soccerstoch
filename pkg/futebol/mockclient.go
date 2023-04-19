package futebol

import (
	"strings"
)

type MockHTTPClient struct{}

func (mc *MockHTTPClient) Get(endpoint ...string) ([]byte, error) {
	endpointStr := strings.Join(endpoint, "/")

	json := ""

	switch endpointStr {
	case "v1/campeonatos":
		json = `[{"campeonato_id":14,"nome":"Campeonato Brasileiro S\u00e9rie B","slug":"campeonato-brasileiro-serie-b","nome_popular":"Brasileir\u00e3o S\u00e9rie B","edicao_atual":{"edicao_id":80,"temporada":"2023","nome":"Campeonato Brasileiro S\u00e9rie B 2023","nome_popular":"Brasileir\u00e3o S\u00e9rie B 2023","slug":"campeonato-brasileiro-serie-b-2023"},"fase_atual":{"fase_id":371,"nome":"Fase \u00danica","slug":"fase-unica","tipo":"pontos-corridos","_link":"\/v1\/campeonatos\/14\/fases\/371"},"rodada_atual":{"nome":"1\u00aa Rodada","slug":"1a-rodada","rodada":1,"status":"agendada"},"status":"andamento","tipo":"Pontos Corridos","logo":"https:\/\/api.api-futebol.com.br\/images\/competicao\/brasileiro-serieb.png","regiao":"nacional","_link":"\/v1\/campeonatos\/14"}]`
	}

	// create a new reader with that JSON
	// r := ioutil.NopCloser(bytes.NewReader([]byte(json)))

	// resp := &http.Response{
	// 	StatusCode: 200,
	// 	Body:       r,
	// }

	// return resp, nil

	return []byte(json), nil
}
