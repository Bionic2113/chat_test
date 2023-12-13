package scylla

import (
	"chat_v2/models"
	"fmt"
	"strconv"

	"github.com/gocql/gocql"
	"go.uber.org/zap"
)

func FindAll(session *gocql.Session, logger *zap.Logger) []models.User {
	logger.Info("Displaying Results:")

	q := session.Query("SELECT * FROM test_users") //"SELECT first_name,last_name,address,picture_location FROM mutant_data")
	var user_id, username string
	it := q.Iter()
	defer func() {
		if err := it.Close(); err != nil {
			logger.Warn("select my_keyspace.test_users", zap.Error(err))
		}
	}()
	users := []models.User{}
	var index int
	for it.Scan(&user_id, &username) {
		users = append(users, models.User{User_id: user_id, Username: username})
		logger.Info("index: " + strconv.Itoa(index) + "\t" + user_id + " " + username)
		index++
	}
	return users
}

func FindFriends(user_id string, session *gocql.Session, logger *zap.Logger) []models.User {
	logger.Info("Displaying Results:")

	q := session.Query("SELECT * FROM test_users WHERE user_id <> ?", user_id) //"SELECT first_name,last_name,address,picture_location FROM mutant_data")
	var username string
	it := q.Iter()
	defer func() {
		if err := it.Close(); err != nil {
			logger.Warn("select my_keyspace.test_users", zap.Error(err))
		}
	}()
	users := []models.User{}
	for it.Scan(&user_id, &username) {
		users = append(users, models.User{User_id: user_id, Username: username})
		logger.Info("\t" + user_id + " " + username)
	}
	return users
}

func FindCurrentUser(user_id string, session *gocql.Session, logger *zap.Logger) (models.User, error) {
	logger.Info("Displaying Results:")

	q := session.Query("SELECT * FROM test_users WHERE user_id = ?", user_id) //"SELECT first_name,last_name,address,picture_location FROM mutant_data")
	user := models.User{}
	if err := q.Scan(&user.User_id, &user.Username); err != nil {
		// на данный момент странно, тк если в куках нет его user_id,
		// то надо было уже создать, а не искать.
		// Мб в будущем прост пригодится
		//return CreateUser(user_id, session, logger)
		return user, fmt.Errorf("user with user_id: %s not found. Err: %s", user_id, err)
	}
	return user, nil
}

func CreateUser(user_id, username string, session *gocql.Session, logger *zap.Logger) error {
	return session.Query("INSERT INTO test_users (user_id, username) VALUES (?, ?)", user_id, username).Exec()
}

func FindAllMessages(sender_id, reciever_id string, session *gocql.Session, logger *zap.Logger) []models.Message {
	logger.Info("Displaying Messages Results:")

	q := session.Query("SELECT * FROM test_message WHERE (sender, reciever) IN ((?,?), (?,?)) ALLOW FILTERING", sender_id, reciever_id, reciever_id, sender_id)
	it := q.Iter()
	defer func() {
		if err := it.Close(); err != nil {
			logger.Warn("select my_keyspace.test_message", zap.Error(err))
		}
	}()
  logger.Info("num rows = " + fmt.Sprint(it.NumRows()))
	messages := make([]models.Message, 0, it.NumRows())
	var message models.Message
	var index int
	for it.Scan(&message.Message_id, &message.Sender_id, &message.Reciever_id, &message.Message) {
		messages = append(messages, message)
		logger.Info("index: " + strconv.Itoa(index) + "\t" + message.Sender_id + " " + message.Reciever_id + " " + message.Message)
		index++
	}
	return messages
}

func CreateMessage(sender_id, reciever_id, message string, session *gocql.Session, logger *zap.Logger) error {
	return session.Query("insert into test_message (message_id, sender, reciever, message) values (uuid(), ?, ?, ?)", sender_id, reciever_id, message).Exec()
}
