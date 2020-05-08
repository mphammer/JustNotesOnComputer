import sys
import datetime
import re
import os
import fileinput

from id import *

def PrintFiles():
    home = os.path.expanduser("~")
    desktopPath = "{}/Desktop".format(home)
    DownloadsPath = "{}/Downloads".format(home)

    print("Desktop:")
    dirpath, dirnames, filenames = ("", "", "")
    for (dirpath, dirnames, filenames) in os.walk(desktopPath):
        break
    filenames.sort()
    for filename in filenames:
        if filename[0] == ".":
            continue
        # filename = filename.replace(" ", "\ ")
        print("\"{}/{}\"".format(dirpath, filename))

    print("")

    print("Downloads:")
    dirpath, dirnames, filenames = ("", "", "")
    for (dirpath, dirnames, filenames) in os.walk(DownloadsPath):
        break
    filenames.sort()
    for filename in filenames:
        if filename[0] == ".":
            continue
        # filename = filename.replace(" ", "\ ")
        print("\"{}/{}\"".format(dirpath, filename))
    

def Get(filePath, newName=""):
    id = getID()

    basename = os.path.basename(filePath)
    name, extension = os.path.splitext(basename)
    if newName != "":
        name = newName
    filename = "{}-{}{}".format(name, id, extension)
    newPath = "../_data/{}".format(filename)
    os.system("mv \"{}\" \"{}\"".format(filePath, newPath))
    print("![{}]({})".format(filename, newPath))

if __name__ == "__main__":
    if len(sys.argv) > 3:
        print("Usage: ./get [filepath] [new_name]")
        sys.exit(1)
    if sys.argv[1] == "":
        PrintFiles()
        sys.exit(0)
    Get(sys.argv[1], sys.argv[2])
    sys.exit(0)
    

    
    