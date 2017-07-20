#!/bin/bash

set -e
set -u

gofmt -w .
sudo docker build  -t scm-build .
sudo docker run -it -v $PWD/out:/out scm-build
