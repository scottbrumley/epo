
## Set GOPATH
export GOPATH=/Users/sbrumley/IdeaProjects/epo/govendor

# Testing Syntax
./scripts/test.sh https://localhost:8443/remote UserName Password

# ePO Syntax Examples
## Find All Systems
go run epo.go -user=someuser -password=somepass -ignoressl=true -url=https://epohost:8443/remote -cmd=system.find -parms=searchText=. -output=json

## Find All Policies
go run epo.go -user=someuser -password=somepass -ignoressl=true -url=https://epohost:8443/remote -cmd=policy.find -parms=searchText=. -output=json

## Find All Repositories
go run epo.go -user=someuser -password=somepass -ignoressl=true -url=https://epohost:8443/remote -cmd=repository.find -parms=searchText=. -output=json

## Find All Queries
go run epo.go -user=someuser -password=somepass -ignoressl=true -url=https://epohost:8443/remote -cmd=core.listQueries  -output=json

## Find All Groups
go run epo.go -user=someuser -password=somepass -ignoressl=true -url=https://epohost:8443/remote -cmd=system.findGroups -output=json
