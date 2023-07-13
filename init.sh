#!/bin/bash

#Set GOPRIVATE to handle both types of git with the comma seperating them
go env -w GOPRIVATE=gitlab.com/*,github.com/*

#Add the two passwords for the git machines
read -r -d '' gitAccess << EOM
machine gitlab.com login mmTristan password $GITLAB_PAT

machine github.com login mmTristan password $GITHUB_PAT
EOM

#Generate the netrc
echo "$gitAccess">$HOME/.netrc

echo "$HOME/.netrc generated"


BRANCH=$(git branch | sed -n -e 's/^\* \(.*\)/\1/p') # extract current branch

go get gitlab.com/mmTristan/tpg-io@latest
go get github.com/mrmxf/opentsg-cote@$BRANCH
