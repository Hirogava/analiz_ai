package main

import (
	"duabi/db"
	"duabi/ai"
	"duabi/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
    if err != nil {
        fmt.Println("Ошибка загрузки файла .env")
        return
    }

	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=%s password=%s port=5432",
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_MODE"),
		os.Getenv("DB_PASSWORD"),
	)

	dB, err := db.NewDBManager("postgres", connStr)
	if err != nil{
		log.Fatalf("Ошибка при подключении к базе данных: %s", err)
	}
	defer dB.Close()

	err = db.Migrate(dB)
	if err != nil {
		log.Fatalf("Ошибка миграции базы данных: %s", err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("static/templates/index.html"))
    	tmpl.Execute(w, nil)
	}).Methods("GET")

	r.HandleFunc("/treatment", func(w http.ResponseWriter, r *http.Request) {
		var answer models.Answer
		question := r.FormValue("question")
		answer.Question = question
		questionId, err := ai.GetCategory(question)

		if err != nil {
			fmt.Println("Ошибка при получении ответа от AI:", err)
			return
		} else if questionId == -1 {
			answer.Answer = "Ваш вопрос не найден, пожалуйста, напишите в техподдержку"
		} else{
			answer.Answer, err = ai.GetAnswer(dB, questionId, question)
			if err != nil {
				fmt.Println("Ошибка при получении ответа от AI:", err)
				return
			}
		}

		tmpl := template.Must(template.ParseFiles("static/templates/response.html"))
    	tmpl.Execute(w, answer)
	}).Methods("POST")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	server := &http.Server{
		Addr: ":8080",
		Handler: r,
	}

	fmt.Println("Server is listening...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
