
<!--- This file is automatically generated by make gen-cli-docs; changes should be made in the go CLI command code (under cmd/kops) -->

## kops completion bash

Generate the autocompletion script for bash

### Synopsis

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(kops completion bash)

To load completions for every new session, execute once:

#### Linux:

	kops completion bash > /etc/bash_completion.d/kops

#### macOS:

	kops completion bash > $(brew --prefix)/etc/bash_completion.d/kops

You will need to start a new shell for this setup to take effect.


```
kops completion bash
```

### Options

```
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
      --config string   yaml config file (default is $HOME/.kops.yaml)
      --name string     Name of cluster. Overrides KOPS_CLUSTER_NAME environment variable
      --state string    Location of state storage (kops 'config' file). Overrides KOPS_STATE_STORE environment variable
  -v, --v Level         number for the log level verbosity
```

### SEE ALSO

* [kops completion](kops_completion.md)	 - Generate the autocompletion script for the specified shell

