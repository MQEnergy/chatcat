APP := chatcat
CREATE_DMG ?= $(shell which create-dmg)
SYS := $(shell uname -s)

define check
	command -v $(1) 1>/dev/null || $(2)
endef
#
#define checkEnv
#ifeq ($(strip $(ENVIRONMENT)),)
#$(warning "ENV is required")
#endif
#endef
#
#$(call checkEnv)

help:
	@echo "$(APP)"
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## dev: 启动调试模式
dev:
	wails dev

## build: 编译app的配置文件
build: clean
	wails build -ldflags="-X main.env=$(ENV)"

## ly-cli: 邻医工具包
ly-cli:
	@go install -ldflags="-s" ./ly-cli

## darwin/amd64: 编译darwin amd64 程序
darwin/amd64:

ifeq ($(SYS),Windows_NT)
	@GOARCH=amd64 GOOS=darwin CGO_ENABLED=1 wails build -platform=darwin/amd64 -ldflags="-X main.env=$(ENV)"
else
	@GOARCH=amd64 GOOS=darwin CGO_ENABLED=1 && wails build -platform=darwin/amd64 -ldflags="-X main.env=$(ENV)"
endif

## darwin/arm64: 编译darwin arm64 程序
darwin/arm64:
ifeq ($(SYS),Windows_NT)
	@GOARCH=arm64 GOOS=darwin CGO_ENABLED=1 wails build -platform=darwin/arm64 -ldflags="-X main.env=$(ENV)"
else
	@GOARCH=arm64 GOOS=darwin CGO_ENABLED=1 && wails build -platform=darwin/arm64 -ldflags="-X main.env=$(ENV)"
endif

## windows/amd64: 编译window amd64 程序
windows/amd64:
ifeq ($(SYS),Windows_NT)
	@GOARCH=amd64 GOOS=windows CGO_ENABLED=1 wails build -nsis -platform=windows/amd64 -ldflags="-X main.env=$(ENV)"
else
	@GOARCH=amd64 GOOS=windows CGO_ENABLED=1 && wails build -nsis -platform=windows/amd64 -ldflags="-X main.env=$(ENV)"
endif

## windows/arm64: 编译window arm64 程序
windows/arm64:
ifeq ($(SYS),Windows_NT)
	@GOARCH=arm64 GOOS=windows CGO_ENABLED=1 wails build -nsis -platform=windows/arm64 -ldflags="-X main.env=$(ENV)"
else
	@GOARCH=arm64 GOOS=windows CGO_ENABLED=1 && wails build -nsis -platform=windows/arm64 -ldflags="-X main.env=$(ENV)"
endif

## windows/386: 编译window 386 程序
windows/386:
ifeq ($(SYS), Windows_NT)
	@GOARCH=386 GOOS=windows CGO_ENABLED=1 wails build -nsis -platform=windows/386 -ldflags="-X main.env=$(ENV)"
else
	@GOARCH=386 GOOS=windows CGO_ENABLED=1 && wails build -nsis -windowsconsole -platform=windows/386 -ldflags="-X main.env=$(ENV)"
endif

## clean: 清除编译的程序
clean:
	rm -rf ./build/bin/*

## dmg: 生成dmg文件
dmg:
	@@$(call check,create-dmg,brew install create-dmg)
	$(CREATE_DMG) \
		--volname "$(APP)" \
	  	--volicon "assets/application_icon.icns" \
		--background "assets/background.png" \
	  	--window-pos 200 140 \
	  	--window-size 800 484 \
		--icon-size 80 \
		--hide-extension "$(APP).app" \
		--icon "$(APP).app" 220 270 \
		--app-drop-link 580 270 \
		"build/bin/$(APP).dmg" \
		"build/bin/$(APP).app"

.PHONY: help clean dev build(test/prod) ly-cli darwin(amd64/arm64) window(amd64/arm64/386) dmg