package app

import "testing_mock/api/controllers/weather_controller"

func routes() {
	router.GET("/weather/:apiKey/:latitude/:longitude", weather_controller.GetWeather)
}