#!/bin/bash
if golangci-lint run; then
  echo "
  _________                                   _____     .__  .__
 /   _____/__ __   ____  ____  ____   _______/ ____\_ __|  | |  | ___.__.
 \_____  \|  |  \_/ ___\/ ___\/ __ \ /  ___/\   __\  |  \  | |  |<   |  |
 /        \  |  /\  \__\  \__\  ___/ \___ \  |  | |  |  /  |_|  |_\___  |
/_______  /____/  \___  >___  >___  >____  > |__| |____/|____/____/ ____|
        \/            \/    \/    \/     \/                       \/
"
  echo "No errors found"
else
  echo "Errors found. Do you have typing errors in your code."
fi
