package Models

import (
	"awesomeProject/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllStudents(students *[]Student) (err error) {
	if err := Config.DB.Model(students).Preload("Scores").Find(students).Error; err != nil {
		return err
	}
	return nil
}

func AddNewStudent(student *Student) error {
	err := Config.DB.Create(student).Error
	fmt.Println(student)
	if err != nil {
		return err
	}
	return nil
}

func GetStudentById(student *Student, id string) (err error) {
	if err := Config.DB.Model(student).Where("id = ?", id).Preload("Scores").First(student).Error; err != nil {
		return err
	}
	return nil
}

func UpdateStudent(student *Student) (err error) {
	fmt.Println(student)
	if err := Config.DB.Save(student).Error; err != nil {
		return err
	}
	return nil
}

func DeleteStudent(student *Student, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).Delete(student).Error; err != nil {
		return err
	}
	return nil
}
