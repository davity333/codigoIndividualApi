package sql

import (
	config "chat/Src/Core"
	entities "chat/Src/Endpoint/Reservations/Domain/Entities"
	"database/sql"
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

func (m *Mysql) GetAllReservations() ([]*entities.Reservation, error) {
	query := `SELECT idReservation, studentId, classId, reservationDate,attendance
				FROM reservations
				ORDER BY idReservation DESC`

	rows, err := m.config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservations []*entities.Reservation
	for rows.Next() {
		reservation, err := scanReservation(rows)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, reservation)
	}

	return reservations, nil
}

func (m *Mysql) GetReservationByID(id int) (*entities.Reservation, error) {
	query := `SELECT idReservation, studentId, classId, reservationDate, attendance
				FROM reservations
				WHERE idReservation = ?`

	rows, err := m.config.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		reservation, scanErr := scanReservation(rows)
		if scanErr != nil {
			return nil, scanErr
		}
		return reservation, nil
	}

	return nil, fmt.Errorf("no existe la reservación con id %d", id)
}

func (m *Mysql) CreateReservation(reservation *entities.Reservation) error {
	query := `INSERT INTO reservations (studentId, classId, reservationDate, attendance)
				VALUES (?, ?, ?, ?)`

	result, err := m.config.DB.Exec(query,
		reservation.StudentID,
		reservation.ClassID,
		reservation.ReservationDate,
		reservation.Attendance,
	)
	if err != nil {
		return err
	}

	lastInsertID, err := result.LastInsertId()
	if err == nil {
		reservation.ID = int(lastInsertID)
	}

	return nil
}

func (m *Mysql) UpdateReservation(id int, reservation *entities.Reservation) error {
	query := `UPDATE reservations
				SET studentId = ?, classId = ?, reservationDate = ?, attendance = ?
				WHERE idReservation = ?`

	result, err := m.config.DB.Exec(query,
		reservation.StudentID,
		reservation.ClassID,
		reservation.ReservationDate,
		reservation.Attendance,
		id,
	)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("no existe la reservación con id %d", id)
	}

	return nil
}

func (m *Mysql) DeleteReservation(id int) error {
	query := `DELETE FROM reservations WHERE idReservation = ?`

	result, err := m.config.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("no existe la reservación con id %d", id)
	}

	return nil
}

func scanReservation(scanner interface {
	Scan(dest ...interface{}) error
}) (*entities.Reservation, error) {
	var reservation entities.Reservation
	var attendance sql.NullBool

	err := scanner.Scan(
		&reservation.ID,
		&reservation.StudentID,
		&reservation.ClassID,
		&reservation.ReservationDate,
		&attendance,
	)
	if err != nil {
		return nil, err
	}

	if attendance.Valid {
		value := attendance.Bool
		reservation.Attendance = &value
	} else {
		reservation.Attendance = nil
	}

	return &reservation, nil
}
