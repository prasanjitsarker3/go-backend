package response

import (
	"encoding/json"
	"net/http"
)




func SendData (w http.ResponseWriter, data interface {}, statusCode int){
     w.WriteHeader(statusCode);
	 encode:= json.NewEncoder(w);
	 encode.Encode(data)
}
