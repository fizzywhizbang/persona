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
	qt "github.com/mappu/miqt/qt6"
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
	app := qt.NewQApplication(os.Args)

	window := qt.NewQMainWindow(nil)
	window.SetMinimumSize2(620, 400)
	window.SetWindowTitle("Persona")

	centralWidget := qt.NewQWidget(nil)
	window.SetCentralWidget(centralWidget)

	verticalLayout := qt.NewQVBoxLayout(nil)

	toolbar := toolbarInit(app, window, qt.NewQToolBar3())

	verticalLayout.SetMenuBar(toolbar.QWidget)
	centralWidget.SetLayout(verticalLayout.QLayout)

	// make the window visible
	window.Show()

	qt.QApplication_Exec()
}

func creditCard(app *qt.QApplication, window *qt.QMainWindow) {
	mainWidget := qt.NewQWidget(nil)
	mainWidget.SetMinimumWidth(610)
	vLayout := qt.NewQVBoxLayout(nil)
	vLayout.SetContentsMargins(4, 0, 4, 0)
	mainWidget.SetLayout(vLayout.QLayout)

	scrollArea := qt.NewQScrollArea(window.QWidget)

	scrollArea.SetHorizontalScrollBarPolicy(qt.ScrollBarAlwaysOff)
	scrollArea.SetVerticalScrollBarPolicy(qt.ScrollBarAlwaysOn)
	scrollArea.SetWidgetResizable(true)
	scrollArea.SetWidget(mainWidget)

	formLayout := qt.NewQFormLayout(nil)

	toolbar := toolbarInit(app, window, qt.NewQToolBar3())
	vLayout.SetMenuBar(toolbar.QWidget)
	person := p.PersonGen(qty)
	ssn := ss.SSNgen(qty)
	card := cc.GenerateCards(matchType(ctype), qty)
	for i := 0; i < qty; i++ {
		//card data
		pan := qt.NewQLineEdit(nil)
		pan.SetText(card[i].Pan.Formatted)
		group0 := qt.NewQHBoxLayout(nil)
		cardLabel := qt.NewQLabel(nil)
		cardLabel.SetText("Card#")
		group0.AddWidget(cardLabel.QWidget)
		group0.AddWidget(pan.QWidget)
		formLayout.InsertRow6(0, group0.QLayout)

		group1 := qt.NewQHBoxLayout(nil)

		explabel := qt.NewQLabel(nil)
		explabel.SetText("exp")
		group1.AddWidget(explabel.QWidget)
		expiry := qt.NewQLineEdit(nil)
		dateString := strconv.Itoa(card[i].ExpiryDate.Month) + "/" + strconv.Itoa(card[i].ExpiryDate.Year)
		expiry.SetText(dateString)
		group1.AddWidget(expiry.QWidget)

		cvvlabel := qt.NewQLabel(nil)
		cvvlabel.SetText("CVV")
		group1.AddWidget(cvvlabel.QWidget)
		cvv := qt.NewQLineEdit(nil)
		cvv.SetText(card[i].CVV)

		group1.AddWidget(cvv.QWidget)

		formLayout.InsertRow6(1, group1.QLayout)

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

		name := qt.NewQLineEdit(nil)

		name.SetText(person[x].First + " " + strings.Title(strings.ToLower(person[x].Last)))

		formLayout.InsertRow5(2, name.QWidget)

		group2 := qt.NewQHBoxLayout(nil)

		ssnLabel := qt.NewQLabel(nil)
		ssnLabel.SetText("SSN")
		group2.AddWidget(ssnLabel.QWidget)
		ssnLineEdit := qt.NewQLineEdit(nil)
		ssnLineEdit.SetText(ssn[x])

		group2.AddWidget(ssnLineEdit.QWidget)

		formLayout.InsertRow6(3, group2.QLayout)
		if qty > 1 {
			spacer := qt.NewQLabel(nil)
			spacer.SetText("###################################################################################")
			formLayout.InsertRow5(4, spacer.QWidget)
		}
	}

	// person data

	vLayout.AddLayout(formLayout.QLayout)
	scrollLayout := qt.NewQVBoxLayout2()
	scrollLayout.AddWidget(scrollArea.QWidget)
	scrollLayout.SetContentsMargins(0, 0, 0, 0)
	containerWidget := qt.NewQWidget(nil)
	containerWidget.SetLayout(scrollLayout.QLayout)
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

func toolbarInit(app *qt.QApplication, window *qt.QMainWindow, toolbar *qt.QToolBar) *qt.QToolBar {
	toolbar.SetToolButtonStyle(qt.ToolButtonTextOnly)
	toolbar.SetMovable(true)
	var card cc.CardProperties
	selector := qt.NewQComboBox(nil)
	items := []string{}
	for _, v := range cc.Bin {

		card = v
		items = append(items, card.LongName)
	}
	// selector := qt.NewQComboBox(nil)
	sort.Strings(items) //sorting items before addind to selector
	selector.AddItems(items)
	if ctype != "" {
		selector.SetCurrentText(ctype)
	} else {
		selector.SetCurrentIndex(0)
	}

	toolbar.AddWidget(selector.QWidget)
	qtyLabel := qt.NewQLabel(nil)
	qtyLabel.SetText("Qty")
	toolbar.AddWidget(qtyLabel.QWidget)
	qtySelector := qt.NewQComboBox(nil)
	qtySelector.AddItems([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20"})
	toolbar.AddWidget(qtySelector.QWidget)
	if qty == 1 {
		qtySelector.SetCurrentIndex(0)
	} else {
		qtySelector.SetCurrentText(strconv.Itoa(qty))
	}

	option := qt.NewQComboBox(nil)
	options := []string{"One Person Multiple Cards", "Persons Multiple Cards"}
	option.AddItems(options)
	if opt != "" {
		option.SetCurrentText(opt)
	} else {
		option.SetCurrentIndex(0)
	}
	toolbar.AddWidget(option.QWidget)

	goButton := qt.NewQPushButton3("Generate")
	toolbar.AddWidget(goButton.QWidget)
	goButton.OnClicked(func() {
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
