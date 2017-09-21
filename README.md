# Paramedic

## Setup

```
$ paramedic setup
```

### Configure Kinesis Streams

Create a stream named `paramedic-logs` and configure subscription from CloudWatch Logs:
http://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/SubscriptionFilters.html#DestinationKinesisExample

Note that the filter pattern should be empty.

## Usage

Prepare YAML file which defines a command to be run like:

```yaml
# required
id: ReloadNginx
command: ['systemctl', 'reload', 'nginx']

# optional
timeout: 1m
workingDirectory: /home/ryotarai
```

```
$ paramedic run command.yaml
```

## Development

### Adding a subcommand

```
$ go get github.com/spf13/cobra/cobra
$ cobra add newcommand
```
