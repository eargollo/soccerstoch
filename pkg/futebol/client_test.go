package futebol

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

type TestHTTPClientMock struct {
	mock.Mock
}

func (o *TestHTTPClientMock) Get(endpoint ...string) ([]byte, error) {
	args := o.Called(endpoint)
	return args.Get(0).([]byte), args.Error(1)
}

func Test_futebolClient_Campeonatos(t *testing.T) {
	testHTTPClientA := new(TestHTTPClientMock)
	testHTTPClientA.On("Get", Endpoints[CAMPEONATOS]).Return(
		[]byte("[]"), nil,
	)
	testHTTPClientB := new(TestHTTPClientMock)
	testHTTPClientB.On("Get", Endpoints[CAMPEONATOS]).Return(
		[]byte(`[{"campeonato_id":2,"nome":"CopadoBrasil","slug":"copa-do-brasil","nome_popular":"CopadoBrasil","edicao_atual":{"edicao_id":69,"temporada":"2023","nome":"CopadoBrasil2023","nome_popular":"CopadoBrasil2023","slug":"copa-do-brasil-2023"},"fase_atual":{"fase_id":312,"nome":"TerceiraFase","slug":"terceira-fase","tipo":"mata-mata","_link":"/v1/campeonatos/2/fases/312"},"rodada_atual":null,"status":"andamento","tipo":"Mata-Mata","logo":"https://api.api-futebol.com.br/images/competicao/copa-do-brasil.png","regiao":"nacional","_link":"/v1/campeonatos/2"}]`), nil,
	)
	testHTTPClientC := new(TestHTTPClientMock)
	testHTTPClientC.On("Get", Endpoints[CAMPEONATOS]).Return(
		[]byte("[]"), fmt.Errorf("error"),
	)
	testHTTPClientD := new(TestHTTPClientMock)
	testHTTPClientD.On("Get", Endpoints[CAMPEONATOS]).Return(
		[]byte("[not a json]"), nil,
	)
	tests := []struct {
		name    string
		client  HTTPClient
		want    []CampeonatoData
		wantErr bool
	}{
		{name: "empty", client: testHTTPClientA, want: []CampeonatoData{}, wantErr: false},
		{name: "sample", client: testHTTPClientB, want: []CampeonatoData{
			{CampeonatoID: 2, Nome: "CopadoBrasil", Slug: "copa-do-brasil", NomePopular: "CopadoBrasil", EdicaoAtual: EdicaoData{EdicaoID: 69, Temporada: "2023", Nome: "CopadoBrasil2023", Slug: "copa-do-brasil-2023", NomePopular: "CopadoBrasil2023"}, FaseAtual: FaseData{FaseID: 312, Nome: "TerceiraFase", Slug: "terceira-fase", Tipo: "mata-mata", Link: "/v1/campeonatos/2/fases/312"}, Status: "andamento", Tipo: "Mata-Mata", Logo: "https://api.api-futebol.com.br/images/competicao/copa-do-brasil.png", Regiao: "nacional", Link: "/v1/campeonatos/2"},
		}, wantErr: false},
		{name: "error get", client: testHTTPClientC, want: []CampeonatoData{}, wantErr: true},
		{name: "error unmarshall", client: testHTTPClientD, want: []CampeonatoData{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fc := futebolClient{
				client: tt.client,
			}
			got, err := fc.Campeonatos()
			if (err != nil) != tt.wantErr {
				t.Errorf("futebolClient.Campeonatos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("futebolClient.Campeonatos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		hc   HTTPClient
		want Client
	}{
		{"client", &MockHTTPClient{}, &futebolClient{client: &MockHTTPClient{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.hc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
