package utils

import(
	"net/http"
	"encoding/json"
	"io"
)

func parseBody(r *http.Request,x interface{}){

	if body ,err := io.ReadAll(r.Body); err == nil{
		if  err:= json.Unmarshal([]byte(body),x); err != nil {
			return
		}
	}

}