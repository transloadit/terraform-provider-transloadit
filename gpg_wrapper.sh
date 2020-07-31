#!/bin/sh
export GPG_TTY=$(tty)
echo $SIGN_PASS|gpg --batch --no-tty --pinentry-mode loopback --passphrase-fd 0 $@