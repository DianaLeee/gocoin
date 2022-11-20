package db

import (
	"github.com/DianaLeee/gocoin/utils"
	"github.com/boltdb/bolt"
)

// bucket (Bolt) = table (SQL)
const (
	dbName = "blockchain.db"
	dataBucket = "data"
	blocksBucket = "blocks"
)

var db *bolt.DB;

func DB() *bolt.DB {
	if db == nil {
		// init db
		dbPointer, err :=bolt.Open(dbName, 0600, nil);
		db = dbPointer;
		utils.HandleErr(err);

		err = db.Update(func(t *bolt.Tx) error {
			_, err := t.CreateBucketIfNotExists([]byte(dataBucket));
			utils.HandleErr(err);
			_, err = t.CreateBucketIfNotExists([]byte(blocksBucket));

			return err
		})
		utils.HandleErr(err);


	}
	return db;
}
