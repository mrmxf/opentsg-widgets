# Import clog files
source clogrc/privatepublic.sh
. clogrc/privatepublic.sh


# run those updates in all the go files
Walk_dir $PWD

# set up initial variables
git_pat=$1
message=$2

repo=https://mmTristan:$git_pat@github.com/mmTristan/tpg-widgets.git

# remove the original git
rm -R .git

# make our new git
# make a temporary file to get the .git we want
# and replace the deleted one with it

mkdir tmp
cd tmp
git clone --no-checkout $repo
ls -all
cd ..
cp -R "tmp/tpg-widgets/.git" ".git" 

# remove the tmp folder after exchanging gits
rm -R tmp


# adding all files to be comitted
# which will create a merge issue.
git add -A
git commit -m "$message"
#assign the branch this can be changed later
#echo $(git tag -l --format='%(contents)' ${GITHUB_REF_NAME})
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
# push those updates to the HEAD:master 
# because sometimes it makes a detached branch from a 3 way merge ?!
git push $repo HEAD:master


# this is for taking all of ours
#git checkout HEAD -- .

#remove the git when done so it doesn't inteerupt the next run
rm -R .git

