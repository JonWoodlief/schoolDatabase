package main

import (
    "./a"
    //"fmt"
    "math/rand"
    "time"
    //"net/http"
)

func newStudents(n int) []databaseTypes.Student {

    students := make([]databaseTypes.Student, n)

    for i := range students {
        students[i].SetID(300+i)
    }

    return students
}
func newClasses(n int) []databaseTypes.Class {
    classes := make([]databaseTypes.Class, n)

    for i := range classes {
        classes[i].SetID(i)
        classes[i].SetRoster(make([]databaseTypes.Student, 0))
    }
    return classes
}
func newTeachers(n int) []databaseTypes.Teacher {

    teachers := make([]databaseTypes.Teacher, n)

    for i := range teachers {
        teachers[i].SetID(100+i)
    }

    return teachers
}

func main () {
    //create a database of students, teachers, and classes randomly assigned to each other
    numStudents := 120
    classSize := 30
    numClasses := 15
    numTeachers := 5

    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    
    students := newStudents(numStudents)
    classes := newClasses(numClasses)
    teachers := newTeachers(numTeachers)
 
    for _, class := range classes {
        for _, studentIndex := range r.Perm(120)[:classSize]   {
            class.AddStudent(students[studentIndex])
        }
        teacherIndex := r.Intn(numTeachers)

        class.UpdateTeacher(teachers[teacherIndex])
    }

    //receive API calls
}
