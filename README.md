# Paramedic

Paramedic is a tool to diagnose and remediate instances, using Amazon EC2 Systems Manager.

With Paramedic:

- can cancel a running command by sending SIGTERM
- can see output of running commands in near-realtime, using Kinesis Streams
- Operations can be delegated to specific users or groups, using IAM policy
- Operations stored as documents can be used repeatedly
- Operations can be reviewed as a code, using a platform like GitHub Pull Requests

## Setup

TODO

### Configure Kinesis Streams

Create a stream named `paramedic-logs` and configure subscription from CloudWatch Logs:
http://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/SubscriptionFilters.html#DestinationKinesisExample

Note that the filter pattern should be empty.

## Usage

Prepare YAML file which defines a document to be run:

```yaml
# reload-nginx.yaml
name: reload-nginx # Optional
description: Reload nginx via systemctl
scriptFile: reload-nginx
```

Save a script named `reload-yaml`:

```bash
#!/bin/bash
exec systemctl reload nginx
```

Upload a document to SSM:

```
$ paramedic documents upload reload-nginx.yaml
```

Run a command:

```
$ paramedic commands run --document-name=reload-nginx --tags=Env=dev,Role=app
2017/09/27 13:26:19 [INFO] This command will be executed on the following instances
2017/09/27 13:26:19 [INFO]   app-i-aaa (i-aaa)
2017/09/27 13:26:19 [INFO]   app-i-bbb (i-bbb)
Are you sure to continue? (y/N): y
2017/09/27 13:26:20 [INFO] A command '...' started
2017/09/27 13:26:20 [INFO] To see the command status, run 'paramedic commands show --command-id=...'
2017/09/27 13:26:20 [INFO] Output logs will be shown below
[2017-09-27T13:26:21+09:00] [i-aaa] [exit status: 0]
[2017-09-27T13:26:21+09:00] [i-bbb] [exit status: 0]

app-i-aaa (i-aaa) Success
app-i-bbb (i-bbb) Success
```

## Development

### Adding a subcommand

```
$ go get github.com/spf13/cobra/cobra
$ cobra add newcommand
```
