package practiceSQL

// pre-requisites: local PgSQL installation
// my_db database creation

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	// Пустой импорт регистрирует драйвер "pgx" во встроенной библиотеке database/sql
	_ "github.com/jackc/pgx/v5/stdlib"
)

// Структура, соответствующая таблице в базе данных
type User struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
}

func main1() {
	// Строка подключения (DSN). Замените на ваши данные.
	dsn := "postgres://postgres:postgres@localhost:5432/my_db?sslmode=disable"

	// 1. Инициализируем пул соединений (sql.Open не устанавливает физическое соединение сразу)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Ошибка конфигурации БД: %v", err)
	}
	defer db.Close() // Закрываем пул при завершении программы

	// 2. Настраиваем пул соединений (важно для продакшена)
	db.SetMaxOpenConns(25)                 // Максимальное количество активных соединений
	db.SetMaxIdleConns(25)                 // Сколько простаивающих соединений держать открытыми
	db.SetConnMaxLifetime(5 * time.Minute) // Время жизни соединения (защищает от утечек)

	// Создаем контекст с таймаутом на случай, если база будет недоступна
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 3. Проверяем реальную доступность базы данных (Ping)
	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("База данных недоступна: %v", err)
	}
	fmt.Println("Успешное подключение к базе данных!")

	// 4. Выполняем команду создания таблицы (Exec)
	_, err = db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			email TEXT UNIQUE NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW()
		);
	`)
	if err != nil {
		log.Fatalf("Не удалось создать таблицу: %v", err)
	}

	// 5. Вставка одной строки (INSERT) c плейсхолдерами во избежание SQL-инъекций
	newUser := User{Name: "Иван", Email: "ivan@example.com"}
	queryInsert := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, created_at`

	// QueryRowContext используем, когда запрос возвращает ровно одну строку (RETURNING)
	err = db.QueryRowContext(ctx, queryInsert, newUser.Name, newUser.Email).
		Scan(&newUser.ID, &newUser.CreatedAt)
	if err != nil {
		log.Printf("Не удалось вставить данные: %v", err)
	} else {
		fmt.Printf("Пользователь успешно создан! ID: %d, Время: %v\n", newUser.ID, newUser.CreatedAt)
	}

	// 6. Получение одной строки по ID (SELECT)
	var user User
	queryGetOne := `SELECT id, name, email, created_at FROM users WHERE id = $1`

	err = db.QueryRowContext(ctx, queryGetOne, newUser.ID).
		Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err == sql.ErrNoRows {
		fmt.Println("Пользователь не найден")
	} else if err != nil {
		log.Fatalf("Ошибка при получении пользователя: %v", err)
	} else {
		fmt.Printf("Получен пользователь по ID: %+v\n", user)
	}

	// 7. Получение списка строк (SELECT)
	queryGetAll := `SELECT id, name, email, created_at FROM users ORDER BY id DESC LIMIT 10`

	rows, err := db.QueryContext(ctx, queryGetAll)
	if err != nil {
		log.Fatalf("Ошибка при запросе списка: %v", err)
	}
	// Обязательно освобождаем соединение обратно в пул!
	defer rows.Close()

	var users []User
	// Итерируемся по полученным строкам
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt); err != nil {
			log.Fatalf("Ошибка сканирования строки: %v", err)
		}
		users = append(users, u)
	}

	// Важная деталь: проверяем, не прервалось ли чтение из-за ошибки (например, сбоя сети)
	if err := rows.Err(); err != nil {
		log.Fatalf("Ошибка при чтении строк: %v", err)
	}

	fmt.Printf("Список последних пользователей: %+v\n", users)
}
