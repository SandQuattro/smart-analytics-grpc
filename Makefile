include vendor.proto.mk

.bin-deps:
	$(info Installing binary dependencies...)

	go install github.com/bufbuild/buf/cmd/buf@v1.41.0
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Форматирование .proto файлов с помощью buf
.buf-format:
	$(info run buf format...)
	buf format -w

# Линтинг .proto файлов с помощью buf
.buf-lint:
	$(info run buf lint...)
	buf lint

# Генерация .pb файлов с помощью buf
.buf-generate:
	$(info run buf generate...)
	buf generate

.tidy:
	go mod tidy

# Генерация .pb файлов
generate: .buf-format .buf-lint .buf-generate .tidy

# Объявляем, что текущие команды не являются файлами и
# инструментируем Makefile не искать изменения в файловой системе
.PHONY: \
	.bin-deps \
	.buf-format \
	.buf-lint \
	.buf-generate \
	.tidy \
	generate