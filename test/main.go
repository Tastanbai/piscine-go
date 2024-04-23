package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Укажите путь к файлу, который вы хотите прочитать
	filePath := "standart.txt"

	// Открываем файл
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Создаем сканер для чтения файла построчно
	scanner := bufio.NewScanner(file)

	// Создаем срез для символов
	var characters []string

	// Читаем файл построчно
	for scanner.Scan() {
		line := scanner.Text()
		characters = append(characters, line)
	}

	// Проверяем ошибки сканера
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Выводим символы
	for _, char := range characters {
		fmt.Println(char)
	}
}
