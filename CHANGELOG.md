## 1.3.0 (unreleased)

- Added support for `unreachable_strategy` configuration
- Added support for `kill_selection` configuration
- Fixed `upgrade_strategy` configuration not being evaluated

## 1.2.0

- Upgrade terraform to 0.9.1
- Support DCOS tokens

## 1.1.0

- Added support for `networkName` to be specified for a task. [Docs](https://mesosphere.github.io/marathon/docs/ip-per-task.html)

## 1.0.4

- Updated terraform vendored deps to 0.7.13

## 1.0.3

- Add support for `fetch` configuration ([marathon docs](http://mesosphere.github.io/marathon/docs/rest-api.html#post-v2-apps), [mesos docs](http://mesos.apache.org/documentation/latest/fetcher/))

## 1.0.2

- Updating go-marathon to `v0.3.0`
- Look at the marathon event stream for deployments
- Fixing diff issue that caused refresh never to update

## 1.0.1

BUILD

- Release `linux` and `osx` binaries.

## 1.0.0

*Note* The released binary was only compiled under OSX. It will not work on linux.

IMPROVEMENTS

- Initial tagged release
