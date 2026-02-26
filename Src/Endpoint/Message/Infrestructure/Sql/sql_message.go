package sql

import (
	config "chat/Src/Core"
	entities "chat/Src/Endpoint/Message/Domain/Entities"
	"fmt"
)

type Mysql struct {
	config *config.ConnMySQL
}

func NewMySQL() (*Mysql, error) {
	conn := config.GetDBPool()
	if conn.Err != "" {
		return nil, fmt.Errorf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &Mysql{config: conn}, nil
}

func (m *Mysql) GetMessagesByUserId(senderId int, receiveId int) ([]*entities.Message, error) {
	query := `SELECT idMessage, senderId, receiveId, content, timeMessage
				FROM messages
				WHERE (senderId = ? AND receiveId = ?) OR (senderId = ? AND receiveId = ?)
				ORDER BY timeMessage ASC;`

	rows, err := m.config.DB.Query(query, senderId, receiveId, receiveId, senderId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*entities.Message
	for rows.Next() {
		var message entities.Message
		err := rows.Scan(&message.ID, &message.SenderId, &message.ReceiveId, &message.Content, &message.TimeMessage)
		if err != nil {
			return nil, err
		}
		messages = append(messages, &message)
	}
	return messages, nil
}

func (m *Mysql) SendMessage(message *entities.Message) error {
	query := `INSERT INTO messages (senderId, receiveId, content, timeMessage) VALUES (?, ?, ?, ?)`
	_, err := m.config.DB.Exec(query, message.SenderId, message.ReceiveId, message.Content, message.TimeMessage)
	return err
}

func (m *Mysql) DeleteMessage(id int) error {
	query := `DELETE FROM messages WHERE idMessage = ?`
	_, err := m.config.DB.Exec(query, id)

	result, err := m.config.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("no existe el mensaje con id %d", id)
	}

	return nil
}
