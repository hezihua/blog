###
go build  -ldflags "-X 'hezihua/version.GIT_TAG=v1.0.0' -X 'hezihua/version.GIT_COMMIT=e15f9598a'" -o dist/vblog-api.exe main.go