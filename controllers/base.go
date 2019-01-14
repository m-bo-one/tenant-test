package controllers

type ApiJson struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func ApiResource(status string, objects interface{}) (apijson *ApiJson) {
	apijson = &ApiJson{Status: status, Data: objects}
	return
}
