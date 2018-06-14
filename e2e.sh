#!/bin/bash

RED='\033[0;31m'
GREEN='\033[0;32m'
RESET='\033[0m'

FAILED_TESTS=""


passed() {
    printf "${GREEN}PASSED\n${RESET}"
}

failed() {
    printf "${RED}FAILED\n${RESET}"
}

# Expected exit code, command
test_command() {
    EXPECTED_CODE=$1
    CMD="${@:2}"
    echo
    echo "TEST: $CMD"
    echo "Output:"
    echo "---------"
    eval $CMD
    ACTUAL_CODE=$?
    echo "---------"
    echo -n "Status: "
    echo $?
    if [ $ACTUAL_CODE -ne $EXPECTED_CODE ]; then
        echo "$ACTUAL_CODE != $EXPECTED_CODE"
        failed
	FAILED_TESTS="$FAILED_TESTS\n - $CMD"
    else
	passed
    fi
    echo "---------"
}

summary() {
    if [ "$FAILED_TESTS" != "" ]; then
        echo "Failed Tests:"
	printf "$FAILED_TESTS\n"
	exit 1
    else
        echo "All tests passed!"
	exit 0
    fi
}

#
# guide subcommand
#
# Failures
test_command 1 ./image-helpgen guide
test_command 1 ./image-helpgen guide -basename /tmp/e2e

# 
# dockerfile subcommand
#
# Successes
test_command 0 ./image-helpgen dockerfile -template ./template.tpl -dockerfile example/Dockerfile -basename /tmp/e2e 
# Failures
test_command 1 ./image-helpgen dockerfile -template ./template.tpl -dockerfile DOESNOTEXIST -basename /tmp/e2e 
test_command 1 ./image-helpgen dockerfile -dockerfile example/Dockerfile -basename /tmp/e2e 

# 
# man subcommand
#
# Successes
test_command 0 ./image-helpgen man -basename /tmp/e2e
# Failures
test_command 1 ./image-helpgen man -basename idonotexist

# 
# version subcommand
#
# Successes
test_command 0 ./image-helpgen version
test_command 0 ./image-helpgen version -verbose


# Summary
summary
