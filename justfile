echo:
    #!/usr/bin/env bash
    pushd echo
    go build .
    maelstrom test -w echo --bin ./echo --node-count 1 --time-limit 10
    popd

unique-ids:
    #!/usr/bin/env bash
    pushd unique-ids
    go build .
    maelstrom test -w unique-ids --bin ./unique-ids --time-limit 30 --rate 1000 --node-count 3 --availability total --nemesis partition
    popd
