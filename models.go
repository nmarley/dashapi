package main

import (
	"fmt"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// TODO: Triggers

// Proposal represents a Proposal object.
type Proposal struct {
	// GovObj fields (hash, vote counts, etc.)
	Hash           string    `json:"hash" sql:",pk"`
	CollateralHash string    `json:"collateralHash"`
	CountYes       uint      `json:"countYes"`
	CountNo        uint      `json:"countNo"`
	CountAbstain   uint      `json:"countAbstain"`
	CreatedAt      time.Time `json:"createdAt"`

	// Proposal detail fields
	StartAt time.Time `json:"startAt"`
	EndAt   time.Time `json:"endAt"`
	Title   string    `json:"name"`
	URL     string    `json:"url"`
	Address string    `json:"address"`
	Amount  float64   `json:"amount"`
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

// IsValid returns whether the Proposal is valid
func (p *Proposal) IsValid() bool {
	zeroTime := time.Time{}
	return (p.Hash != "" &&
		p.CollateralHash != "" &&
		p.Title != "" &&
		p.URL != "" &&
		p.Address != "" &&
		p.Amount != 0 &&
		p.CreatedAt != zeroTime &&
		p.StartAt != zeroTime &&
		p.EndAt != zeroTime)
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

// alreadyHaveProposal returns whether or not this Proposal Hash has been
// recorded in the database.
func alreadyHaveProposal(db *pg.DB, prop *Proposal) bool {
	count, _ := db.Model((*Proposal)(nil)).
		Where("hash = ?", prop.Hash).
		Count()

	return count != 0
}

// upsertProposal does what it says
func upsertProposal(db *pg.DB, prop *Proposal) error {
	var err error
	if alreadyHaveProposal(db, prop) {
		// update proposal with fields in prop
		// TODO: debug logging
		// fmt.Printf("Already have %s, updating!\n", prop.Hash)
		err = db.Update(prop)
	} else {
		// Insert proposal
		// TODO: debug logging
		// fmt.Printf("Do NOT have %s, insert!\n", prop.Hash)
		err = db.Insert(prop)
	}
	if err != nil {
		return err
	}
	return nil
}
