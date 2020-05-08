import sys
import datetime
import re
import os
import fileinput

def Move(startPath, destPath):
    startPath = startPath.strip("/")
    destPath = destPath.strip("/")
    if os.path.isfile(destPath):
        print("ERROR: destPath must be a directory - cannot rename the file")

    dirpath, dirnames, filenames = ("", "", "")
    if os.path.isfile(startPath):
        dirpath = os.path.dirname(startPath)
        filenames = [os.path.basename(startPath)]
        dirnames = []
    else:
        for (dirpath, dirnames, filenames) in os.walk(startPath):
            break
    for filename in filenames:
        oldPath = "{}/{}".format(dirpath, filename)
        newPath = "{}/{}".format(destPath, filename)
        
        # Change all references to this file
        # print("find ../ -type f -name \"*\.md\" -print0 | xargs -0 sed -i '' -e 's~{}~{}~g'".format(oldPath[3:], newPath[3:]))
        os.system("find ../ -type f -name \"*\.md\" -print0 | xargs -0 sed -i '' -e 's~{}~{}~g'".format(oldPath[3:], newPath[3:]))

        # Change references within this file
        oldDepth = len(oldPath.split("/")) - 2
        oldRootPath = "../" * oldDepth
        newDepth = len(newPath.split("/")) - 2
        newRootPath = "../" * newDepth
        # print("replace: {} -> {}".format(oldRootPath, newRootPath))
        reading_file = open(oldPath, "r")
        new_file_content = ""
        regOldRootPath = oldRootPath.replace(".", "\.")
        for line in reading_file:
            new_line = re.sub(regOldRootPath, newRootPath, line)
            new_file_content += new_line
        reading_file.close()
        writing_file = open(oldPath, "w")
        writing_file.write(new_file_content)
        writing_file.close()

        # Create new directories
        # print("mkdir -p {}".format(destPath))
        os.system("mkdir -p {}".format(destPath))

        # Move the file
        # print("mv {} {}".format(oldPath, newPath))
        os.system("mv  {} {}".format(oldPath, newPath))

        # print("")
    for dirname in dirnames:
        Move("{}/{}".format(startPath,dirname), "{}/{}".format(destPath,dirname))

if __name__ == "__main__":
    if len(sys.argv) != 3:
        print("Usage: ./mv <file_or_directory_path> <directory_path>")
        sys.exit(1)
    startLoc = "../{}".format(sys.argv[1])
    destLoc = "../{}".format(sys.argv[2])
    Move(startLoc, destLoc)
    