#!/bin/bash
if [ ! -d /root/.ethereum/keystore ]; then
    echo "/root/.ethereum/keystore not found, running 'geth init'..."
    geth init /root/ethconfig/genesis.json
#    geth dumpconfig > /root/geth.env
    echo "...done!"
fi

geth "$@"