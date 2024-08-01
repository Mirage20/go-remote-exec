# Go Remote Execution

A REST API that allows you to execute commands on a remote machine.


## Example

```bash
curl -X POST http://localhost:12084/execute -d '{"command": "ls -la"}'
```
