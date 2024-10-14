package main

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Contact struct {
	firstName   string
	lastName    string
	email       string
	phoneNumber string
	state       string
	business    bool
}

var contacts []Contact

var app = tview.NewApplication()

var pages = tview.NewPages()

var text = tview.NewTextView().
	SetTextColor(tcell.ColorGreen).
	SetText("(a) to add a new contact\n(q) to quit")

var contactList = tview.NewList().ShowSecondaryText(false)

func addContactList() {
	contactList.Clear()
	for index, contact := range contacts {
		contactList.AddItem(contact.firstName+" "+contact.lastName, "", rune(49+index), nil)
	}
}

var flex = tview.NewFlex()

var contactText = tview.NewTextView()

func setContactText(contact *Contact) {
	contactText.Clear()
	contactText.SetText(
		"First Name: " + contact.firstName + "\n" +
			"Last Name: " + contact.lastName + "\n" +
			"Email: " + contact.email + "\n" +
			"Phone Number: " + contact.phoneNumber + "\n" +
			"State: " + contact.state + "\n" +
			"Business: " + strconv.FormatBool(contact.business) + "\n",
	)
}

var states = []string{"Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Delaware"}
var form = tview.NewForm()

func addContactForm() {
	contact := Contact{}

	form.AddInputField("First Name", "", 20, nil, func(s string) { contact.firstName = s })

	form.AddInputField("Last Name", "", 20, nil, func(s string) { contact.lastName = s })

	form.AddInputField("Email", "", 20, nil, func(s string) { contact.email = s })

	form.AddInputField("Phone Number", "", 20, nil, func(s string) { contact.phoneNumber = s })

	form.AddDropDown("State", states, 0, func(s string, index int) { contact.state = s })

	form.AddCheckbox("Business", false, func(b bool) { contact.business = b })

	form.AddButton("Save", func() {
		contacts = append(contacts, contact)
		addContactList()
		pages.SwitchToPage("Menu")
	})
}

func main() {
	flex.SetDirection(tview.FlexRow).
		AddItem(
			tview.NewFlex().
				AddItem(contactList, 0, 1, true).
				AddItem(contactText, 0, 4, false),
			0, 6, true,
		).
		AddItem(text, 0, 1, false)

	contactList.SetChangedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		setContactText(&contacts[index])
	})

	pages.AddPage("Menu", flex, true, true)
	pages.AddPage("Add Contact", form, true, false)

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' {
			app.Stop()
		} else if event.Rune() == 'a' {
			form.Clear(true)
			addContactForm()
			pages.SwitchToPage("Add Contact")
		}
		return event
	})

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
