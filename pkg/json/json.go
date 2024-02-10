package json

import (
	"encoding/json"
	"net/http"

	// "github.com/gin-gonic/gin"
	"project/pkg/err"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	erro := decoder.Decode(result)
	err.PanicIfError(erro)
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	erro := encoder.Encode(response)
	err.PanicIfError(erro)
}
