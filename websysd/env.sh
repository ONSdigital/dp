# This file defines env vars required by websysd scripts, with a default value if they are not set.

# Source bash profile to get the en vars defined there
source ~/.bash_profile

# Path where Go Modules will be cloned (default: ~/gomod).
export GO_MOD_PATH=${GO_MOD_PATH:-"$HOME/gomod"}