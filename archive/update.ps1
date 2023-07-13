# set up initial variables
$git_pat=$args[0]
$message=$args[1]

# remove the original git
Remove-Item -Force -Recurse -Path .git
# make our new git

# make a temporary file to get the .git we want
# and replace the deleted one with it
mkdir tmp
cd tmp
git clone --no-checkout https://mmTristan:$git_pat@github.com/mmTristan/tpg-widgets.git
cd ..
Copy-Item -Path "tmp\tpg-widgets\.git" -Destination ".git" -Recurse

# remove the tmp folder
Remove-Item -Force -Recurse -Path tmp


# adding all files to be comitted
# which will create a merge issue.
git add -A
git commit -m "$message"
#assign the branch this can be changed later
git branch -M master


# Solver the merge issue by rebasing
git rebase origin/master 
# always take our updates, which is confusingly theirs as the update
git checkout --theirs .
# commit again
git add .
echo "commit message: $message"
git commit -m "$message"

# continue the rebase now we've fixed the merge
git rebase --continue
echo "pushing to github.com/mmTristan/tpg-widgets.git"
# push those updates
git push https://mmTristan:$git_pat@github.com/mmTristan/tpg-widgets.git

# this is for taking all of ours
#git checkout HEAD -- .

#git rebase origin/master 
# the above gives all the commebted code
# git add .   
#git rebase --continue

