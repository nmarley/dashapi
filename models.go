package main

import (
	"fmt"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// MNReport represents a Masternode report.
type MNReport struct {
	// somehow identify a unique report via MN ID and timestamp
	ID        string    `json:"hash" sql:",pk" sql:",notnull"`
	Hash      string    `json:"hash" sql:",pk" sql:",notnull"`
	CreatedAt time.Time `json:"createdAt" sql:",notnull"`

	// Proposal detail fields
	MNCount uint16 `json:"mnCount" sql:",notnull"`

	SBCycle uint16 `json:"sbCycle" sql:",notnull"`
	LastSB  int    `json:"lastSB" sql:",notnull"`
	NextSB  int    `json:"lastSB" sql:",notnull"`
	// in Satoshis
	NextSBBudget uint64 `json:"nextSBBudget" sql:",notnull"`
}

// GetLatestReport returns the latest MN report
func GetLatestReport(db *pg.DB) MNReport {
	report := MNReport{}

	query := `
    select r.*
      from mn_reports r
     inner join (
           select id
                , max(created_at) as maxdate
             from mn_reports
            group by id
           ) latest
        on r.id = latest.id
       and r.created_at = latest.maxdate
    `

	_, err := db.Query(&report, query)
	if err != nil {
		// Log it... & return an empty report
		fmt.Println("error: ", err.Error())
		// return MNReport{}
	}

	return report
}

// Proposal represents a Proposal object.
type Proposal struct {
	// GovObj fields (hash, vote counts, etc.)
	Hash           string    `json:"hash" sql:",pk" sql:",notnull"`
	CollateralHash string    `json:"collateralHash" sql:",notnull"`
	CountYes       uint      `json:"countYes" sql:",notnull"`
	CountNo        uint      `json:"countNo" sql:",notnull"`
	CountAbstain   uint      `json:"countAbstain" sql:",notnull"`
	CreatedAt      time.Time `json:"createdAt" sql:",notnull"`

	// Proposal detail fields
	StartAt time.Time `json:"startAt" sql:",notnull"`
	EndAt   time.Time `json:"endAt" sql:",notnull"`
	Title   string    `json:"name" sql:",notnull"`
	URL     string    `json:"url" sql:",notnull"`
	Address string    `json:"address" sql:",notnull"`
	Amount  float64   `json:"amount" sql:",notnull"`
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

// currentProposals returns a list of current Proposals.
func currentProposals(db *pg.DB) ([]Proposal, error) {
	proposals := []Proposal{}

	query := `
	select p.*
	  from proposals p
	 where p.start_at <= now()
	   and p.end_at >= now()
     order by (p.count_yes - p.count_no) desc
	`

	_, err := db.Query(&proposals, query)
	return proposals, err
}
