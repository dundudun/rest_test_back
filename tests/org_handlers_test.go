package handlers_tests

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"

	"github.com/dundudun/rest_test_back/internal/handlers"
)

var h handlers.Handler

func SetupDB() func() {
	DATABASE_URL := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)

	h.Ctx = context.Background()
	var err error
	if h.Db, err = pgx.Connect(h.Ctx, DATABASE_URL); err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	return func() { h.Db.Close(h.Ctx) }
}

func FuzzCreateOrg(f *testing.F) {
	cleanup := SetupDB()
	defer cleanup()

	f.Add("test_org_from_fuzzing_1", "100", "200", "10", "20", "200", "0")
	f.Add("test_org_from_fuzzing_2", "100", "200", "10", "20", "200")
	f.Fuzz(func(t *testing.T, name, plasticLimit, glassLimit, biowasteLimit, producedPlastic, producedGlass, producedBiowaste string) {
		// data->json->reqsBody
		// h.CreateOrganization(req)
		// check statusOk
		// row select ... where ... ->parse to json
		// compare jsons - send and got
		r := gin.Default()
		r.POST("/api/organizations", h.CreateOrganization)

		recorder := httptest.NewRecorder()
		//move routes to separate package for universal use in main code and in tests
		newRequest, err := http.NewRequest("POST", "/api/organizations", bytes.NewBuffer([]byte(body)))
		r.ServeHTTP(recorder, newRequest)

		if recorder.Code != http.StatusOK { // if == http.InternalError in case i won't find the way to dance around bad json load
			t.Errorf("Expected status code 200 - OK, but got %d", recorder.Code)
		} else if recorder.Code == http.StatusBadRequest {
			//TOCHECK: what is in error -> can i dance around that as not failed test but not ignore completely
		}

		responseData, _ := io.ReadAll(recorder.Body)
	})

}

//Command Timeout=0;
