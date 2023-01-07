package routes

import (
	"github.com/Emotions-s/go-classroom-management-basic/pkg/controllers"
	"github.com/gorilla/mux"
)

var ClassManagerRoutes = func(router *mux.Router) {
	router.HandleFunc("/classroom", controllers.CreateClassroom).Methods("POST")
	router.HandleFunc("/classroom", controllers.GetClassroom).Methods("GET")
	router.HandleFunc("/classroom/{roomId}", controllers.GetClassroomById).Methods("GET")
	router.HandleFunc("/classroom/{roomId}", controllers.UpdateClassroom).Methods("PUT")
	router.HandleFunc("/classroom/{roomId}", controllers.DeleteClassroom).Methods("DELETE")
	router.HandleFunc("/student", controllers.CreateStudent).Methods("POST")
	router.HandleFunc("/student/{studentId}", controllers.GetStudentById).Methods("GET")
	router.HandleFunc("/student/{studentId}", controllers.UpdateStudent).Methods("PUT")
	router.HandleFunc("/student/{studentId}", controllers.DeleteStudent).Methods("DELETE")
	router.HandleFunc("/teacher", controllers.CreateStudent).Methods("POST")
	router.HandleFunc("/teacher/{teacherId}", controllers.GetTeacherById).Methods("GET")
	router.HandleFunc("/teacher/{teacherId}", controllers.UpdateTeacher).Methods("PUT")
	router.HandleFunc("/teacher/{teacherId}", controllers.DeleteTeacher).Methods("DELETE")
}
