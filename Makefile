dep:
	go mod vendor

run: dep
	go run main.go

deploy:
	kubectl create -f httpserver.namespace.yaml
	kubectl create -f httpserver.deployment.yaml
	kubectl create -f httpserver.deployment2.yaml
	kubectl create -f httpserver.deployment3.yaml
	kubectl create -f httpserver.istio.yaml
