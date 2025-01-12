package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client     *mongo.Client
	collection *mongo.Collection
}

// NewMongoDB initializes a new MongoDB instance.
func NewMongoDB(uri, dbName, collectionName string) (*MongoDB, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return &MongoDB{
		client:     client,
		collection: client.Database(dbName).Collection(collectionName),
	}, nil
}

// SaveArticle saves an article to the database.
func (m *MongoDB) SaveArticle(ctx context.Context, article Article) error {
	_, err := m.collection.InsertOne(ctx, article)
	return err
}

// GetArticlesByUser retrieves articles by user ID.
func (m *MongoDB) GetArticlesByUser(ctx context.Context, userID string) ([]Article, error) {
	cursor, err := m.collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var articles []Article
	if err = cursor.All(ctx, &articles); err != nil {
		return nil, err
	}
	return articles, nil
}

// GetArticle retrieves a single article by ID.
func (m *MongoDB) GetArticle(ctx context.Context, articleID string) (*Article, error) {
	var article Article
	err := m.collection.FindOne(ctx, bson.M{"id": articleID}).Decode(&article)
	if err != nil {
		return nil, err
	}
	return &article, nil
}

// UpdateArticle updates an article by ID.
func (m *MongoDB) UpdateArticle(ctx context.Context, articleID string, updates map[string]interface{}) error {
	_, err := m.collection.UpdateOne(ctx, bson.M{"id": articleID}, bson.M{"$set": updates})
	return err
}

// DeleteArticle deletes an article by ID.
func (m *MongoDB) DeleteArticle(ctx context.Context, articleID string) error {
	_, err := m.collection.DeleteOne(ctx, bson.M{"id": articleID})
	return err
}

func (m *MongoDB) Ping(ctx context.Context) error {
	return m.client.Ping(ctx, nil)
}
