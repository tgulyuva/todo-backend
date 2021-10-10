package cdc

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/tgulyuva/todo-backend"
)

func TestProvider_TodoPost(t *testing.T) {
	pact := &dsl.Pact{
		Provider: "todo-list-backend",
	}

	port := 3001

	pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL:            fmt.Sprintf("http://127.0.0.1:%d", port),
		BrokerURL:                  "https://gulyuva.pactflow.io",
		BrokerToken:                "KhJOPfQdxj90G8mg_W6quQ",
		PublishVerificationResults: true,
		ProviderVersion:            "asdasdasdas-asd121dasd-12dsadasd",

		BeforeEach: func() error {

			app := main.EndPoints()
			go http.ListenAndServe(fmt.Sprint("localhost:", port), app)

			return nil
		},
		StateHandlers: types.StateHandlers{

			"there is task": func() error {
				// msg := models.ResponseModel{Success: false}
				// return msg
				return nil // hata almak için nil dönüş yapıldı.
			},
		}},
	)
}
