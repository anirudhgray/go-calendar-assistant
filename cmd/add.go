package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"go-calendar-assistant/internal/calendar"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new calendar event",
	Long:  "Create a new event in your Google Calendar with a title and date.",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter event summary: ")
		summary, _ := reader.ReadString('\n')
		summary = summary[:len(summary)-1]

		fmt.Print("Enter event date (YYYY-MM-DD): ")
		date, _ := reader.ReadString('\n')
		date = date[:len(date)-1]

		srv, err := calendar.GetCalendarService()
		if err != nil {
			log.Fatalf("Unable to connect to Google Calendar: %v", err)
		}

		event := calendar.CreateEvent(srv, summary, date)
		fmt.Printf("Event created: %s\n", event.HtmlLink)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
