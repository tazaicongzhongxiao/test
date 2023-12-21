#!/bin/bash
set -ex
make
./serverHello --config=config/config.conf
