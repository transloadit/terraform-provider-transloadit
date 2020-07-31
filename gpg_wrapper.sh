#!/bin/sh
export GPG_TTY=$(tty)
echo $SIGN_PASS|gpg $@