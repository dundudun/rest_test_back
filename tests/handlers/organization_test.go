package handlers_tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dundudun/rest_test_back/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func FuzzCreateOrganization(f *testing.F) {
	cleanup := SetupDB()
	defer cleanup()

	//TOCHECK: can Fuzz generate nil?
	f.Add(63, "test_org_from_fuzzing_1", "100", "200", "10", "20", "200", "0")
	f.Add(62, "test_org_from_fuzzing_2", "100", "200", "10", "20", "200", "")
	//TODO: more cases for f.Add()

	// fieldSet sets what optional fields will be in request body (1 means will be present, 0 - won't)
	//   63 -> 111111 means all 6 optional fields will be present
	//   39 -> 111001 means optional fields {plasticLimit, glassLimit, biowasteLimit and producedBiowaste} will be present
	f.Fuzz(func(t *testing.T, fieldSet int, name, plasticLimit, glassLimit, biowasteLimit, producedPlastic, producedGlass, producedBiowaste string) {
		sendBody := `{"name":` + name
		i := 0
		for fieldSet %= 64; fieldSet != 0; fieldSet /= 2 {
			if fieldSet%2 == 1 {
				switch i {
				case 0:
					sendBody += `,"plastic_limit":"` + plasticLimit + `"`
				case 1:
					sendBody += `,"glass_limit":"` + glassLimit + `"`
				case 2:
					sendBody += `,"biowaste_limit":"` + biowasteLimit + `"`
				case 3:
					sendBody += `,"produced_plastic":"` + producedPlastic + `"`
				case 4:
					sendBody += `,"produced_glass":"` + producedGlass + `"`
				case 5:
					sendBody += `,"produced_biowaste":"` + producedBiowaste + `"`
				}
			}
			i++
		}
		sendBody += `}`

		gin.SetMode(gin.TestMode)
		router := gin.Default()
		router.POST("/api/organizations", handler.CreateOrganization)

		recorder := httptest.NewRecorder()
		//move routes to separate package for universal use in main code and in tests
		newRequest, err := http.NewRequest("POST", "/api/organizations", bytes.NewBuffer([]byte(sendBody)))
		if err != nil {
			t.Fatalf("INTERNAL TEST FUNCTION FAIL: fail to create NewRequest: %v", err)
		}
		router.ServeHTTP(recorder, newRequest)

		if recorder.Code == http.StatusInternalServerError { // if == http.InternalError in case i won't find the way to dance around bad json load
			t.Error("fail to create organization")
		} else if recorder.Code == http.StatusBadRequest {
			//TOCHECK: what is in error -> can i dance around that as not failed test but not ignore completely (i'm interested if in err i will have specific field that caused error)
		} else if recorder.Code == http.StatusCreated {
			var createdOrganization sqlc.Organization
			if err = json.Unmarshal(recorder.Body.Bytes(), &createdOrganization); err != nil {
				t.Fatalf("INTERNAL TEST FUNCTION FAIL: fail to decode recorder's body from json: %v", err)
			}
			//TOCHECK: if encode works as expected (especially what will be in optional fields if db returns nulls)
			got, err := json.Marshal(struct {
				Name             pgtype.Text `json:"name"`
				PlasticLimit     pgtype.Int4 `json:"plastic_limit,omitempty"`
				GlassLimit       pgtype.Int4 `json:"glass_limit,omitempty"`
				BiowasteLimit    pgtype.Int4 `json:"biowaste_limit,omitempty"`
				ProducedPlastic  pgtype.Int4 `json:"produced_plastic,omitempty"`
				ProducedGlass    pgtype.Int4 `json:"produced_glass,omitempty"`
				ProducedBiowaste pgtype.Int4 `json:"produced_biowaste,omitempty"`
			}{
				createdOrganization.Name,
				createdOrganization.PlasticLimit,
				createdOrganization.GlassLimit,
				createdOrganization.BiowasteLimit,
				createdOrganization.ProducedPlastic,
				createdOrganization.ProducedGlass,
				createdOrganization.ProducedBiowaste,
			})
			if err != nil {
				t.Fatalf("INTERNAL TEST FUNCTION FAIL: fail to encode created organization to json: %v", err)
			}
			if sendBody != string(got) {
				t.Errorf("send: %v\ngot:%v", sendBody, got)
			}
			// TOCHECK: what error handler will return if name isn't present in requestBody and move this code to according place (BadRequest should be i guess)
			// if !createdOrganization.Name.Valid {
			// 	t.Error("organization's name is required and can't be null")
			// }
		}
	})
}

func TestGetOrganization(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/api/organizations/:id", handler.GetOrganization)

	recorder := httptest.NewRecorder()
	newRequest, err := http.NewRequest("GET", "/api/organizations/"+id, nil)
	if err != nil {
		t.Fatalf("INTERNAL TEST FUNCTION FAIL: fail to create NewRequest: %v", err)
	}
	router.ServeHTTP(recorder, newRequest)

	if recorder.Code == http.StatusInternalServerError { // if == http.InternalError in case i won't find the way to dance around bad json load
		t.Error("fail to get the organization")
	} else if recorder.Code == http.StatusBadRequest { // && substring(recorder.Body.Bytes(), "invalid organization ID") == trues

	} else if recorder.Code == http.StatusOK {
		var returnedOrganization sqlc.Organization
		if err = json.Unmarshal(recorder.Body.Bytes(), &returnedOrganization); err != nil {
			t.Fatalf("INTERNAL TEST FUNCTION FAIL: fail to decode recorder's body from json: %v", err)
		}
		//TOCHECK: if encode works as expected (especially what will be in optional fields if db returns nulls)
		got, err := json.Marshal(struct {
			Name             pgtype.Text `json:"name"`
			PlasticLimit     pgtype.Int4 `json:"plastic_limit,omitempty"`
			GlassLimit       pgtype.Int4 `json:"glass_limit,omitempty"`
			BiowasteLimit    pgtype.Int4 `json:"biowaste_limit,omitempty"`
			ProducedPlastic  pgtype.Int4 `json:"produced_plastic,omitempty"`
			ProducedGlass    pgtype.Int4 `json:"produced_glass,omitempty"`
			ProducedBiowaste pgtype.Int4 `json:"produced_biowaste,omitempty"`
		}{
			returnedOrganization.Name,
			returnedOrganization.PlasticLimit,
			returnedOrganization.GlassLimit,
			returnedOrganization.BiowasteLimit,
			returnedOrganization.ProducedPlastic,
			returnedOrganization.ProducedGlass,
			returnedOrganization.ProducedBiowaste,
		})
		if err != nil {
			t.Fatalf("INTERNAL TEST FUNCTION FAIL: fail to encode created organization to json: %v", err)
		}
		if sendBody != string(got) {
			t.Errorf("send: %v\ngot:%v", sendBody, got)
		}
		// TOCHECK: what error handler will return if name isn't present in requestBody and move this code to according place (BadRequest should be i guess)
		// if !createdOrganization.Name.Valid {
		// 	t.Error("organization's name is required and can't be null")
		// }
	}
}
