# usage> lint
# short> set up and run the golinter
# long>  set up and run the golinter as well as any necessary files
#   _      _          _
#  | |    (_)  _ _   | |_
#  | |__  | | | ' \  |  _|
#  |____| |_| |_||_|  \__|

# export the commit ID for the build

# ------------------------------------------------------------------------------
# -- sample go build (with versions)

# Import clog files
source clogrc/core/inc.sh
# Import lint filess
source clogrc/lintYamls.sh

# ------------------------------------------------------------------------------


# Check the linter is installed
if ! command -v golangci-lint --version &> /dev/null ; then
    fnInfo "golangci-lint not found, now installing ..."
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.1
    fnInfo "golangci-lint installed"
fi

# Check there is a configuration file
FILE=".golangci.yaml"
#if [ ! -f "$FILE" ]; then
    #create the yaml for the linter to run
    #ensure the spacing in the variable is preserved quotation marks
echo "$StandardYaml" >$PWD/$FILE
#fi

go_lint() {
pathname="$1"

if go_dir_fence $pathname; then 
    #update the lint to extract the failures
    #then just the number
    errors=$(golangci-lint run $pathname --enable gocognit| grep -E "\.go")
    count=$(echo "$errors"| grep -E -o "failures=\"[0-9]{1,}\">")

    if [ -z "${count}" ]; then
        fnOk "$pathname has no linting errors";
    else 

        #extract the total values here as an array
        NUMBER=($(echo "$count" | grep -o -E '[0-9]{1,}')) ; 
        #then loop through and add all the numbers
        sum="0"
        for i in "${NUMBER[@]}"
            do
            sum=$(($sum + $i))
        done
        
        errorfile="/linterErrors.xml"

        fnWarning "$pathname has $sum linting errors. The errors have compiled and saved in $pathname$errorfile";
        echo "$errors" >"$pathname$errorfile"
    fi       
fi
}

# walk_dir searches through every folder within a system and runs golint
# with every folder with a go file in it.
walk_dir () {  
    shopt -s nullglob dotglob

  

    for pathname in "$1"/*; do
        #Check if the file is a directory
        
        if [ -d "$pathname" ]; then
           #check it is a go folder
           go_lint $pathname
            if go_dir_fence $pathname; then 
                #update the lint to extract the failures
                #then just the number
                errors=$(golangci-lint run $pathname --enable gocognit| grep -E "\.go")
                count=$(echo "$errors"| grep -E -o "failures=\"[0-9]{1,}\">")
                if [ -z "${count}" ]; then
                    fnOk "$pathname has no linting errors";
                else 

                    #extract the total values here as an array
                    NUMBER=($(echo "$count" | grep -o -E '[0-9]{1,}')) ; 
                    #then loop through and add all the numbers
                    sum="0"
                    for i in "${NUMBER[@]}"
                        do
                        sum=$(($sum + $i))
                    done
                    
                    errorfile="/linterErrors.xml"

                    fnWarning "$pathname has $sum linting errors. The errors have compiled and saved in $pathname$errorfile";
                    echo "$errors" >"$pathname$errorfile"
                fi       
            fi
            # walk the path regardless
            # in search of nested go files
            walk_dir $pathname 
        fi
     #   echo "$pathname, file"
    done
}

go_dir_fence () {
    #check every file in the directory
    for goname in "$1"/*; do
       found=`echo $goname | grep -e "\.go$"`
        #found just gives the matching string
        if [ "$found" = "$goname" ]; then
            return 0
        fi
    done
    #return that no go files were indeed found
    return 1

}


fnInfo "beginning linting of $PWD"
# lint the working directory
go_lint $PWD
#then search the rest to be linted
walk_dir $PWD


