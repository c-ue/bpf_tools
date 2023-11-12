#!/bin/bash

# install golang
wget -q -O - https://git.io/vQhTU | bash

# install tool chain
apt install build-essential make cmake clang llvm linux-headers-$(uname -r) gcc-multilib libbpf-dev