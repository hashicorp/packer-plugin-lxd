#!/bin/bash

getent group sudo >/dev/null 2>&1 || groupadd --system sudo
useradd --create-home -s /bin/bash -G sudo -U bjk