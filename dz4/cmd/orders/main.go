package main

import (
	"dz4/internal/config"
	"dz4/internal/handler"
	"dz4/internal/repository"
	"dz4/internal/service"
	"fmt"
)

func main() {
	repo := repository.NewMemory()
	orders := service.NewService(repo)
	handler := handler.NewHandler(orders)

	address := config.GetAddress()
	order, err := handler.Create("Покупатель")
	if err != nil {
		fmt.Println("ошибка:", err)
		return
	}
	fmt.Printf("сервис слушал бы %s, создан заказ: %+v\n", address, order)
}
