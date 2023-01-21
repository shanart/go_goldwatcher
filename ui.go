package main

import "fyne.io/fyne/v2/container"

func (app *Config) makeUI() {
	// Get current price of gold ( https://data-asg.goldprice.org/dbXrates/USD )
	openPrice, currentPrice, priceChange := app.getPriceText()

	// put price information into a container
	priceContent := container.NewGridWithColumns(3,
		openPrice,
		currentPrice,
		priceChange,
	)

	app.PriceContainer = priceContent

	// add container to window
	finalContent := container.NewVBox(priceContent)

	app.MainWindow.SetContent(finalContent)
}
