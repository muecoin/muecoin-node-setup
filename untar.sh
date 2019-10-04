#!/bin/sh
tar -xzvf mon-2.1.4-x86_64-linux-gnu.tar.gz
cd mon/bin/
cp monetaryunitd /usr/bin/monetaryunitd
cp monetaryunit-cli /usr/bin/monetaryunit-cli
