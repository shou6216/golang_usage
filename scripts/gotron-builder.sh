#!/bin/sh -eu

cd scripts;
ls;
./gotron-builder-amd64-win.exe --go=../cmd/deposit_chart --app=../web;
