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

	checkpoint = "checkpoint"
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

func SaveBlock(hash string, data []byte) {
	// fmt.Printf("Saving Block %s\nData: %b\n", hash, data);
	err := DB().Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blocksBucket));
		err := bucket.Put([]byte(hash), data); // Put KEY & VALUE
		return err;
	})
	utils.HandleErr(err);
}

func SaveBlockchain(data []byte) {
	err := DB().Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket));
		err := bucket.Put([]byte(checkpoint), data);
		return err;
	})
	utils.HandleErr(err);

}

func Checkpoint() []byte {
	var data []byte
	DB().View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket));
		data = bucket.Get([]byte(checkpoint))
		return nil
	})

	return data;
}