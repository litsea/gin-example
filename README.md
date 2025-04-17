# gin-example

## Build

```shell
make build
```

## Run

### Complete Example

```shell
./app complate
```

Routes:

* `/`: Success
* `/user`: Request with validation and validation error translations
* `/panic`: Test panic error recovery
* `/panic-recovery`: Test new goroutine panic error recovery
* `/no-translate`: Test the log of no translation `msgID`
* `/err-test`: Test `errcode` response and error translations
* `/err-unknown`: Test unknown error(no `errcode` assigned) response and error translations
* `/check-admin`: Test request body fetch in log and Sentry
* `/long-time`: Test timeout middleware
* `/rate-limit`: Test `ratelimit` middleware
* `/update-log-lvl`: Runtime update log level
* `/log`: Test runtime update log level
