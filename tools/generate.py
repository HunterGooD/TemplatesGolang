import os
import sys

def scan(path):
    with os.scandir(path) as it:
        for entry in it:
            splited = entry.name.split(".")
            if not entry.name.startswith('.'):
                if entry.is_dir():
                    scan(path+"/"+entry.name)
                elif entry.is_file and len(splited) == 2 and splited[1] == "go" or splited[1] == "mod":
                    replaceInFile(path+"/"+entry.name,"templatego.test", sys.argv[1])

def replaceInFile(nameFile, oldString, newString):
    with open(nameFile, "r") as w:
        data = w.read()
    with open(nameFile, "w") as w:
        w.write(data.replace(oldString, newString))

if len(sys.argv) != 2:
    print("Не передан аргумент нового модуля, "+ str(len(sys.argv)) + " Args: ",sys.argv)
    exit(1)

scan("./")