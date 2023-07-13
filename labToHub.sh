# set up initial variables
git_pat=$1
message=$2
target="main"


# Import clog files
source clogrc/privatepublic.sh
. clogrc/privatepublic.sh

#update the names to reflect the new repo
# remove the old mod to stop changes
rm go.mod

Walk_dir "$PWD" "github.com/mrmxf/opentsg-widgets" "github.com/mmTristan/tpg-widgets-private"
Walk_dir "$PWD" "github.com/mrmxf/opentsg-cote" "github.com/mmTristan/tpg-core-private"
Walk_dir "$PWD" "gitlab.com/mmTristan/tpg-io" "github.com/mmTristan/tpg-io-private"


# make the new mod with the new name
go mod init "github.com/mmTristan/tpg-widgets-private"
go mod tidy

# change these names to be whatever config
git config --global user.email "tristan@mrmxf.com"
git config --global user.name "mmTristan"

FULL=$(git log --format="%H" -n 1)
Branch=$(echo $FULL | head -c 38)


# make a temporary file to get the .git we want
# and replace the deleted one with it
mkdir tmp
cd tmp

if ! git clone --no-checkout https://mm-sandbox:$git_pat@github.com/mm-sandbox/tpg-widgets-private.git ; then
	exit 1
fi



cd ..
# remove the original git
rm -R .git
# replace with the github target
cp -R "tmp/tpg-widgets-private/.git" ".git"
# remove the tmp folder
rm -R tmp


# add the branch to the hosted forked repo
git add .
# remove uneeded pipeline files
git rm labToHub.sh -f
git rm .gitlab-ci.yml -f
git commit -m "$message"
git branch -M "$Branch"

if ! git push https://mm-sandbox:$git_pat@github.com/mm-sandbox/tpg-widgets-private.git ; then
    echo "error pushing update to mm-sandbox/tpg-widgets-private.git"
	exit 1
fi

#set authorisation here
export GH_TOKEN="$git_pat"
# pull request to the new library
if ! gh pr create --repo mmTristan/tpg-widgets-private  --title "$message" --body "$message" --head "mm-sandbox:$Branch" --base $target ; then
    export GH_TOKEN=""
    echo "error generating pull request to mmTristan/tpg-widgets-private.git"
	exit 1
fi
export GH_TOKEN=""

# remove the replaced git
rm -R .git