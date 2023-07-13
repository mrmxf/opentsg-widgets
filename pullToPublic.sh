# Import clog files
source clogrc/privatepublic.sh
. clogrc/privatepublic.sh

#set variables for gitlab and github for swappping between them

# run those updates in all the go files changing from
Walk_dir $PWD "github.com/mmTristan/tpg-widgets-private" "github.com/mmTristan/tpg-widgets"
Walk_dir "$PWD" "github.com/mrmxf/opentsg-cote-private" "github.com/mmTristan/tpg-core"
Walk_dir "$PWD" "gitlab.com/mmTristan/tpg-io-private" "github.com/mmTristan/tpg-io"


# set up initial variables
git_pat=$1
message=$2
target="main"

repo=https://mm-sandbox:$git_pat@github.com/mm-sandbox/tpg-widgets.git

# set up stuff here to clone the repos history etc


# remove the original git
rm -R .git

# make our new git
# make a temporary file to get the .git we want
# and replace the deleted one with it

mkdir tmp
cd tmp

if ! git clone --no-checkout $repo ; then
    exit 1
fi

ls -all
cd ..
cp -R "tmp/tpg-widgets/.git" ".git"

# remove the tmp folder after exchanging gits
rm -R tmp

# adding all files to be comitted
git add .
# not including files relating to this workflow
git rm pullToPublic.sh -f
git rm .github/* -f
git commit -m "$message"

FULL=$(git log --format="%H" -n 1)
Branch=$(echo $FULL | head -c 38)


#assign the branch this can be changed later
#echo $(git tag -l --format='%(contents)' ${GITHUB_REF_NAME})
git branch -M $Branch

export GH_TOKEN="$git_pat"
#create a new repo of the thing
# gh repo fork --remote-name $Branch

if ! git push $repo; then
    exit 1
fi

if ! gh pr create --repo mmTristan/tpg-widgets  --title "$message" --body "$message" --head "mm-sandbox:$Branch" --base $target ; then
    export GH_TOKEN=""
    exit 1
fi

export GH_TOKEN=""

# pull request to the new library
#gh pr create --title "here is $message" --body "message of $message" --base main --head "mm-sandbox/tpg-widgets/$Branch" --project mmTristan/tpg-widgets
