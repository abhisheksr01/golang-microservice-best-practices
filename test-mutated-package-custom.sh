#!/bin/bash

# This exec script implements
# - the replacement of the original file with the mutation,
# - the execution of all tests originating from the package of the mutated file,
# - and the reporting if the mutation was killed.

if [ -z ${MUTATE_CHANGED+x} ]; then echo "MUTATE_CHANGED is not set"; exit 1; fi
if [ -z ${MUTATE_ORIGINAL+x} ]; then echo "MUTATE_ORIGINAL is not set"; exit 1; fi
if [ -z ${MUTATE_PACKAGE+x} ]; then echo "MUTATE_PACKAGE is not set"; exit 1; fi

function clean_up {
	if [ -f $MUTATE_ORIGINAL.tmp ];
	then
		mv $MUTATE_ORIGINAL.tmp $MUTATE_ORIGINAL
	fi
}

function sig_handler {
	clean_up

	exit $GOMUTESTING_RESULT
}
trap sig_handler SIGHUP SIGINT SIGTERM

export GOMUTESTING_DIFF=$(diff -u $MUTATE_ORIGINAL $MUTATE_CHANGED)

mv $MUTATE_ORIGINAL $MUTATE_ORIGINAL.tmp
cp $MUTATE_CHANGED $MUTATE_ORIGINAL

export MUTATE_TIMEOUT=${MUTATE_TIMEOUT:-10}

if [ -n "$TEST_RECURSIVE" ]; then
	TEST_RECURSIVE="/..."
fi

touch report.html
GOMUTESTING_TEST=$(go test -timeout $(printf '%ds' $MUTATE_TIMEOUT) $MUTATE_PACKAGE$TEST_RECURSIVE 2>&1)
export GOMUTESTING_RESULT=$?

if [ "$MUTATE_DEBUG" = true ] ; then
    echo "<pre><code>" >> report.html
	echo "$GOMUTESTING_TEST" >> report.html
    echo "</pre></code>" >> report.html
fi

clean_up

case $GOMUTESTING_RESULT in
0) # tests passed -> FAIL
    printf "\033[0;31m"
    echo "<pre><code><p style=\"color:red;\">" >> go-mutesting-error-report.html
	echo "$GOMUTESTING_DIFF" >> go-mutesting-error-report.html
    echo "</p></pre></code></br><hr></br>" >> go-mutesting-error-report.html
    printf "\033[0m"

	exit 1
	;;
1) # tests failed -> PASS
	if [ "$MUTATE_DEBUG" = true ] ; then
        echo "<pre><code><p style=\"color:black;\">" >> go-mutesting-ok-report.html
		echo "$GOMUTESTING_DIFF"  >> go-mutesting-ok-report.html
        echo "</p></pre></code></br><hr></br>" >> go-mutesting-ok-report.html
	fi

	exit 0
	;;
2) # did not compile -> SKIP
	if [ "$MUTATE_VERBOSE" = true ] ; then
		echo "Mutation did not compile"
	fi

	if [ "$MUTATE_DEBUG" = true ] ; then
		echo "$GOMUTESTING_DIFF"
	fi

	exit 2
	;;
*) # Unkown exit code -> SKIP
	echo "Unknown exit code"
	echo "$GOMUTESTING_DIFF"

	exit $GOMUTESTING_RESULT
	;;
esac