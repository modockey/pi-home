generate:
	java -jar openapi-generator-cli.jar generate \
	-i ../openapi.yaml \
	-g go-server \
	-o ./gen \
	--additional-properties sourceFolder=openapi \
	--additional-properties outputAsLibrary=true
	goimports -w ./gen/openapi/

swagger-generate:
	swagger generate server -t gen -f ../swagger.yaml