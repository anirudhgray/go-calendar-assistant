package cmd

import (
	"fmt"
	"log"
	"time"

	"go-calendar-assistant/internal/calendar"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List upcoming calendar events",
	Long:  "Retrieve and display a list of upcoming events from your Google Calendar.",
	Run: func(cmd *cobra.Command, args []string) {
		srv, err := calendar.GetCalendarService()
		if err != nil {
			log.Fatalf("Unable to connect to Google Calendar: %v", err)
		}

		t := time.Now().Format(time.RFC3339)
		events, err := srv.Events.List("primary").ShowDeleted(false).
			SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
		if err != nil {
			log.Fatalf("Unable to retrieve events: %v", err)
		}

		fmt.Println("Upcoming events:")
		if len(events.Items) == 0 {
			fmt.Println("No upcoming events found.")
		} else {
			for _, item := range events.Items {
				date := item.Start.DateTime
				if date == "" {
					date = item.Start.Date
				}
				fmt.Printf("%s (%s)\n", item.Summary, date)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
