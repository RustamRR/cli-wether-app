/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/RustamRR/cli-wether-app/internal/app"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var CityName string

// getTodayCmd represents the getToday command
var getTodayCmd = &cobra.Command{
	Use:   "getToday",
	Short: "Погода на сегодня",
	Long: `Погода на сегодня в выбранном городе
Название города необходимо вводить на английском, например Bishkek
`,
	Run: func(cmd *cobra.Command, args []string) {
		server := app.New(viper.GetViper())
		weather, err := server.GetWeatherForCity(CityName)
		if err != nil {
			color.Red("Не удалось узнать погоду")
			return
		}

		temperature := fmt.Sprintf("%.2f°C", weather.Temperature)
		windSpeed := fmt.Sprintf("%.2f м/с", weather.WindSpeed)
		windDirection := fmt.Sprintf("%.2f°", weather.WindDirection)

		fmt.Printf(
			"%s. \nТемпература воздуха %s. \nСкорость ветра: %s. \nНаправление ветра: %s.\n",
			color.New(color.FgHiGreen).SprintfFunc()(weather.Name),
			color.New(color.FgRed).SprintfFunc()(temperature),
			color.New(color.FgBlue).SprintfFunc()(windSpeed),
			color.New(color.FgCyan).SprintfFunc()(windDirection),
		)
	},
}

func init() {
	getTodayCmd.Flags().StringVarP(&CityName, "city", "c", "Berlin", "Название вашего города на английском")
	rootCmd.AddCommand(getTodayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getTodayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getTodayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
