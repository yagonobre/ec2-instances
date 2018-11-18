#!/bin/bash

curl -Lo hack/instances.json https://www.ec2instances.info/instances.json

if [[ `git status --porcelain` ]]; then
	  git checkout master
    go generate

    git config user.name "Batata.Bot"
    git config user.email "batata.the.bot@gmail.com"

    git add instances_data.go
    git add hack/instances.json
    git commit -m "Update instances info"

		git fetch --tags
    VERSION=`git tag -l --sort -version:refname | head -n 1 | awk -F. '{$NF+=1; OFS="."; print $0}'| sed 's/ /./g'`
    git tag ${VERSION}

    git remote add github https://${GH_TOKEN}@${GH_REF} > /dev/null 2>&1
    git push github master
    git push github ${VERSION}
fi
