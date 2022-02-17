# ENV printer

Support: 
- [x] json
- [X] yaml

Install:
```sh
go get -u github.com/phamvinhdat/mo
```

    Usage:
    mo [command]

    Available Commands:
    completion  Generate the autocompletion script for the specified shell
    help        Help about any command
    k8s         k8s env format

    Flags:
    -e, --extension string   config file extension, require if file path don't contain extension
    -f, --file string        config file. ex: ./config.yaml
    -h, --help               help for mo
    -p, --prefix string      env's prefix

    Use "mo [command] --help" for more information about a command.