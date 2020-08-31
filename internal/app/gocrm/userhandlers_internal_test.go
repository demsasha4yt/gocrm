package gocrm_test

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/demsasha4yt/gocrm.git/internal/app/store/sqlstore"
// )

// func TestUserHandler_HandleUsersCreate(t *testing.T) {
// 	db, teardown := sqlstore.TestDB(t)
// 	defer teardown("users")
// 	s := NewServer(db)

// 	testcases := []struct {
// 		name         string
// 		payload      interface{}
// 		expectedCode int
// 	}{
// 		{
// 			name: "valid",
// 			payload: map[string]string{
// 				"login":    "login",
// 				"email":    "username@example.org",
// 				"password": "password",
// 			},
// 			expectedCode: http.StatusCreated,
// 		},
// 		{
// 			name: "dublicate login",
// 			payload: map[string]string{
// 				"login":    "login",
// 				"email":    "username@example.org",
// 				"password": "password",
// 			},
// 			expectedCode: http.StatusUnprocessableEntity,
// 		},
// 		{
// 			name: "dublicate email",
// 			payload: map[string]string{
// 				"login":    "login2",
// 				"email":    "username@example.org",
// 				"password": "password",
// 			},
// 			expectedCode: http.StatusUnprocessableEntity,
// 		},
// 		{
// 			name:         "Invalid payload",
// 			payload:      "invalid",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name: "no login or email",
// 			payload: map[string]string{
// 				"password": "password",
// 			},
// 			expectedCode: http.StatusUnprocessableEntity,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			rec = httptest.NewRecorder()
// 			b := &bytes.Buffer{}
// 			json.NewEncoder(b).Encode(tc.payload)
// 			req, _ := http.NewRequest(http.MethodPost, "/users", b)
// 			s.ServeHTTP(rec, req)
// 		})
// 	}
// }

// func TestUserHandler_HandleSessionCreate(t *testing.T) {

// }
