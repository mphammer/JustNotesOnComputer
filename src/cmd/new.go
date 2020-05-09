package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"SecondBrain/src/util"
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
	PreRun: func(cmd *cobra.Command, args []string) {
		LoadConfig()
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

func createGeneric(filename, templatePath string, patches map[string]string) error {
	path := Config.Project
	id := util.GetID()
	noteName := fmt.Sprintf("%s-%s.md", filename, id)
	notePath := fmt.Sprintf("%s/%s", path, noteName)

	// Get lines from template file
	lines, err := util.ReadFileLines(templatePath)
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
		for key, val := range patches {
			txt = strings.Replace(txt, key, val, -1)
		}
		txt = strings.Replace(txt, "TODO_ZETTLE_ID", id, -1)
		pathDepth := strings.Repeat("../", Config.ProjectDepth)
		txt = strings.Replace(txt, "TODO_FILE_PATH", fmt.Sprintf("%s%s", pathDepth, noteName), -1)
		fmt.Fprint(noteFile, txt)
	}

	noteFile.Close()

	fmt.Printf("%s\n", notePath)

	return nil
}

func createBookSummary() error {
	title := util.Input("Book Title: ")
	title = strings.Title(title)
	filename := strings.Replace(title, " ", "", -1)

	patches := map[string]string{
		"TODO_MAIN_TITLE":     title,
		"TODO_BOOK_REFERENCE": filename,
	}
	templatePath := "_templates/BookSummary.md"
	createGeneric(filename, templatePath, patches)

	return nil
}

func createContact() error {
	name := util.Input("Name: ")
	name = strings.Title(name)
	filename := strings.Replace(name, " ", "", -1)

	patches := map[string]string{
		"TODO_NAME": name,
		"TODO_TAG":  filename,
	}
	templatePath := "_templates/Contact.md"
	createGeneric(filename, templatePath, patches)

	return nil
}

func createJournal() error {
	t := time.Now()
	year := t.Year()
	month := t.Month()
	day := t.Day()
	dateString := fmt.Sprintf("%s %d %d", month, day, year)
	filename := "DailyNote"

	patches := map[string]string{
		"TODO_DATE": dateString,
	}
	templatePath := "_templates/Journal.md"
	createGeneric(filename, templatePath, patches)

	return nil
}

func createNote() error {
	noteTopic := util.Input("Note Topic: ")
	noteTopic = strings.Title(noteTopic)
	filename := strings.Replace(noteTopic, " ", "", -1)

	patches := map[string]string{
		"TODO_NOTE_TOPIC": noteTopic,
	}
	templatePath := "_templates/Note.md"
	createGeneric(filename, templatePath, patches)

	return nil
}
