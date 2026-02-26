package sql

import (
	config "chat/Src/Core"
	entities "chat/Src/Endpoint/Contacts/Domain/Entities"
	userEntities "chat/Src/Endpoint/User/Domain/Entities"
	"fmt"
	"strings"
)

type Mysql struct {
	config *config.ConnMySQL
}

func NewMySQL() (*Mysql, error) {
	conn := config.GetDBPool()
	if conn.Err != "" {
		return nil, fmt.Errorf("error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &Mysql{config: conn}, nil
}

func (m *Mysql) GetAll(userID int) ([]entities.ContactResponse, error) {
	query := `
        SELECT 
            c.id,
            c.userId,
            c.contactId,
            c.createdAt,
            u.firstname,
            u.lastname,
            u.username,
            u.email,
            u.role
        FROM user_contacts c
        JOIN users u ON c.contactId = u.id
        WHERE c.userId = ?
    `

	rows, err := m.config.DB.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener contactos: %v", err)
	}
	defer rows.Close()

	var contacts []entities.ContactResponse

	for rows.Next() {
		var c entities.ContactResponse
		if err := rows.Scan(
			&c.ID,
			&c.UserID,
			&c.ContactID,
			&c.CreatedAt,
			&c.FirstName,
			&c.LastName,
			&c.Username,
			&c.Email,
			&c.Role,
		); err != nil {
			return nil, fmt.Errorf("error al escanear contacto: %v", err)
		}
		contacts = append(contacts, c)
	}

	return contacts, nil
}

func (m *Mysql) CreateContact(contact entities.Contact) error {
	query := `
        INSERT INTO user_contacts (userId, contactId)
        VALUES (?, ?)
    `

	_, err := m.config.DB.Exec(query, contact.UserID, contact.ContactID)
	if err != nil {
		// Error de foreign key (usuario no existe)
		if strings.Contains(err.Error(), "Error 1452") {
			return fmt.Errorf("contacto no existe")
		}
		return fmt.Errorf("error al crear contacto: %v", err)
	}

	return nil
}

func (m *Mysql) GetContactByName(username string) (*userEntities.User, error) {
	query := ` SELECT id, firstname, lastname, username FROM users WHERE username = ? LIMIT 1 `

	var u userEntities.User

	err := m.config.DB.QueryRow(query, username).Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Username,
		&u.Role,
	)

	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return nil, nil // no existe, el use case lo manejará
		}
		return nil, fmt.Errorf("error al obtener usuario por username: %v", err)
	}

	return &u, nil
}

func (m *Mysql) DeleteContact(userID int, contactID int) error {
	query := `
        DELETE FROM user_contacts
        WHERE userId = ? AND contactId = ?
    `

	_, err := m.config.DB.Exec(query, userID, contactID)
	if err != nil {
		return fmt.Errorf("error al eliminar contacto: %v", err)
	}

	return nil
}

func (m *Mysql) Exists(userID int, contactID int) (bool, error) {
	query := `
        SELECT COUNT(*) 
        FROM user_contacts
        WHERE userId = ? AND contactId = ?
    `

	var count int
	err := m.config.DB.QueryRow(query, userID, contactID).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error al verificar existencia del contacto: %v", err)
	}

	return count > 0, nil
}
