package handlers_tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dundudun/rest_test_back/db/sqlc"
	"github.com/gin-gonic/gin"
)

func FuzzCreateOrg(f *testing.F) {
	cleanup := SetupDB()
	defer cleanup()

	//TOCHECK: can Fuzz generate nil?
	f.Add(63, "test_org_from_fuzzing_1", "100", "200", "10", "20", "200", "0")
	f.Add(62, "test_org_from_fuzzing_2", "100", "200", "10", "20", "200")
	//TODO: more cases for f.Add()

	// fieldSet sets what optional fields will be in request body (1 means will be present, 0 - won't)
	//   63 -> 111111 means all 6 optional fields will be present
	//   39 -> 111001 means optional fields {plasticLimit, glassLimit, biowasteLimit and producedBiowaste} will be present
	f.Fuzz(func(t *testing.T, fieldSet int, name, plasticLimit, glassLimit, biowasteLimit, producedPlastic, producedGlass, producedBiowaste string) {
		// data->json->reqsBody
		// h.CreateOrganization(req)
		// check statusOk
		// row select ... where ... ->parse to json
		// compare jsons - send and got

		body := `{"name":` + name
		i := 0
		for fieldSet %= 64; fieldSet != 0; fieldSet /= 2 {
			if fieldSet%2 == 1 {
				switch i {
				case 0:
					body += `,"plastic_limit":"` + plasticLimit + `"`
				case 1:
					body += `,"glass_limit":"` + glassLimit + `"`
				case 2:
					body += `,"biowaste_limit":"` + biowasteLimit + `"`
				case 3:
					body += `,"produced_plastic":"` + producedPlastic + `"`
				case 4:
					body += `,"produced_glass":"` + producedGlass + `"`
				case 5:
					body += `,"produced_biowaste":"` + producedBiowaste + `"`
				}
			}
			i++
		}
		body += `}`

		router := gin.Default()
		router.POST("/api/organizations", handler.CreateOrganization)

		recorder := httptest.NewRecorder()
		//move routes to separate package for universal use in main code and in tests
		newRequest, err := http.NewRequest("POST", "/api/organizations", bytes.NewBuffer([]byte(body)))
		if err != nil {
			t.Fatalf("fail to create NewRequest: %v", err)
		}
		router.ServeHTTP(recorder, newRequest)

		if recorder.Code == http.StatusInternalServerError { // if == http.InternalError in case i won't find the way to dance around bad json load
			t.Error("Failed to create organization")
		} else if recorder.Code == http.StatusBadRequest {
			//TOCHECK: what is in error -> can i dance around that as not failed test but not ignore completely (i'm interested if in err i will have specific field that caused error)
		} else if recorder.Code == http.StatusOK {
			// responseData, err := io.ReadAll(recorder.Body)
			// if err != nil {
			// 	t.Fatalf("fail to read recorder's body: %v", err)
			// }
			var createdOrganization sqlc.Organization
			if err = json.Unmarshal(recorder.Body.Bytes(), createdOrganization); err != nil {
				t.Fatalf("fail to decode recorder's body to json: %v", err)
			}
			//TOFIX: encode new organization to json but without ID to compare with body
			json.Marshal()
			if body != string(responseData) {
				t.Errorf("send: %v\ngot:\v", body, responseData)
			}
		}

	})
}
