# set up initial variables
$git_pat=$args[0]
$message=$args[1]
$target=$args[2]




# change these names to be whatever config
git config --global user.email "tristan@mrmxf.com"
git config --global user.name "mmTristan"

$FULL = git log --format="%H" -n 1

$Branch = $FULL.SubString(0,38)



# make a temporary file to get the .git we want
# and replace the deleted one with it
mkdir tmp
cd tmp

git clone --no-checkout https://mm-sandbox:$git_pat@github.com/mm-sandbox/tpg-widgets-private.git
if (-Not $?) {
	exit 1
}



cd ..
# remove the original git
Remove-Item -Force -Recurse -Path .git
# replace with the github target
Copy-Item -Path "tmp\tpg-widgets-private\.git" -Destination ".git" -Recurse
# remove the tmp folder
Remove-Item -Force -Recurse -Path tmp


# add the branch to the hosted forked repo
# add all changes apart from files that are needed to be transferred
git add . 
git rm pullgit.sh

git commit -m "$message"
git branch -M "$Branch"
git push https://mm-sandbox:$git_pat@github.com/mm-sandbox/tpg-widgets-private.git

if (-Not $?) {
    echo "error pushing update to mm-sandbox/tpg-widgets-private.git"
	exit 1
}

#set authorisation here
$env:GH_TOKEN="$git_pat"
# pull request to the new library
gh pr create --repo mmTristan/tpg-widgets-private  --title "$message" --body "$message" --head "mm-sandbox:$Branch" --base $target

#reset env variable




if (-Not $?) {
    $env:GH_TOKEN=""
    echo "error generating pull request to mmTristan/tpg-widgets-private.git"
	exit 1
}
$env:GH_TOKEN=""

# remove the replaced git
Remove-Item -Force -Recurse -Path .git