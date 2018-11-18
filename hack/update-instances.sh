#!bin/bash

curl -Lo hack/instances.json https://www.ec2instances.info/instances.json

if ! git diff-index --quiet HEAD --; then
    go generate
		git add instances_data.go
		git add hack/instances.json
		git commit -m "Update instances info"
		# bump tag
		git push
fi
