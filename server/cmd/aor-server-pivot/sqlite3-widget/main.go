package main

import (
	"fmt"
	"time"

	"github.com/ghetzel/pivot"
	"github.com/ghetzel/pivot/dal"
	"github.com/ghetzel/pivot/mapper"
)

var Widgets mapper.Mapper

var WidgetsSchema = &dal.Collection{
	Name:                   `widgets`,
	IdentityFieldType:      dal.StringType,
	IdentityFieldFormatter: dal.GenerateUUID,
	Fields: []dal.Field{
		{
			Name:        `type`,
			Description: `The type of widget.`,
			Type:        dal.StringType,
			Validator:   dal.ValidateIsOneOf(`foo`, `bar`, `baz`),
			Required:    true,
		}, {
			Name:        `usage`,
			Description: `Short description on how to use this widget.`,
			Type:        dal.StringType,
		}, {
			Name:        `created_at`,
			Description: `When the widget was created.`,
			Type:        dal.TimeType,
			Formatter:   dal.CurrentTimeIfUnset,
		}, {
			Name:        `updated_at`,
			Description: `Last time the widget was updated.`,
			Type:        dal.TimeType,
			Formatter:   dal.CurrentTime,
		},
	},
}

type Widget struct {
	ID        string    `pivot:"id,identity"`
	Type      string    `pivot:"type"`
	Usage     string    `pivot:"usage"`
	CreatedAt time.Time `pivot:"created_at"`
	UpdatedAt time.Time `pivot:"updated_at"`
}

func main() {
	// setup a new backend instance based on the supplied connection string
	if backend, err := pivot.NewDatabase(`sqlite:///./test.db`); err == nil {

		// initialize the backend (connect to/open it)
		if err := backend.Initialize(); err == nil {

			// register models to this database backend
			Widgets = mapper.NewModel(backend, WidgetsSchema)

			// create the model tables if they don't exist
			if err := Widgets.Migrate(); err != nil {
				fmt.Printf("failed to create widget table: %v\n", err)
				return
			}

			// make a new Widget instance, containing the data we want to see
			// the ID field will be populated after creation with the auto-
			// generated UUID.
			newWidget := Widget{
				Type:  `foo`,
				Usage: `A fooable widget.`,
			}

			// insert a widget (ID will be auto-generated because of dal.GenerateUUID)
			if err := Widgets.Create(&newWidget); err != nil {
				fmt.Printf("failed to insert widget: %v\n", err)
				return
			}

			// retrieve the widget using the ID we just got back
			var gotWidget Widget

			if err := Widgets.Get(newWidget.ID, &gotWidget); err != nil {
				fmt.Printf("failed to retrieve widget: %v\n", err)
				return
			}

			fmt.Printf("Got Widget: %#+v", gotWidget)

			// delete the widget
			if err := Widgets.Delete(newWidget.ID); err != nil {
				fmt.Printf("failed to delete widget: %v\n", err)
				return
			}
		} else {
			fmt.Printf("failed to initialize backend: %v\n", err)
			return
		}
	} else {
		fmt.Printf("failed to create backend: %v\n", err)
		return
	}

}
