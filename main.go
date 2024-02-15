package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"log"
	"net/http"
	"workshop/polkadot/blockchain"
)

func (h *Header) Render() app.UI {
	return app.Div().Body(
		app.H1().Body(
			app.Text("Connect Blockchain"),
		).Style("width", "10em").Style("margin-left", "auto").Style("margin-right", "auto"),
		app.P().Body(
			app.Input().Style("width", "500px").
				Type("text").
				Value(h.Address).
				Placeholder("Address blockchain").
				AutoFocus(true).
				OnChange(h.ValueTo(&h.Address)),
		).Style("width", "10em").Style("margin-right", "auto"),

		// Node Version
		app.H1().Body(
			app.Text("Get Node Version"),
		).Style("width", "10em").Style("margin-left", "auto").Style("margin-right", "auto"),
		app.Div().Body(
			app.Div().Body(
				app.Button().Text("Version").OnClick(h.GetNodeVersion).Style("float", "right").Style("height", "200px").Style("width", "500px"),
				app.Div().Body(
					app.H2().Text("Version: "),
					app.If(h.ChainId != "",
						app.H2().Text(h.ChainId),
					).Else(
						app.H2().Text("None"),
					)).Style("overflow", "hidden").Style("height", "200px").Style("width", "300px"),
			),
		).Style("background-color", "deepskyblue").Style("display", "table").Style("clear", "both").Style("display", "block").Style("height", "200px"),

		// Block Number
		app.H1().Body(
			app.Text("Last BlockNumber"),
		).Style("width", "10em").Style("margin-left", "auto").Style("margin-right", "auto"),
		app.Div().Body(
			app.Div().Body(
				app.Button().Text("BlockNumber").OnClick(h.GetBlockNumber).Style("float", "right").Style("height", "200px").Style("width", "500px"),
				app.Div().Body(
					app.H2().Text("Block Number:"),
					app.If(h.BlockNumber != 0,
						app.H2().Text(h.BlockNumber),
					).Else(
						app.H2().Text("0"),
					)).Style("overflow", "hidden").Style("height", "200px").Style("width", "300px"),
			),
		).Style("background-color", "deepskyblue").Style("display", "table").Style("clear", "both").Style("display", "block").Style("height", "200px"),

		// Node Name
		app.H1().Body(
			app.Text("Get Node Name"),
		).Style("width", "10em").Style("margin-left", "auto").Style("margin-right", "auto"),
		app.Div().Body(
			app.Div().Body(
				app.Button().Text("Node Name").OnClick(h.GetNodeName).Style("float", "right").Style("height", "200px").Style("width", "500px"),
				app.Div().Body(
					app.H2().Text("Node Name: \n"),
					app.If(h.NodeName != "",
						app.H2().Text(h.NodeName),
					).Else(
						app.H2().Text("None"),
					)).Style("overflow", "hidden").Style("height", "200px"),
			),
		).Style("background-color", "deepskyblue").Style("display", "table").Style("clear", "both").Style("display", "block").Style("height", "200px"),
	)
}

type Header struct {
	app.Compo
	Address     string
	BlockNumber int64
	Name        string
	NodeName    string
	ChainId     string
}

func (header *Header) GetBlockNumber(ctx app.Context, e app.Event) {
	header.BlockNumber = blockchain.BlockNumber(header.Address)
}
func (header *Header) GetNodeName(ctx app.Context, e app.Event) {
	header.NodeName = blockchain.NodeName(header.Address)
}
func (header *Header) GetNodeVersion(ctx app.Context, e app.Event) {
	if header.ChainId != blockchain.NodeVersion(header.Address) {
		header.ChainId = blockchain.NodeVersion(header.Address)
	} else {
		header.ChainId = "None"
	}
}

func main() {
	// Components routing:
	app.Route("/", &Header{})
	app.RunWhenOnBrowser()

	// HTTP routing:
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
