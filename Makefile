# Swagger-codegenのバージョン
SWAGGER_V="2.4.4"

.PHONY: all
all: build

.PHONY: clean
clean:
	rm bin/swagger-codegen-cli.jar
	rm -rf api-android/

.PHONY: build
build: build-android build-go

.PHONY: build-android
build-android: docs/swagger.yaml install
	java -jar ./bin/swagger-codegen-cli.jar generate -i docs/swagger.yaml -l java -o ./api-android -c android-config.json
	cd api-android; mvn package

.PHONY: build-go
build-go: docs/swagger.yaml install
	java -jar ./bin/swagger-codegen-cli.jar generate -i docs/swagger.yaml -l go-server -o ./api-go

# Install swagger-codegen-cli
# ヘルプ確認方法
# $ java -jar ./bin/swagger-codegen-cli.jar help
# 使えるconfigの確認方法
# $ java -jar ./bin/swagger-codegen-cli.jar config-help -l java
.PHONY: install
install: bin/swagger-codegen-cli.jar

bin/swagger-codegen-cli.jar:
	mkdir -p bin
	wget http://central.maven.org/maven2/io/swagger/swagger-codegen-cli/$(SWAGGER_V)/swagger-codegen-cli-$(SWAGGER_V).jar -O bin/swagger-codegen-cli.jar

