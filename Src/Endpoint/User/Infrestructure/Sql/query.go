package sql

import (
	config "chat/Src/Core"
	entities "chat/Src/Endpoint/User/Domain/Entities"
	"fmt"
	"log"
)

type Mysql struct {
	config *config.ConnMySQL
}

func NewMySQL() (*Mysql, error) {
	conn := config.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &Mysql{config: conn}, nil
}

func (m *Mysql) GetUserByName(username string) ([]*entities.User, error) {
    query := `
        SELECT id, username, email, password, firstname, lastname, rol
        FROM users
        WHERE username = ?
    `

    rows, err := m.config.DB.Query(query, username)
    if err != nil {
        return nil, fmt.Errorf("error al buscar usuario: %v", err)
    }
    defer rows.Close()

    var users []*entities.User

    for rows.Next() {
        var u entities.User
        if err := rows.Scan(
            &u.ID,
            &u.Username,
            &u.Email,
            &u.Password,
            &u.FirstName,
            &u.LastName,
            &u.Role,
        ); err != nil {
            return nil, fmt.Errorf("error al escanear usuario: %v", err)
        }

        users = append(users, &u)
    }

    if len(users) == 0 {
        return nil, fmt.Errorf("no existe el usuario")
    }

    return users, nil
}


func (m *Mysql) GetAllUsers() ([]*entities.User, error) {
	query := "SELECT id, username, email, firstname, lastname, rol FROM users"
	rows, err := m.config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entities.User
	for rows.Next() {
		var user entities.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.FirstName, &user.LastName, &user.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (mysql *Mysql) LoginUser(email string, password string) (*entities.User, error) {
    query := "SELECT id, username, email, password, firstname, lastname, rol FROM users WHERE email = ?"
    rows, err := mysql.config.FetchRows(query, email)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var user entities.User
    if rows.Next() {
        err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.Role)
        if err != nil {
            return nil, err
        }
    } else {
        return nil, fmt.Errorf("usuario no encontrado")
    }

    return &user, nil
}

func (mysql *Mysql) CreateUser(save *entities.User) error {

	query := "INSERT INTO users (username, email, password, firstname, lastname, rol) VALUES (?, ?, ?, ?, ?, ?)"

	result, err := mysql.config.ExecutePreparedQuery(query, save.Username, save.Email, save.Password, save.FirstName, save.LastName, save.Role)

	if err != nil {
		log.Println("Error al insertar el usuario:", save.Email, save.Username, save.Password, err)
		return err
	}
	if result != nil {
		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 1 {
			log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)

			if err != nil {
				fmt.Println(err)
				return err
			}

		} else {
			log.Printf("[MySQL] - Ninguna fila fue afectada.")
		}
	} else {
		log.Printf("[MySQL] - Resultado de la consulta es nil.")
	}
	return nil
}
