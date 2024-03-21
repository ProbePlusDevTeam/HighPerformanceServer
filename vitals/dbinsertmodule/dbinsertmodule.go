package dbinsertmodule

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	streamDocuments      = make(map[string][]interface{})
	streamDocumentsMutex = &sync.Mutex{}
)

// BulkInsert represents a MongoDB bulk insert operation.
type BulkInsert struct {
	collection *mongo.Collection
	documents  chan []interface{}
}

// NewBulkInsert initializes a new BulkInsert instance.
func NewBulkInsert(client *mongo.Client, dbName, collectionName string) (*BulkInsert, error) {
	collection := client.Database(dbName).Collection(collectionName)
	return &BulkInsert{
		collection: collection,
		documents:  make(chan []interface{}),
	}, nil
}

// Execute executes the bulk insert operation with documents received from the channel.
func (bi *BulkInsert) Execute(ctx context.Context, documents []interface{}) error {
	if len(documents) == 0 {
		return nil // No documents to insert
	}

	_, err := bi.collection.InsertMany(ctx, documents)
	if err != nil {
		return err
	}
	return nil
}

// Cancel closes the documents channel.
func (bi *BulkInsert) Cancel() {
	close(bi.documents)
}

// StartInserting starts inserting documents into MongoDB collection
// every specified duration.
func StartInserting(client *mongo.Client, groupMap map[string]context.Context, dbName, collectionName string, duration time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	// Create a new BulkInsert instance
	bulkInsert, err := NewBulkInsert(client, dbName, collectionName)
	if err != nil {
		log.Fatal(err)
	}
	defer bulkInsert.Cancel()

	// Start a goroutine to collect and insert documents periodically

	ticker := time.NewTicker(duration)
	defer ticker.Stop()
	for {
		select {
		case <-groupMap[collectionName].Done():
			fmt.Println("exiting ")
			return
		case <-ticker.C:
			streamDocumentsMutex.Lock()
			if len(streamDocuments[collectionName]) == 0 {
				delete(streamDocuments, collectionName)
				delete(groupMap, collectionName)
				fmt.Println("exiting ", streamDocuments, groupMap, collectionName)
				streamDocumentsMutex.Unlock()
				return
			}
			startTime := time.Now()
			if err := bulkInsert.Execute(groupMap[collectionName], streamDocuments[collectionName]); err != nil {
				log.Printf("Error inserting documents: %v\n", err)
			} else {
				log.Printf("Inserted documents %d number of go routines %d groupid %s\n", len(streamDocuments[collectionName]), runtime.NumGoroutine(), collectionName)
			}
			fmt.Println(time.Since(startTime))
			streamDocuments[collectionName] = nil
			streamDocumentsMutex.Unlock()
		}
	}

	//// Wait until the context is done
	//<-ctx.Done()
	//
	//// Signal to stop inserting
	//close(stopInserting)
}

func AddData(data bson.D, groupId string) {
	streamDocumentsMutex.Lock()
	streamDocuments[groupId] = append(streamDocuments[groupId], data)
	streamDocumentsMutex.Unlock()

}
