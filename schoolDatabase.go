package main

import (
    //"fmt"
    "math/rand"
    "time"
)

type Student struct {
    id int
}
type Class struct {
    id int
    roster []Student
}
type Teacher struct {
    id int
    classes []Class
}

func newStudents(n int) []Student {

    students := make([]Student, n)

    for i := range students {
        students[i].id = 300 + i
    }

    return students
}
func newClasses(n int) []Class {
    classes := make([]Class, n)

    for i := range classes {
        classes[i].id = i
        classes[i].roster = make([]Student, 0)
    }

    return classes
}
func newTeachers(n int) []Teacher {

    teachers := make([]Teacher, n)

    for i := range teachers {
        teachers[i].id = 100 + i
    }

    return teachers
}


func main() {
    numStudents := 120
    classSize := 30
    numClasses := 15
    numTeachers := 5

    r := rand.New(rand.NewSource(time.Now().UnixNano()))
     
    students := newStudents(numStudents)
    classes := newClasses(numClasses)
    teachers := newTeachers(numTeachers)
 
    for classIndex := range classes {
        for _, studentIndex := range r.Perm(120)[:classSize]   {
            classes[classIndex].roster = append(classes[classIndex].roster, students[studentIndex])
        }
        teacherIndex := r.Intn(numTeachers)

        teachers[teacherIndex].classes = append(teachers[teacherIndex].classes, classes[classIndex])
    }
}
