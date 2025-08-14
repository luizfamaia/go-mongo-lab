//go:build integration

package integrations

import (
	"context"
	"os"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID    string `bson:"id,omitempty"`
	Name  string `bson:"name"`
	Email string `bson:"email"`
}

func mongoURI() string {
	if uri := os.Getenv("MONGO_URI"); uri != "" {
		return uri
	}
	return "mongodb://localhost:27017"
}

func TestUserCRUD(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI()))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = client.Disconnect(context.Background()) })

	db := client.Database("testdb_integration")
	coll := db.Collection("users_it")

	// limpa a coleção antes do teste
	_ = coll.Drop(ctx)

	// CREATE
	u := User{ID: "it-1", Name: "Luiz IT", Email: "luiz.it@example.com"}
	if _, err := coll.InsertOne(ctx, u); err != nil {
		t.Fatalf("insert: %v", err)
	}

	// READ
	var got User
	if err := coll.FindOne(ctx, bson.M{"id": u.ID}).Decode(&got); err != nil {
		t.Fatalf("find: %v", err)
	}
	if got.Name != u.Name {
		t.Fatalf("name mismatch: got=%s want=%s", got.Name, u.Name)
	}

	// UPDATE
	if _, err := coll.UpdateOne(ctx, bson.M{"id": u.ID}, bson.M{"$set": bson.M{"name": "Luiz Updated"}}); err != nil {
		t.Fatalf("update: %v", err)
	}
	if err := coll.FindOne(ctx, bson.M{"id": u.ID}).Decode(&got); err != nil {
		t.Fatalf("verify update: %v", err)
	}
	if got.Name != "Luiz Updated" {
		t.Fatalf("verify update name mismatch: got=%s", got.Name)
	}

	// DELETE
	if _, err := coll.DeleteOne(ctx, bson.M{"id": u.ID}); err != nil {
		t.Fatalf("delete: %v", err)
	}
	err = coll.FindOne(ctx, bson.M{"id": u.ID}).Decode(&got)
	if err == nil || err != mongo.ErrNoDocuments {
		t.Fatalf("expected ErrNoDocuments after delete, got=%v", err)
	}
}
