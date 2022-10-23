package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	SenderField   string = "sender"
	ReceiverField string = "receiver"
	MessageField  string = "message"
)

type Connection struct {
	Client     *mongo.Client
	dbName     string
	collection string
}

type Querier interface {
	SendMessage(ctx context.Context, param MessageParam) (string, error)
	GetMessage(ctx context.Context, param Messanger) MessageParam
}

func NewConnection(ctx context.Context, connString string, dbName, collection string, opt ...*options.ServerAPIOptions) (*Connection, error) {
	clientOptions := options.Client().ApplyURI(connString)
	for _, v := range opt {
		if v != nil {
			clientOptions.SetServerAPIOptions(v)
		}
	}

	client, err := mongo.Connect(ctx, clientOptions)

	conn := Connection{
		Client:     client,
		dbName:     dbName,
		collection: collection,
	}

	return &conn, err
}

func (conn *Connection) Collection() *mongo.Collection {
	return conn.Client.Database(conn.dbName).Collection(conn.collection)
}

type Messanger struct {
	Sender   string
	Receiver string
}

type MessageParam struct {
	Sender   string `bson:"sender"`
	Receiver string `bson:"receiver"`
	Message  string `bson:"message"`
}

func (conn Connection) SendMessage(ctx context.Context, param MessageParam) (string, error) {
	coll := conn.Collection()
	if param.Sender == param.Receiver {
		return "", fmt.Errorf("cannot send message on self")
	}
	doc := bson.D{
		{SenderField, param.Sender},
		{ReceiverField, param.Receiver},
		{MessageField, param.Message},
	}

	res, err := coll.InsertOne(ctx, doc)
	if err != nil {
		return "", err
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("error convert uuid")
	}
	return id.String(), nil
}

func (conn Connection) GetMessage(ctx context.Context, param Messanger) MessageParam {
	var message MessageParam
	coll := conn.Collection()
	filter := interface{}(bson.D{{SenderField, param.Sender}, {ReceiverField, param.Receiver}})
	res := coll.FindOne(ctx, filter)
	res.Decode(&message)
	return message
}

func (conn *Connection) Disconnect(ctx context.Context) error {
	return conn.Client.Disconnect(ctx)
}
