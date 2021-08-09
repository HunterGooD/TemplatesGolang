#!/bin/bash

# $1 - path directory
# $2 - new mod name
function scan {
    for file in `find $1/* -type f -name "*.go"`
        do
            echo "Замена в файле $file..."
            replaceInFile $file "templatego.test" $2
            echo "$file Закончено"
    done
}

# $1 - path to file 
# $2 - oldString 
# $3 - newString
function replaceInFile {
    sed -i "s!$2!$3!" $1
}

if [ -z "$1" ]
    then
        # scan .
        echo "Не передан аргумент нового модуля"
    else
        newName=${1@Q}
        echo $newName
        scan . $newName
        echo "Замена в файле ./go.mod..."
        replaceInFile ./go.mod "templatego.test" $newName
        echo "./go.mod Закончено"
fi
