mackerel-plugin-delayed-job-count
====

delayed\_job custom metrics plugin for mackerel.io agent.

## Synopsis

```
mackerel-plugin-delayed-job-count -dbconf=</path/to/database.yml> [-env=<environment>]
```

## Requirements

- [delayed_job](https://rubygems.org/gems/delayed_job)

## Example of mackerel-agent.conf

```
[plugin.metrics.delayed_job_count]
command = "/path/to/mackerel-plugin-delayed-job-count -dbconf=/path/to/database.yml -env=production"
```
