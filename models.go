package main

import (
	"fmt"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// TODO: Models for Proposal, Trigger...

// Proposal represents a Proposal object.
type Proposal struct {
	StartAt time.Time `json:"startAt"`
	EndAt   time.Time `json:"endAt"`
	Title   string    `json:"name"`
	URL     string    `json:"url"`
	Address string    `json:"address"`
	Amount  float64   `json:"amount"`

	CreatedAt time.Time `json:"createdAt"`
}

// String implements the Stringer interface for Proposal
func (p Proposal) String() string {
	return fmt.Sprintf(
		"Proposal(Title: %s, URL: %s, Address: %s, Amount %v, StartAt: %v, EndAt: %v)",
		p.Title,
		p.URL,
		p.Address,
		p.Amount,
		p.StartAt.UTC(),
		p.EndAt.UTC(),
	)
}

// createSchema makes the DB tables if they don't exist
func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{
		// Add models here...
		(*Proposal)(nil),
	} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// getCurrentVotesOnly returns a list of the latest votes for each address
// func getCurrentVotesOnly(db *pg.DB) ([]Vote, error) {
// 	countingVotes := []Vote{}
//
// 	query := `
// 	select distinct t.address
// 	     , t.message
// 	     , t.signature
// 	     , t.created_at
// 	  from votes t
// 	 inner join (
// 	       select address
// 	            , max(created_at) as maxdate
// 	         from votes
// 	        group by address
// 	       ) tm
// 	    on t.address = tm.address
// 	   and t.created_at = tm.maxdate
// 	`
//
// 	_, err := db.Query(&countingVotes, query)
// 	return countingVotes, err
// }

// getAllProposals returns a list of all proposals in the database
func getAllProposals(db *pg.DB) ([]Proposal, error) {
	proposals := []Proposal{}

	err := db.Model(&proposals).Select()
	return proposals, err
}
