## medic profile list

List profiles within a mediator control plane

### Synopsis

The medic profile list subcommand lets you list profiles within a
mediator control plane for an specific project.

```
medic profile list [flags]
```

### Options

```
  -h, --help              help for list
  -o, --output string     Output format (json, yaml or table) (default "table")
  -p, --provider string   Provider to list profiles for
```

### Options inherited from parent commands

```
      --config string      config file (default is $PWD/config.yaml)
      --grpc-host string   Server host (default "staging.stacklok.dev")
      --grpc-insecure      Allow establishing insecure connections
      --grpc-port int      Server port (default 443)
```

### SEE ALSO

* [medic profile](medic_profile.md)	 - Manage profiles within a mediator control plane
