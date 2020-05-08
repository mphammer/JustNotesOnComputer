import sys
import datetime
import re

from id import *

def printHelp():
    print("b - Book Summary")
    print("c - Contact")
    print("z - Zettle")
    print("j - Journal")

def templateZettle(t):
    if t == "b":
        createBookSummary()
    elif t == "c":
        createContact()
    elif t == "z":
        createZettle()
    elif t == "j":
        createDailyJournal()
    else:
        print("ERROR: input '{}' is invalid".format(t))
        printHelp()
        return 

def createBookSummary():
    noteID = getID()
    main_title = raw_input("Book Title: ").title()
    fileName = main_title.replace(" ", "")
    with open("_BookSummaries.md") as f:
        noteName = "{}-{}.md".format(fileName, noteID)
        with open("../Zettles/{}".format(noteName), "w+") as f1:
            for line in f:
                line = re.sub("TODO_ZETTLE_ID", noteID, line,)
                line = re.sub("TODO_FILENAME", noteName, line)
                line = re.sub("TODO_MAIN_TITLE", main_title, line)
                line = re.sub("TODO_BOOK_REFERENCE", fileName, line)
                f1.write(line)
        print("Note: Zettles/{}".format(noteName))

def createContact():
    noteID = getID()
    name = raw_input("Name: ").title()
    nameTag = name.replace(" ", "")
    with open("_Contacts.md") as f:
        noteName = "{}-{}.md".format(nameTag, noteID)
        with open("../Zettles/{}".format(noteName), "w+") as f1:
            for line in f:
                line = re.sub("TODO_ZETTLE_ID", noteID, line)
                line = re.sub("TODO_FILENAME", noteName, line)
                line = re.sub("TODO_NAME", name, line)
                line = re.sub("TODO_TAG", nameTag, line)
                f1.write(line)
        print("Note: Zettles/{}".format(noteName))

def createDailyJournal():
    noteID = getID()
    x = datetime.datetime.now()
    dateString = x.strftime("%B %d %Y")
    with open("_DailyNotes.md") as f:
        noteName = "DailyNote-{}.md".format(noteID)
        with open("../Zettles/{}".format(noteName), "w+") as f1:
            for line in f:
                line = re.sub("TODO_DATE", dateString, line)
                line = re.sub("TODO_ZETTLE_ID", noteID, line)
                line = re.sub("TODO_FILENAME", noteName, line)
                f1.write(line)
        print("Note: Zettles/{}".format(noteName))

def createZettle():
    noteID = getID()
    x = datetime.datetime.now()
    dateString = x.strftime("%B %d %Y")
    with open("_Zettles.md") as f:
        noteTopic = raw_input("Note Topic: ").title()
        fileName = noteTopic.replace(" ", "")
        noteName = "{}-{}.md".format(fileName, noteID)
        with open("../Zettles/{}".format(noteName), "w+") as f1:
            for line in f:
                line = re.sub("TODO_NOTE_TOPIC", noteTopic, line)
                line = re.sub("TODO_ZETTLE_ID", noteID, line)
                line = re.sub("TODO_FILENAME", noteName, line)
                f1.write(line)
        print("Note: Zettles/{}".format(noteName))

if __name__ == "__main__":
    if len(sys.argv) <= 1:
        print("ERROR: Must provide argument")
        printHelp()
    else:
        zettleType = sys.argv[1]
        templateZettle(zettleType)