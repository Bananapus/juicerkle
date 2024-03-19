package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const treeDb = "trees.db"

var db *sql.DB

func initDb() error {
	var err error

	db, err = sql.Open("sqlite3", treeDb)
	if err != nil {
		return err
	}

	if _, err = db.Exec(`CREATE TABLE IF NOT EXISTS suckers (
		chain_id INTEGER,
		contract_address TEXT,
		peer_chain_id INTEGER,
		peer_contract_address TEXT,
		PRIMARY KEY (chain_id, contract_address)
	);`); err != nil {
		return err
	}

	if _, err = db.Exec(`CREATE TABLE IF NOT EXISTS trees (
		chain_id INTEGER,
		contract_address TEXT,
		token_address TEXT,	
		current_root TEXT,
		count INTEGER,
		PRIMARY KEY (chain_id, contract_address, token_address)
		FOREIGN KEY (chain_id, contract_address) REFERENCES suckers (chain_id, contract_address)
	);`); err != nil {
		return err
	}

	if _, err = db.Exec(`CREATE TABLE IF NOT EXISTS leaves (
		chain_id INTEGER,
		contract_address TEXT,
		token_address TEXT,
		leaf_index INTEGER,
		leaf_hash TEXT,
		PRIMARY KEY (chain_id, contract_address, token_address, leaf_index)
		FOREIGN KEY (chain_id, contract_address, token_address) REFERENCES tree (chain_id, contract_address, token_address)
	);`); err != nil {
		return err
	}

	if _, err = db.Exec(`CREATE INDEX IF NOT EXISTS idx_leaves_chain_contract_token 
		ON leaves (chain_id, contract_address, token_address
	);`); err != nil {
		return err
	}

	return nil
}
