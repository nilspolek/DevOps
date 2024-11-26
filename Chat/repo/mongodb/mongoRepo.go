package mongodb

import (
	"context"
	"os"

	"github.com/google/uuid"
	messageservice "github.com/nilspolek/DevOps/Chat/direct_message_service"
	groupmessageservice "github.com/nilspolek/DevOps/Chat/group_message_service"
	groupservice "github.com/nilspolek/DevOps/Chat/group_service"
	reactionservice "github.com/nilspolek/DevOps/Chat/reaction_service"
	"github.com/nilspolek/DevOps/Chat/repo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongorepo struct {
	DMCollection           mongo.Collection
	GroupMessageCollection mongo.Collection
	GroupCollection        mongo.Collection
}

func New() (repo.Repo, error) {
	const (
		DB_NAME                   = "MessagingDB"
		DIRECT_MESSAGE_COLLECTION = "direct_messages"
		GROUP_COLLECTION          = "groups"
		GROUP_MESSAGE_COLLECTION  = "group_messages"
	)

	var (
		mr  mongorepo
		uri string
		ctx = context.Background()
		err error
	)

	if uri = os.Getenv("MONGO_URI"); uri == "" {
		uri = "mongodb://localhost:27017"
	}

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return mr, err
	}

	mr.DMCollection = *client.Database(DB_NAME).Collection(DIRECT_MESSAGE_COLLECTION)
	mr.GroupCollection = *client.Database(DB_NAME).Collection(GROUP_COLLECTION)
	mr.GroupMessageCollection = *client.Database(DB_NAME).Collection(GROUP_MESSAGE_COLLECTION)

	return mr, nil
}

func (mr mongorepo) GetDirectMessages(userID messageservice.ID) ([]messageservice.Message, error) {
	var (
		messages []messageservice.Message
		ctx      = context.Background()
		filter   = bson.M{
			"$or": []bson.M{
				{"senderid": userID},
				{"receiverid": userID},
			},
		}
	)
	cursor, err := mr.DMCollection.Find(ctx, filter)
	if err != nil {
		return messages, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var result messageservice.Message
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		messages = append(messages, result)
	}
	return messages, err
}

func (mr mongorepo) SendDirectMessage(msg messageservice.Message) error {
	var (
		ctx = context.Background()
	)
	_, err := mr.DMCollection.InsertOne(ctx, msg)
	return err
}

func (mr mongorepo) ReplaceDirecMessage(messageID messageservice.ID, msg messageservice.Message) error {
	var (
		ctx = context.Background()
	)
	_, err := mr.DMCollection.ReplaceOne(ctx, bson.M{"id": messageID}, msg)
	return err
}

func (mr mongorepo) DeleteDirectMessage(messageID messageservice.ID) error {
	var (
		ctx = context.Background()
	)
	_, err := mr.DMCollection.DeleteOne(ctx, bson.M{"id": messageID})
	return err
}

func (mr mongorepo) GetGroupMessages(groupID groupmessageservice.ID) ([]groupmessageservice.Message, error) {
	var (
		messages []groupmessageservice.Message
		ctx      = context.Background()
	)
	cursor, err := mr.GroupMessageCollection.Find(ctx, bson.M{"id": groupID})
	if err != nil {
		return messages, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var result groupmessageservice.Message
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		messages = append(messages, result)
	}
	return messages, err
}

func (mr mongorepo) SendMessageToGroup(groupID groupmessageservice.ID, msg groupmessageservice.Message) error {
	var (
		ctx = context.Background()
	)
	msg.GroupId = groupID
	_, err := mr.GroupMessageCollection.InsertOne(ctx, msg)
	return err
}

func (mr mongorepo) ReplaceGroupMessage(messageID groupmessageservice.ID, msg groupmessageservice.Message) error {
	var (
		ctx = context.Background()
	)
	_, err := mr.GroupMessageCollection.ReplaceOne(ctx, bson.M{"id": messageID}, msg)
	return err
}

func (mr mongorepo) DeleteGroupMessage(messageID groupmessageservice.ID) error {
	var (
		ctx = context.Background()
	)
	_, err := mr.GroupMessageCollection.DeleteOne(ctx, bson.M{"id": messageID})
	return err
}
func (mr mongorepo) GetAllGroups(userId groupservice.ID) ([]groupservice.Group, error) {
	var (
		ctx    = context.Background()
		groups []groupservice.Group
	)
	filter := bson.M{"members": bson.M{"$in": []groupservice.ID{userId}}}

	cursor, err := mr.GroupCollection.Find(ctx, filter)
	if err != nil {
		return groups, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var result groupservice.Group
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		groups = append(groups, result)
	}
	return groups, err
}
func (mr mongorepo) CreateGroup(group groupservice.Group) (groupservice.ID, error) {
	var (
		ctx = context.Background()
	)
	group.Id = groupservice.ID(uuid.New())
	_, err := mr.GroupCollection.InsertOne(ctx, group)
	return group.Id, err
}

