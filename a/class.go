package databaseTypes

type Class struct {
	id      int
	roster  []Student
	teacher Teacher
}

func (c Class) GetID() int {
	return c.id
}
func (c Class) SetID(newID int) int {
	c.id = newID
	if c.id == newID {
		return 0
	}
	return 1 //failed to set id
}
func (c Class) GetRoster() []Student {
    return c.roster
}
func (c Class) SetRoster(newRoster []Student) {
    c.roster = newRoster
}
func (c Class) GetTeacher() Teacher {
    return c.teacher
}
func (c Class) AddStudent(newStu Student) int {
    for _, student := range c.roster {
        if student.id == newStu.id {
            return 1 //student already in class
        }
    }

    c.roster = append(c.roster, newStu)
    return 0
}
func (c Class) DelStudent(delStu Student) int {
    for i, student := range c.roster {
        if student.id == delStu.id {
            c.roster = append(c.roster[:i], c.roster[i+1:]...)
            return 0
        }
    }

    return 1 //student not found
}
func (c Class) UpdateTeacher(newTeacher Teacher) int {
    if c.teacher.id == newTeacher.id {
        return 1 //teacher is already the head of the class
    }
    newTeacher.DelClass(c)
    c.teacher.AddClass(c)
    c.teacher = newTeacher
    return 0
}
