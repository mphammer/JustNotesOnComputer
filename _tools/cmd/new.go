package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"SecondBrain/_tools/util"
)

func init() {
	rootCmd.AddCommand(newCmd)
}

var newCmd = &cobra.Command{
	Use:     "new NOTE_TYPE",
	Aliases: []string{"n"},
	Short:   "Create a new note",
	Long: `Create a new note:
		   Note Types:
		   - booksummary, bs 
		   - contact, c
		   - journal, j  
		   - note, n`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("this command takes 1 argument")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		noteType := args[0]
		switch noteType {
		case "booksummary", "bs":
			createBookSummary()
		case "contact", "c":
			createContact()
		case "journal", "j":
			createJournal()
		case "note", "n":
			createNote()
		default:
			return fmt.Errorf("'%s' is not a valid note type", noteType)
		}
		return nil
	},
}

func createBookSummary() error {
	id := util.GetID()
	title := util.Input("Book Title: ")
	filename := strings.Replace(title, " ", "", -1)
	noteName := fmt.Sprintf("%s-%s.md", filename, id)
	path := "../Zettles"
	notePath := fmt.Sprintf("%s/%s", path, noteName)

	// Get lines from template file
	lines, err := util.ReadFileLines("templates/_BookSummaries.md")
	if err != nil {
		return fmt.Errorf("failed to read template: %+v", err)
	}

	// Create the new note
	noteFile, err := os.Create(notePath)
	if err != nil {
		return fmt.Errorf("failed create file: %s", err)
	}

	// Write lines to note
	for _, txt := range lines {
		txt = strings.Replace(txt, "TODO_ZETTLE_ID", id, -1)
		txt = strings.Replace(txt, "TODO_FILENAME", noteName, -1)
		txt = strings.Replace(txt, "TODO_MAIN_TITLE", title, -1)
		txt = strings.Replace(txt, "TODO_BOOK_REFERENCE", filename, -1)
		fmt.Fprintln(noteFile, txt)
	}

	noteFile.Close()

	fmt.Printf("%s\n", notePath)

	return nil
}

func createContact() error {
	id := util.GetID()
	name := util.Input("Name: ")
	nameTag := strings.Replace(name, " ", "", -1)
	noteName := fmt.Sprintf("%s-%s.md", nameTag, id)
	path := "../Zettles"
	notePath := fmt.Sprintf("%s/%s", path, noteName)

	// Get lines from template file
	lines, err := util.ReadFileLines("templates/_Contacts.md")
	if err != nil {
		return fmt.Errorf("failed to read template: %+v", err)
	}

	// Create the new note
	noteFile, err := os.Create(notePath)
	if err != nil {
		return fmt.Errorf("failed create file: %s", err)
	}

	// Write lines to note
	for _, txt := range lines {
		txt = strings.Replace(txt, "TODO_ZETTLE_ID", id, -1)
		txt = strings.Replace(txt, "TODO_FILENAME", noteName, -1)
		txt = strings.Replace(txt, "TODO_NAME", name, -1)
		txt = strings.Replace(txt, "TODO_TAG", nameTag, -1)
		fmt.Fprintln(noteFile, txt)
	}

	noteFile.Close()

	fmt.Printf("%s\n", notePath)

	return nil
}

func createJournal() error {
	id := util.GetID()
	t := time.Now()
	year := t.Year()
	month := t.Month()
	day := t.Day()
	dateString := fmt.Sprintf("%s %d %d", month, day, year)
	noteName := fmt.Sprintf("DailyNote-%s.md", id)
	path := "../Zettles"
	notePath := fmt.Sprintf("%s/%s", path, noteName)

	// Get lines from template file
	lines, err := util.ReadFileLines("templates/_DailyNotes.md")
	if err != nil {
		return fmt.Errorf("failed to read template: %+v", err)
	}

	// Create the new note
	noteFile, err := os.Create(notePath)
	if err != nil {
		return fmt.Errorf("failed create file: %s", err)
	}

	// Write lines to note
	for _, txt := range lines {
		txt = strings.Replace(txt, "TODO_DATE", dateString, -1)
		txt = strings.Replace(txt, "TODO_ZETTLE_ID", id, -1)
		txt = strings.Replace(txt, "TODO_FILENAME", noteName, -1)
		fmt.Fprintln(noteFile, txt)
	}

	noteFile.Close()

	fmt.Printf("%s\n", notePath)

	return nil
}

func createNote() error {
	id := util.GetID()
	noteTopic := util.Input("Note Topic: ")
	fileName := strings.Replace(noteTopic, " ", "", -1)
	noteName := fmt.Sprintf("%s-%s.md", fileName, id)
	path := "../Zettles"
	notePath := fmt.Sprintf("%s/%s", path, noteName)

	// Get lines from template file
	lines, err := util.ReadFileLines("templates/_Zettles.md")
	if err != nil {
		return fmt.Errorf("failed to read template: %+v", err)
	}

	// Create the new note
	noteFile, err := os.Create(notePath)
	if err != nil {
		return fmt.Errorf("failed create file: %s", err)
	}

	// Write lines to note
	for _, txt := range lines {
		txt = strings.Replace(txt, "TODO_NOTE_TOPIC", noteTopic, -1)
		txt = strings.Replace(txt, "TODO_ZETTLE_ID", id, -1)
		txt = strings.Replace(txt, "TODO_FILENAME", noteName, -1)
		fmt.Fprintln(noteFile, txt)
	}

	noteFile.Close()

	fmt.Printf("%s\n", notePath)

	return nil
}
