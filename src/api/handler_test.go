package api

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_ReadHandler_Input_Id_1_Should_Be_Id_1_Description_helloworld_Status_200(t *testing.T) {
	url := "/helloworld/read?id=1"
	request := httptest.NewRequest("GET", url, nil)
	responseWriter := httptest.NewRecorder()
	expectedStatusCode := 200
	expectedResultJSON := `{"id":1,"description":"hello world"}`

	ReadHandler(responseWriter, request)
	actualResponse := responseWriter.Result()
	body, _ := ioutil.ReadAll(actualResponse.Body)

	if expectedStatusCode != actualResponse.StatusCode {
		t.Errorf("status code is %d but got %d", expectedStatusCode, actualResponse.StatusCode)
	}
	if expectedResultJSON != string(body) {
		t.Errorf("result is\n'%s' \nbut got \n'%s'", expectedResultJSON, body)
	}
}
