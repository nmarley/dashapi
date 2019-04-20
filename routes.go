package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

// JWTSecretKey is used to verify the JWT was signed w/the same, used for
// authorization.
// See also: https://jwt.io/#debugger
var JWTSecretKey []byte

// DashNetwork is used for validating the address network byte
var DashNetwork string

// ServeHTTP passes requests thru to the router.
func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// routes defines the routes the server will handle
func (s *server) routes() {
	// health check
	s.router.HandleFunc("/health", s.handleHealthCheck())

	// TODO: remove these and add Data recording + retrieval routes here...

	// route to record incoming proposals
	s.router.HandleFunc("/proposal", s.handleProposal())

	// audit routes
	// s.router.HandleFunc("/validVotes", isAuthorized(s.handleValidVotes()))
	s.router.HandleFunc("/allProposals", isAuthorized(s.handleAllProposals()))

	// catch-all (404)
	s.router.PathPrefix("/").Handler(s.handleIndex())
}

// isAuthorized is used to wrap handlers that need authz
func isAuthorized(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bearerToken, ok := r.Header["Authorization"]
		if !ok {
			writeError(http.StatusUnauthorized, w, r)
			return
		}

		// strip the "Bearer " from the beginning
		actualTokenStr := strings.TrimPrefix(bearerToken[0], "Bearer ")

		// Parse and validate token from request header
		token, err := jwt.Parse(actualTokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return "invalid signing method", nil
			}
			return JWTSecretKey, nil
		})
		if err != nil {
			writeError(http.StatusUnauthorized, w, r)
			return
		}

		// JWT is valid, pass the request thru to protected route
		if token.Valid {
			f(w, r)
		}
	}
}

// handleProposal handles the proposal route
func (s *server) handleProposal() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse proposal body
		var p Proposal
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			writeError(http.StatusBadRequest, w, r)
			return
		}

		// Basic input validation
		if !p.IsValid() {
			writeError(http.StatusBadRequest, w, r)
			return
		}

		// Upsert proposal
		err = upsertProposal(s.db, &p)
		if err != nil {
			// TODO: debug logging
			// fmt.Println("err =", err.Error())
			writeError(http.StatusInternalServerError, w, r)
			return
		}

		// Return response
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_ = json.NewEncoder(w).Encode(JSONResult{
			Status:  http.StatusCreated,
			Message: "Proposal Recorded",
		})
	}
}

// handleVoteClosed handles the vote route once voting is Closed
// func (s *server) handleVoteClosed() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		// Return response
// 		w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 		_ = json.NewEncoder(w).Encode(JSONResult{
// 			Status:  http.StatusForbidden,
// 			Message: "Voting Closed",
// 		})
// 	}
// }

// handleValidVotes is the route for vote tallying, and returns only most
// current vote per MN collateral address.
// TODO: consider pagination if this gets too big.
// func (s *server) handleValidVotes() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		votes, err := getCurrentVotesOnly(s.db)
// 		if err != nil {
// 			writeError(http.StatusInternalServerError, w, r)
// 			return
// 		}
//
// 		w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 		err = json.NewEncoder(w).Encode(&votes)
// 		if err != nil {
// 			writeError(http.StatusInternalServerError, w, r)
// 			return
// 		}
// 	}
// }

// handleAllProposals is the route for listing all proposals.
// TODO: consider pagination if this gets too big.
func (s *server) handleAllProposals() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		proposals, err := getAllProposals(s.db)
		if err != nil {
			writeError(http.StatusInternalServerError, w, r)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		err = json.NewEncoder(w).Encode(&proposals)
		if err != nil {
			writeError(http.StatusInternalServerError, w, r)
			return
		}
	}
}

// handleHealthCheck handles the health check route, an unauthenticated route
// needed for load balancers to know this service is still "healthy".
func (s *server) handleHealthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_ = json.NewEncoder(w).Encode(JSONResult{
			Status:  http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		})
	}
}

// JSONErrorMessage represents the JSON structure of an error message to be
// returned.
type JSONErrorMessage struct {
	Status int    `json:"status"`
	URL    string `json:"url"`
	Error  string `json:"error"`
}

// JSONResult represents the JSON structure of the success message to be
// returned.
type JSONResult struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// writeError returns a generic JSON error blob.
func writeError(errorCode int, w http.ResponseWriter, r *http.Request) {
	msg := JSONErrorMessage{
		Status: errorCode,
		URL:    r.URL.Path,
		Error:  http.StatusText(errorCode),
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(errorCode)
	_ = json.NewEncoder(w).Encode(msg)
	return
}

// handleIndex is catch-all route handler.
func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		writeError(http.StatusNotFound, w, r)
		return
	}
}

func init() {
	JWTSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
}
