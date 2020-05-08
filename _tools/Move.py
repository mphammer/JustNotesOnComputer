import sys
import datetime
import re
import os


def Move(startPath, destPath):
    startPath = startPath.strip("/")
    destPath = destPath.strip("/")

    isFile = os.path.isfile(startPath)
    dirpath, dirnames, filenames = ("", "", "")
    if isFile:
        print("siFile")
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
        print("find ../ -type f -name \"*\.md\" -print0 | xargs -0 sed -i '' -e 's~{}~{}~g'".format(oldPath[3:], newPath[3:]))
        # os.system("find ../ -type f -name \"*\.md\" -print0 | xargs -0 sed -i '' -e 's~{}~{}~g'".format(oldPath[3:], newPath[3:]))

        # Change references within this file
        oldDepth = len(oldPath.split("/")) - 2
        oldRootPath = "../" * oldDepth
        newDepth = len(newPath.split("/")) - 2
        newRootPath = "../" * newDepth
        print("sed -i '' -e 's~{}~{}~g' {}".format(oldRootPath, newRootPath, oldPath))
        os.system("sed -i '' -e 's~{}~{}~g' {}".format(oldRootPath, newRootPath, oldPath))

        # Create new directories
        print("mkdir -p {}".format(destPath))
        os.system("mkdir -p {}".format(destPath))

        # Move the file
        print("mv {} {}".format(oldPath, newPath))
        os.system("mv  {} {}".format(oldPath, newPath))

        print("")
    for dirname in dirnames:
        Move("{}/{}".format(startPath,dirname), "{}/{}".format(destPath,dirname))

if __name__ == "__main__":
    startLoc = "../{}".format(sys.argv[1])
    destLoc = "../{}".format(sys.argv[2])
    Move(startLoc, destLoc)
    