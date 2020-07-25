package app

import "github.com/milanojs/gofordevops/testing_mock/api/controllers/weather_controller"

func routes() {
	router.GET("/weather/:apiKey/:latitude/:longitude", weather_controller.GetWeather)
}
