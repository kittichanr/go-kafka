package main

import (
	"producer/controllers"
	"producer/services"
	"strings"

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {
	producer, err := sarama.NewSyncProducer(viper.GetStringSlice("kafka.servers"), nil)
	if err != nil {
		panic(err)
	}

	defer producer.Close()

	eventProducer := services.NewEventProducer(producer)
	accountService := services.NewAccountService(eventProducer)
	accountController := controllers.NewAccountController(accountService)

	app := fiber.New()

	app.Post("/openaccount", accountController.OpenAccount)
	app.Post("/depositfund", accountController.DepositFund)
	app.Post("/withdrawfund", accountController.WithdrawFund)
	app.Post("/closeaccount", accountController.CloseAccount)

	app.Listen(":8000")
}

// func main() {
// 	servers := []string{"localhost:9092"}

// 	producer, err := sarama.NewSyncProducer(servers, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer producer.Close()

// 	msg := sarama.ProducerMessage{
// 		Topic: "petchhello",
// 		Value: sarama.StringEncoder("Hello world"),
// 	}
// 	p, o, err := producer.SendMessage(&msg)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("partition=%v, offset=%v", p, o)

// }
