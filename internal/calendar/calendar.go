package calendar

import (
	"log"

	"google.golang.org/api/calendar/v3"
)

func CreateEvent(srv *calendar.Service, summary, date string) *calendar.Event {
	event := &calendar.Event{
		Summary: summary,
		Start: &calendar.EventDateTime{
			Date: date,
		},
		End: &calendar.EventDateTime{
			Date: date,
		},
	}

	newEvent, err := srv.Events.Insert("primary", event).Do()
	if err != nil {
		log.Fatalf("Unable to create event: %v", err)
	}
	return newEvent
}
