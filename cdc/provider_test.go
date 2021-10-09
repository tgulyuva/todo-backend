package cdc

import (
	"fmt"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func TestProvider_TodoGetPost(t *testing.T) {
	pact := &dsl.Pact{
		Provider: "todo-list-backend",
	}

	port := 3000

	pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL:            fmt.Sprintf("http://127.0.0.1:%d", port),
		BrokerURL:                  "https://gulyuva.pactflow.io",
		BrokerToken:                "KhJOPfQdxj90G8mg_W6quQ",
		PublishVerificationResults: true,
		ProviderVersion:            "asdasdasdas-asd121dasd-12dsadasd",

		BeforeEach: func() error {

			//	app := main.EndPoints()

			//go app.Listen(fmt.Sprint(":", port))
			//go http.ListenAndServe(fmt.Sprint("localhost:", port), app)

			return nil
		},
		StateHandlers: types.StateHandlers{
			"there is 3 todo": func() error {
				//err := repository.AddTodoRepository(&TodoModel{ID: "1", Todo: "test 1"})
				//err = repository.AddTodoRepository(&TodoModel{ID: "2", Todo: "test 2"})
				//err = repository.AddTodoRepository(&TodoModel{ID: "3", Todo: "test 3"})

				return nil
			},
			"there is 1 todo": func() error {
				//err := repository.AddTodoRepository(&TodoModel{ID: "4", Todo: "Buy the milk"})
				return nil
			},
		}},
	)
}
