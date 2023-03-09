/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/RustamRR/cli-wether-app/internal/app"
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
		weather := server.GetWeatherForCity(CityName)

		fmt.Printf(
			"%s. \nТемпература воздуха %.2f. \nСкорость ветра: %.2f. \nНаправление ветра: %.2f.\n",
			weather.Name,
			weather.Temperature,
			weather.WindSpeed,
			weather.WindDirection,
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
