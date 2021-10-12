package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	cc "github.com/fizzywhizbang/persona/ccgen"
	p "github.com/fizzywhizbang/persona/persongen"
	ss "github.com/fizzywhizbang/persona/ssngen"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

var ctype string
var qty = 1
var opt string

func main() {
	path, _ := os.Getwd()
	fmt.Println(path)
	gui()
}

// func gennames() {
// 	person := p.PersonGen(100)
// 	f, err := os.Create("first.txt")
// 	if err != nil {
// 		fmt.Println("Could not write file")
// 	}

// 	defer f.Close()
// 	payload := ""
// 	for i := 0; i < len(person); i++ {
// 		payload += "\"" + person[i].First + "\","
// 	}
// 	_, err2 := f.WriteString(payload)

// 	if err2 != nil {
// 		fmt.Println("Could not write to file")
// 	}

// }

//will return in next itteration
// func writeToFileFunc(filename string, payload string) {
// 	f, err := os.Create(filename)
// 	if err != nil {
// 		fmt.Println("Could not write file")
// 	}

// 	defer f.Close()

// 	_, err2 := f.WriteString(payload)

// 	if err2 != nil {
// 		fmt.Println("Could not write to file")
// 	}

// }

func gui() {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(620, 400)
	window.SetWindowTitle("Persona")

	centralWidget := widgets.NewQWidget(nil, 0)
	window.SetCentralWidget(centralWidget)

	verticalLayout := widgets.NewQVBoxLayout()

	toolbar := toolbarInit(app, window, widgets.NewQToolBar2(nil))

	verticalLayout.SetMenuBar(toolbar)
	centralWidget.SetLayout(verticalLayout)

	// make the window visible
	window.Show()

	app.Exec()
}

func creditCard(app *widgets.QApplication, window *widgets.QMainWindow) {
	mainWidget := widgets.NewQWidget(nil, 0)
	mainWidget.SetMinimumWidth(610)
	vLayout := widgets.NewQVBoxLayout()
	vLayout.SetContentsMargins(4, 0, 4, 0)
	mainWidget.SetLayout(vLayout)

	scrollArea := widgets.NewQScrollArea(window)

	scrollArea.SetHorizontalScrollBarPolicy(core.Qt__ScrollBarAlwaysOff)
	scrollArea.SetVerticalScrollBarPolicy(core.Qt__ScrollBarAlwaysOn)
	scrollArea.SetWidgetResizable(true)
	scrollArea.SetWidget(mainWidget)

	formLayout := widgets.NewQFormLayout(nil)

	toolbar := toolbarInit(app, window, widgets.NewQToolBar2(nil))
	vLayout.SetMenuBar(toolbar)
	person := p.PersonGen(qty)
	ssn := ss.SSNgen(qty)
	card := cc.GenerateCards(matchType(ctype), qty)
	for i := 0; i < qty; i++ {
		//card data
		pan := widgets.NewQLineEdit(nil)
		pan.SetText(card[i].Pan.Formatted)
		group0 := widgets.NewQHBoxLayout()
		cardLabel := widgets.NewQLabel(nil, 0)
		cardLabel.SetText("Card#")
		group0.AddWidget(cardLabel, 0, 0)
		group0.AddWidget(pan, 0, 0)
		formLayout.InsertRow6(0, group0)

		group1 := widgets.NewQHBoxLayout()

		explabel := widgets.NewQLabel(nil, 0)
		explabel.SetText("exp")
		group1.AddWidget(explabel, 0, 0)
		expiry := widgets.NewQLineEdit(nil)
		dateString := strconv.Itoa(card[i].ExpiryDate.Month) + "/" + strconv.Itoa(card[i].ExpiryDate.Year)
		expiry.SetText(dateString)
		group1.AddWidget(expiry, 0, 0)

		cvvlabel := widgets.NewQLabel(nil, 0)
		cvvlabel.SetText("CVV")
		group1.AddWidget(cvvlabel, 0, 0)
		cvv := widgets.NewQLineEdit(nil)
		cvv.SetText(card[i].CVV)

		group1.AddWidget(cvv, 0, 0)

		formLayout.InsertRow6(1, group1)

		///person

		x := 0
		for !ss.CheckValid(ssn[x]) { //loop until we have a valid SSN
			ssn = ss.SSNgen(qty)
		}
		fmt.Println(opt)
		if opt == "Persons Multiple Cards" {
			//allow to increment
			x = i
		}

		name := widgets.NewQLineEdit(nil)

		name.SetText(person[x].First + " " + strings.Title(strings.ToLower(person[x].Last)))

		formLayout.InsertRow5(2, name)

		group2 := widgets.NewQHBoxLayout()

		ssnLabel := widgets.NewQLabel(nil, 0)
		ssnLabel.SetText("SSN")
		group2.AddWidget(ssnLabel, 0, 0)
		ssnLineEdit := widgets.NewQLineEdit(nil)
		ssnLineEdit.SetText(ssn[x])

		group2.AddWidget(ssnLineEdit, 0, 0)

		formLayout.InsertRow6(3, group2)
		if qty > 1 {
			spacer := widgets.NewQLabel(nil, 0)
			spacer.SetText("###################################################################################")
			formLayout.InsertRow5(4, spacer)
		}
	}

	// person data

	vLayout.AddLayout(formLayout, 0)
	scroll_layout := widgets.NewQVBoxLayout2(nil)
	scroll_layout.AddWidget(scrollArea, 0, 0)
	scroll_layout.SetContentsMargins(0, 0, 0, 0)
	containerWidget := widgets.NewQWidget(nil, 0)
	containerWidget.SetLayout(scroll_layout)
	window.SetCentralWidget(containerWidget)
	window.Show()
}

