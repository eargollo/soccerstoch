package futebol_test

import (
	"reflect"
	"testing"

	"github.com/eargollo/soccrtstoch/pkg/futebol"
)

func TestMockHTTPClient_Get(t *testing.T) {
	tests := []struct {
		name     string
		mc       *futebol.MockHTTPClient
		endpoint []string
		want     []byte
		wantErr  bool
	}{
		{"campeonatos", &futebol.MockHTTPClient{}, futebol.Endpoints[futebol.CAMPEONATOS], []byte(`[{"campeonato_id":14,"nome":"Campeonato Brasileiro S\u00e9rie B","slug":"campeonato-brasileiro-serie-b","nome_popular":"Brasileir\u00e3o S\u00e9rie B","edicao_atual":{"edicao_id":80,"temporada":"2023","nome":"Campeonato Brasileiro S\u00e9rie B 2023","nome_popular":"Brasileir\u00e3o S\u00e9rie B 2023","slug":"campeonato-brasileiro-serie-b-2023"},"fase_atual":{"fase_id":371,"nome":"Fase \u00danica","slug":"fase-unica","tipo":"pontos-corridos","_link":"\/v1\/campeonatos\/14\/fases\/371"},"rodada_atual":{"nome":"1\u00aa Rodada","slug":"1a-rodada","rodada":1,"status":"agendada"},"status":"andamento","tipo":"Pontos Corridos","logo":"https:\/\/api.api-futebol.com.br\/images\/competicao\/brasileiro-serieb.png","regiao":"nacional","_link":"\/v1\/campeonatos\/14"}]`), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := &futebol.MockHTTPClient{}
			got, err := mc.Get(tt.endpoint...)
			if (err != nil) != tt.wantErr {
				t.Errorf("MockHTTPClient.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockHTTPClient.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
