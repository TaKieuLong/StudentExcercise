package main

import "fmt"

type Student struct {
	name   string
	age    int
	gender string
}

func searchStudent(name string, studentMap map[int]Student) (Student, bool) {
	for _, student := range studentMap {
		if student.name == name {
			return student, true
		}
	}
	return Student{}, false
}

// Hàm thêm sinh viên
func addStudent(id int, name string, age int, gender string, studentMap map[int]Student) {
	studentMap[id] = Student{name: name, age: age, gender: gender}
	fmt.Printf("Đã thêm sinh viên: %s\n", name)
}

// Hàm cập nhật sinh viên
func updateStudent(id int, name string, age int, gender string, studentMap map[int]Student) {
	if _, exists := studentMap[id]; exists {
		studentMap[id] = Student{name: name, age: age, gender: gender}
		fmt.Printf("Đã cập nhật sinh viên với ID %d\n", id)
	} else {
		fmt.Printf("Không tìm thấy sinh viên với ID %d để cập nhật\n", id)
	}
}

// Hàm xóa sinh viên
func deleteStudent(id int, studentMap map[int]Student) {
	if _, exists := studentMap[id]; exists {
		delete(studentMap, id)
		fmt.Printf("Đã xóa sinh viên với ID %d\n", id)
	} else {
		fmt.Printf("Không tìm thấy sinh viên với ID %d để xóa\n", id)
	}
}

// Hàm xuất ra danh sách sinh viên theo giới tính
func exportByGender(gender string, studentMap map[int]Student) {
	fmt.Printf("Danh sách sinh viên có giới tính %s:\n", gender)
	for _, student := range studentMap {
		if student.gender == gender {
			fmt.Printf("Tên: %s, Tuổi: %d\n", student.name, student.age)
		}
	}
}
func main() {
	var studentMap = make(map[int]Student)
	studentMap[1] = Student{
		name:   "Long",
		age:    12,
		gender: "nam",
	}
	studentMap[2] = Student{
		name:   "Toàn",
		age:    200,
		gender: "pede",
	}
	studentMap[3] = Student{
		name:   "Trọng",
		age:    20,
		gender: "nữ",
	}
	studentToFind := "Toàn"

	addStudent(1, "Khang", 20, "nam", studentMap)

	updateStudent(2, "Toàn Năng", 21, "nam", studentMap)

	deleteStudent(3, studentMap)

	if student, found := searchStudent(studentToFind, studentMap); found {
		fmt.Printf("Đã tìm thấy sinh viên %s %d tuổi và có giới tính là %s\n", student.name, student.age, student.gender)
	} else {
		fmt.Printf("Không tìm thấy sinh viên %s\n", studentToFind)
	}

	exportByGender("nam", studentMap)
}