func (mr mongorepo) EditGroup(group groupservice.Group, groupId groupservice.ID) error {
	ctx := context.Background()

	filter := bson.M{"id": groupId}

	update := bson.M{
		"$set": bson.M{
			"title":    group.Title,
			"imageUrl": group.ImageUrl,
		},
	}

	_, err := mr.GroupCollection.UpdateOne(ctx, filter, update)
	return err
}

func (mr mongorepo) DeleteGroup(groupID groupservice.ID) error {
	ctx := context.Background()

	filter := bson.M{"id": groupID}
	_, err := mr.GroupCollection.DeleteOne(ctx, filter)
	return err
}

func (mr mongorepo) AddUserToGroup(groupId, userID groupservice.ID) error {
	ctx := context.Background()

	filter := bson.M{"id": groupId}
	update := bson.M{
		"$push": bson.M{
			"members": userID,
		},
	}
	_, err := mr.GroupCollection.UpdateOne(ctx, filter, update)

	return err
}

func (mr mongorepo) RemoveUserFromGroup(groupId, userID groupservice.ID) error {
	ctx := context.Background()

	filter := bson.M{"_id": groupId}
	update := bson.M{
		"$pull": bson.M{
			"members": userID,
		},
	}

	_, err := mr.GroupCollection.UpdateOne(ctx, filter, update)
	return err
}

func (mr mongorepo) AddReactionToDM(messageID, userID reactionservice.ID, reaction messageservice.Reaction) error {
	ctx := context.Background()

	filter := bson.M{"id": messageID}
	update := bson.M{
		"$push": bson.M{
			"reactions": reaction,
		},
	}
	_, err := mr.DMCollection.UpdateOne(ctx, filter, update)
	return err
}

func (mr mongorepo) ChangeReactionToDM(messageID, userID reactionservice.ID, reaction messageservice.Reaction) error {
	ctx := context.Background()

	filter := bson.M{"id": messageID}

	update := bson.M{
		"$set": bson.M{
			"reactions.$[elem].reaction": reaction.Reaction,
		},
	}

	arrayFilters := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{
			bson.M{"elem.sender": userID},
		},
	})

	_, err := mr.DMCollection.UpdateOne(ctx, filter, update, arrayFilters)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			update = bson.M{
				"$push": bson.M{
					"reactions": reaction,
				},
			}

			_, err = mr.DMCollection.UpdateOne(ctx, filter, update)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func (mr mongorepo) RemoveReactionFromDM(messageID, userID reactionservice.ID) error {
	ctx := context.Background()
	filter := bson.M{"id": messageID}
	update := bson.M{
		"$pull": bson.M{
			"reactions": bson.M{"sender": userID},
		},
	}

	_, err := mr.DMCollection.UpdateOne(ctx, filter, update)
	return err
}

func (mr mongorepo) AddReactionToGroup(messageID, userID groupmessageservice.ID, reaction groupmessageservice.Reaction) error {
	ctx := context.Background()
	filter := bson.M{"id": messageID}
	update := bson.M{
		"$push": bson.M{
			"reactions": reaction,
		},
	}

	_, err := mr.GroupMessageCollection.UpdateOne(ctx, filter, update)
	return err
}

func (mr mongorepo) ChangeReactionToGroup(messageID, reaction groupmessageservice.Reaction) error {
	ctx := context.Background()
	var err error
	filter := bson.M{"id": messageID}

	update := bson.M{
		"$set": bson.M{
			"reactions.$[elem].reaction": reaction,
		},
	}

	arrayFilters := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{
			bson.M{"elem.sender": reaction.Sender.Id},
		},
	})

	result, err := mr.GroupMessageCollection.UpdateOne(ctx, filter, update, arrayFilters)

	if result.MatchedCount == 0 {

		update = bson.M{
			"$push": bson.M{
				"reactions": reaction,
			},
		}
		_, err = mr.GroupMessageCollection.UpdateOne(ctx, filter, update)
	}
	return err
}

func (mr mongorepo) RemoveReactionFromGroup(messageID, userID groupmessageservice.ID) error {
	ctx := context.Background()
	filter := bson.M{"id": messageID}

	update := bson.M{
		"$pull": bson.M{
			"reactions": bson.M{"sender": userID},
		},
	}

	_, err := mr.GroupMessageCollection.UpdateOne(ctx, filter, update)
	return err
}
