git config --global user.email "tristan@mrmxf.com"
git config --global user.name "mmTristan"

Get-Content $HOME/.netrc
Get-Content ~/.gitconfig

echo $HOME

go env

go env -w GOPRIVATE=gitlab.com/*,github.com/*


go env
git config --list

go get github.com/mmTristan/tpg-core-private
