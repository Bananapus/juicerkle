package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const treeDb = "trees.db"

var db *sql.DB

// The db stores address hex values with 0x prefixes (standard for go-ethereum's common.address).
// Roots and leaf hashes are stored as raw bytes for performance reasons.
func initDb() error {
	var err error

	db, err = sql.Open("sqlite3", treeDb)
	if err != nil {
		return err
	}

	// These are the inbox trees for each sucker. The leaves are read from the peer sucker.
	if _, err = db.Exec(`CREATE TABLE IF NOT EXISTS trees (
		chain_id TEXT,
		contract_address TEXT,
		token_address TEXT,	
		current_root TEXT,
		PRIMARY KEY (chain_id, contract_address, token_address)
	);`); err != nil {
		return err
	}

	// Leaves are associated with their inbox tree, not outbox trees.
	// Claimed is a boolean
	if _, err = db.Exec(`CREATE TABLE IF NOT EXISTS leaves (
		chain_id TEXT,
		contract_address TEXT,
		token_address TEXT,
		leaf_index TEXT,
		leaf_beneficiary TEXT,
		leaf_project_token_amount TEXT,
		leaf_terminal_token_amount TEXT,
		leaf_hash TEXT,
		is_claimed INTEGER,
		PRIMARY KEY (chain_id, contract_address, token_address, leaf_index)
	);`); err != nil {
		return err
	}

	if _, err = db.Exec(`CREATE INDEX IF NOT EXISTS idx_leaves_chain_contract_token 
		ON leaves (chain_id, contract_address, token_address);`); err != nil {
		return err
	}

	return nil
}