func matchType(t string) string {
	var card cc.CardProperties
	for k, v := range cc.Bin {

		card = v
		if card.LongName == t {
			return k
		}

	}
	return ""
}

func toolbarInit(app *widgets.QApplication, window *widgets.QMainWindow, toolbar *widgets.QToolBar) *widgets.QToolBar {
	toolbar.SetToolButtonStyle(core.Qt__ToolButtonTextOnly)
	toolbar.SetMovable(true)
	var card cc.CardProperties
	selector := widgets.NewQComboBox(nil)
	items := []string{}
	for _, v := range cc.Bin {

		card = v
		items = append(items, card.LongName)
	}
	// selector := widgets.NewQComboBox(nil)
	sort.Strings(items) //sorting items before addind to selector
	selector.AddItems(items)
	if ctype != "" {
		selector.SetCurrentText(ctype)
	} else {
		selector.SetCurrentIndex(0)
	}

	toolbar.AddWidget(selector)
	qtyLabel := widgets.NewQLabel(nil, 0)
	qtyLabel.SetText("Qty")
	toolbar.AddWidget(qtyLabel)
	qtySelector := widgets.NewQComboBox(nil)
	qtySelector.AddItems([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20"})
	toolbar.AddWidget(qtySelector)
	if qty == 1 {
		qtySelector.SetCurrentIndex(0)
	} else {
		qtySelector.SetCurrentText(strconv.Itoa(qty))
	}

	option := widgets.NewQComboBox(nil)
	options := []string{"One Person Multiple Cards", "Persons Multiple Cards"}
	option.AddItems(options)
	if opt != "" {
		option.SetCurrentText(opt)
	} else {
		option.SetCurrentIndex(0)
	}
	toolbar.AddWidget(option)

	goButton := widgets.NewQPushButton2("Generate", nil)
	toolbar.AddWidget(goButton)
	goButton.ConnectClicked(func(checked bool) {
		ctype = selector.CurrentText()
		opt = option.CurrentText()
		q, err := strconv.Atoi(qtySelector.CurrentText())
		if err != nil { // if for some reason a number isn't entered default to one
			qty = 1
		} else {
			qty = q
		}

		creditCard(app, window)

	})
	return toolbar
}
