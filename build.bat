@echo off
set cur_dir=%~dp0%
set file_dir=%cur_dir%src\webdemo\
pushd %file_dir%
set gofile=main.go route.go loginController.go ajaxController.go adminController.go
echo go build %gofile%
go build -o ../../main.exe %gofile%
popd
echo go build finished!
pause