#!/bin/bash

set -e
set -u

sudo docker build  -t scm-build .
sudo docker run -it -v $PWD/out:/out scm-build
