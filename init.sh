#!/bin/bash
chmod 666 /mofahezi
chmod +x /mofahezi/check.out
cd /mofahezi
nohup ./check.out >/dev/null 2>&1 &