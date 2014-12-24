# doing a dev build

make dev

# running an acceptance test

make testacc TEST=./builtin/providers/marathon TESTARGS='-run=MarathonApp_basic'
