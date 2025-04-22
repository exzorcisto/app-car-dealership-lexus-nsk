package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"app_car_dealership_lexus_nsk/models" // Замените "your_module" на имя вашего модуля

	"github.com/gorilla/mux"
)

// db - глобальное подключение к базе данных (возможно, лучше переместить в отдельный пакет db)
var db *sql.DB

// SetDB устанавливает соединение с базой данных для обработчиков.
func SetDB(database *sql.DB) {
	db = database
}

// GetModelsHandler обрабатывает GET-запрос для получения всех моделей автомобилей.
func GetModelsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Устанавливаем заголовок Content-Type для JSON

	rows, err := db.Query("SELECT ModelID, ModelName, Description FROM Models") // Запрос на получение всех моделей
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Отправляем ошибку сервера, если запрос не удался
		log.Printf("Ошибка при запросе моделей: %v", err)          // Логируем ошибку
		return
	}
	defer rows.Close() // Закрываем rows после завершения функции

	var modelsList []models.Model // Создаем слайс для хранения моделей
	for rows.Next() {
		var model models.Model                                                 // Создаем переменную для одной модели
		err := rows.Scan(&model.ModelID, &model.ModelName, &model.Description) // Сканируем значения из строки базы данных в структуру model
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) // Отправляем ошибку сервера, если сканирование не удалось
			log.Printf("Ошибка при сканировании строки: %v", err)      // Логируем ошибку
			return
		}
		modelsList = append(modelsList, model) // Добавляем модель в слайс
	}

	err = rows.Err() // Проверяем, не было ли ошибок во время итерации
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Отправляем ошибку сервера, если произошла ошибка
		log.Printf("Ошибка при итерации по строкам: %v", err)      // Логируем ошибку
		return
	}

	json.NewEncoder(w).Encode(modelsList)  // Кодируем слайс моделей в JSON и отправляем в ответ
	log.Println("Модели успешно получены") // Логируем успех
}

// GetModelHandler обрабатывает GET-запрос для получения модели автомобиля по ID.
func GetModelHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Устанавливаем заголовок Content-Type для JSON
	vars := mux.Vars(r)                                // Получаем переменные маршрута из запроса
	id := vars["id"]                                   // Получаем ID из переменных маршрута

	var model models.Model                                                                                                                                     // Создаем переменную для хранения модели
	err := db.QueryRow("SELECT ModelID, ModelName, Description FROM Models WHERE ModelID = $1", id).Scan(&model.ModelID, &model.ModelName, &model.Description) // Запрашиваем модель из базы данных по ID
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Модель не найдена", http.StatusNotFound) // Отправляем ошибку 404, если модель не найдена
			log.Printf("Модель с ID %s не найдена", id)             // Логируем ошибку
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)   // Отправляем ошибку сервера, если запрос не удался
		log.Printf("Ошибка при запросе модели с ID %s: %v", id, err) // Логируем ошибку
		return
	}

	json.NewEncoder(w).Encode(model)                  // Кодируем модель в JSON и отправляем в ответ
	log.Printf("Модель с ID %s успешно получена", id) // Логируем успех
}

// CreateModelHandler обрабатывает POST-запрос для создания новой модели автомобиля.
func CreateModelHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Устанавливаем заголовок Content-Type для JSON

	var model models.Model                        // Создаем переменную для хранения модели
	err := json.NewDecoder(r.Body).Decode(&model) // Декодируем JSON из тела запроса в структуру model
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)            // Отправляем ошибку 400, если декодирование не удалось
		log.Printf("Ошибка при декодировании тела запроса: %v", err) // Логируем ошибку
		return
	}

	_, err = db.Exec("INSERT INTO Models (ModelName, Description) VALUES ($1, $2)", model.ModelName, model.Description) // Вставляем новую модель в базу данных
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Отправляем ошибку сервера, если вставка не удалась
		log.Printf("Ошибка при вставке модели: %v", err)           // Логируем ошибку
		return
	}

	w.WriteHeader(http.StatusCreated)                                                 // Устанавливаем код статуса 201 Created
	json.NewEncoder(w).Encode(map[string]string{"message": "Модель успешно создана"}) // Отправляем сообщение об успехе
	log.Println("Новая модель успешно создана")                                       // Логируем успех
}

// UpdateModelHandler обрабатывает PUT-запрос для обновления существующей модели автомобиля.
func UpdateModelHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Устанавливаем заголовок Content-Type для JSON
	vars := mux.Vars(r)                                // Получаем переменные маршрута из запроса
	id := vars["id"]                                   // Получаем ID из переменных маршрута

	var model models.Model                        // Создаем переменную для хранения модели
	err := json.NewDecoder(r.Body).Decode(&model) // Декодируем JSON из тела запроса в структуру model
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)            // Отправляем ошибку 400, если декодирование не удалось
		log.Printf("Ошибка при декодировании тела запроса: %v", err) // Логируем ошибку
		return
	}

	result, err := db.Exec("UPDATE Models SET ModelName = $1, Description = $2 WHERE ModelID = $3", model.ModelName, model.Description, id) // Обновляем модель в базе данных
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Отправляем ошибку сервера, если обновление не удалось
		log.Printf("Ошибка при обновлении модели: %v", err)        // Логируем ошибку
		return
	}

	rowsAffected, err := result.RowsAffected() // Получаем количество затронутых строк
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)              // Отправляем ошибку сервера, если получение количества не удалось
		log.Printf("Ошибка при получении количества затронутых строк: %v", err) // Логируем ошибку
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Модель не найдена", http.StatusNotFound) // Отправляем ошибку 404, если модель не найдена
		log.Printf("Модель с ID %s не найдена", id)             // Логируем ошибку
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Модель успешно обновлена"}) // Отправляем сообщение об успехе
	log.Printf("Модель с ID %s успешно обновлена", id)                                  // Логируем успех
}

// DeleteModelHandler обрабатывает DELETE-запрос для удаления модели автомобиля.
func DeleteModelHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Устанавливаем заголовок Content-Type для JSON
	vars := mux.Vars(r)                                // Получаем переменные маршрута из запроса
	id := vars["id"]                                   // Получаем ID из переменных маршрута

	result, err := db.Exec("DELETE FROM Models WHERE ModelID = $1", id) // Удаляем модель из базы данных

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // Отправляем ошибку сервера, если удаление не удалось
		log.Printf("Ошибка при удалении модели: %v", err)          // Логируем ошибку
		return
	}

	rowsAffected, err := result.RowsAffected() // Получаем количество затронутых строк
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)              // Отправляем ошибку сервера, если получение количества не удалось
		log.Printf("Ошибка при получении количества затронутых строк: %v", err) // Логируем ошибку
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "Модель не найдена", http.StatusNotFound) // Отправляем ошибку 404, если модель не найдена
		log.Printf("Модель с ID %s не найдена", id)             // Логируем ошибку
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Модель успешно удалена"}) // Отправляем сообщение об успехе
	log.Printf("Модель с ID %s успешно удалена", id)                                  // Логируем успех
}
