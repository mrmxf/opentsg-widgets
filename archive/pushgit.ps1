# set up initial variables
$git_pat=$args[0]
$message=$args[1]


#what's my name
git config --global user.email "tristan@mrmxf.com"
git config --global user.name "mmTristan"

Get-ChildItem -Hidden
dir

# remove the original git
Remove-Item -Force -Recurse -Path .git
# make our new git

# make a temporary file to get the .git we want
# and replace the deleted one with it
mkdir tmp
cd tmp
git clone --no-checkout https://mmTristan:$git_pat@github.com/mmTristan/tpg-widgets-private.git
cd ..


Copy-Item -Path "tmp\tpg-widgets-private\.git" -Destination ".git" -Recurse

# remove the tmp folder
Remove-Item -Force -Recurse -Path tmp



# adding all files to be comitted
# which will create a merge issue.
git add -A
git commit -m "$message"

#assign the branch this can be changed later
git branch -M main
echo "commited our changes starting merge"

# Solver the merge issue by rebasing
git rebase origin/master 
# always take our updates, which is confusingly theirs as the update
git checkout --theirs .
echo "taking our updates"
# commit again
git add .
git remove .gitignore
echo "commit message: $message"
git commit -m "$message"

# continue the rebase now we've fixed the merge
git rebase --continue
echo "pushing to github.com/mmTristan/tpg-widgets-private.git"
# push those updates
git push https://mmTristan:$git_pat@github.com/mmTristan/tpg-widgets-private.git



