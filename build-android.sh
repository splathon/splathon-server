##!/usr/bin/env bash
# Swagger-codegenのバージョン
SWAGGER_V="2.4.4"

# 環境をクリアする
rm swagger-codegen-cli.jar
rm -rf api-android/

# swagger-codegenをmavenから取得
wget http://central.maven.org/maven2/io/swagger/swagger-codegen-cli/$SWAGGER_V/swagger-codegen-cli-$SWAGGER_V.jar -O swagger-codegen-cli.jar

# swagger.yamlからAPIライブラリを作成
# 以下、各ビルド用にカスタマイズすること(これはAndroid用)
java -jar swagger-codegen-cli.jar generate -i docs/swagger.yaml -l java -o ./api-android -c android-config.json

# ヘルプ確認方法
#java -jar swagger-codegen-cli.jar help

# 使えるconfigの確認方法
#java -jar swagger-codegen-cli.jar config-help -l java

# swagger-codegenを削除
rm swagger-codegen-cli.jar

###### Swagger固有の処理はここまで ######

# 以下、android用のライブラリ作成用コマンド
cd api-android
mvn package
