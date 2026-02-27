package sql

import (
	config "chat/Src/Core"
	entities "chat/Src/Endpoint/Class/Domain/Entities"
	"database/sql"
)

type ClassSQL struct {
	config *config.ConnMySQL
}

func NewClassSQL() (*ClassSQL, error) {
	conn := config.GetDBPool()
	return &ClassSQL{config: conn}, nil
}

func (c *ClassSQL) GetAllClasses() ([]entities.ClassWithTeacher, error) {
    query := `
        SELECT 
            cl.idClass,
            cl.teacherId,
            cl.title,
            cl.description,
            cl.classDate,
            cl.startTime,
            cl.endTime,
            cl.capacity,
            cl.status,
            u.firstname,
            u.lastname
        FROM classes cl
        JOIN users u ON cl.teacherId = u.id
        WHERE cl.status = 'Activa'
    `

    rows, err := c.config.FetchRows(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var classes []entities.ClassWithTeacher
    for rows.Next() {
        class, err := scanClassWithTeacher(rows)
        if err != nil {
            return nil, err
        }
        classes = append(classes, class)
    }
    return classes, nil
}



func (c *ClassSQL) GetClassByID(classID int64) (*entities.Class, error) {
	query := `SELECT idClass, teacherId, title, description, classDate, startTime, endTime, capacity, status FROM classes WHERE idClass = ?`
	row := c.config.QueryRow(query, classID)
	class, err := scanClassRow(row)
	if err != nil {
		return nil, err
	}
	return &class, nil
}

func (c *ClassSQL) GetClassesByTeacherID(teacherID int) ([]entities.Class, error) {
	query := `SELECT idClass, teacherId, title, description, classDate, startTime, endTime, capacity, status FROM classes WHERE teacherId = ? AND status = 'Activa'`
	rows, err := c.config.FetchRows(query, teacherID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var classes []entities.Class
	for rows.Next() {
		class, err := scanClass(rows)
		if err != nil {
			return nil, err
		}
		classes = append(classes, class)
	}
	return classes, nil
}

func (c *ClassSQL) CreateClass(class *entities.Class) (*entities.Class, error) {
	query := `INSERT INTO classes (teacherId, title, description, classDate, startTime, endTime, capacity, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := c.config.ExecutePreparedQuery(query, class.TeacherID, class.Title, class.Description, class.ClassDate, class.StartTime, class.EndTime, class.Capacity, "Activa")
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	class.ID = id
	class.Status = "Activa"
	return class, nil
}

func (c *ClassSQL) UpdateClass(class *entities.Class) error {
	query := `UPDATE classes SET title = ?, description = ?, classDate = ?, startTime = ?, endTime = ?, capacity = ?, status = ? WHERE idClass = ?`
	_, err := c.config.ExecutePreparedQuery(query, class.Title, class.Description, class.ClassDate, class.StartTime, class.EndTime, class.Capacity, class.Status, class.ID)
	return err
}

func (c *ClassSQL) DeleteClass(classID int64) error {
	query := `UPDATE classes SET status = 'Cancelada' WHERE idClass = ?`
	_, err := c.config.ExecutePreparedQuery(query, classID)
	return err
}

func scanClass(rows *sql.Rows) (entities.Class, error) {
    var class entities.Class
    var classDate string

    err := rows.Scan(
        &class.ID,
        &class.TeacherID,
        &class.Title,
        &class.Description,
        &classDate,
        &class.StartTime,
        &class.EndTime,
        &class.Capacity,
        &class.Status,
    )
    if err != nil {
        return class, err
    }

    class.ClassDate = classDate
    return class, nil
}


func scanClassRow(row *sql.Row) (entities.Class, error) {
    var class entities.Class
    var classDate string

    err := row.Scan(
        &class.ID,
        &class.TeacherID,
        &class.Title,
        &class.Description,
        &classDate,
        &class.StartTime,
        &class.EndTime,
        &class.Capacity,
        &class.Status,
    )
    if err != nil {
        return class, err
    }

    class.ClassDate = classDate
    return class, nil
}


func (c *ClassSQL) HasScheduleConflict(teacherID int, startTime, endTime string, classDate string) (bool, error) {
    query := `
        SELECT COUNT(*)
        FROM classes
        WHERE teacherId = ?
          AND classDate = ?
          AND status = 'Activa'
          AND (
                startTime < ?
            AND endTime   > ?
          )
    `

    var count int
    err := c.config.DB.QueryRow(query, teacherID, classDate, endTime, startTime).Scan(&count)
    if err != nil {
        return false, err
    }

    return count > 0, nil
}

func scanClassWithTeacher(rows *sql.Rows) (entities.ClassWithTeacher, error) {
    var class entities.ClassWithTeacher
    var classDate string

    err := rows.Scan(
        &class.ID,
        &class.TeacherID,
        &class.Title,
        &class.Description,
        &classDate,
        &class.StartTime,
        &class.EndTime,
        &class.Capacity,
        &class.Status,
        &class.TeacherFirstName,
        &class.TeacherLastName,
    )
    if err != nil {
        return class, err
    }

    class.ClassDate = classDate
    return class, nil
}

func (m *ClassSQL) GetClassesByDate(date string) ([]entities.ClassWithTeacher, error) {

    query := `
        SELECT 
            cl.idClass,
            cl.teacherId,
            cl.title,
            cl.description,
            cl.classDate,
            cl.startTime,
            cl.endTime,
            cl.capacity,
            cl.status,
            u.firstname,
            u.lastname
        FROM classes cl
        JOIN users u ON cl.teacherId = u.id
        WHERE DATE(cl.classDate) = ?
        AND CONCAT(cl.classDate, ' ', cl.startTime) > NOW()
        ORDER BY cl.startTime ASC
    `

    rows, err := m.config.DB.Query(query, date)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var classes []entities.ClassWithTeacher

    for rows.Next() {
        var c entities.ClassWithTeacher
        err := rows.Scan(
            &c.ID,
            &c.TeacherID,
            &c.Title,
            &c.Description,
            &c.ClassDate,
            &c.StartTime,
            &c.EndTime,
            &c.Capacity,
            &c.Status,
            &c.TeacherFirstName,
            &c.TeacherLastName,
        )
        if err != nil {
            return nil, err
        }
        classes = append(classes, c)
    }

    return classes, nil
}
