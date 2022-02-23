## [v1.0.88] _2022-02-10_
- **CHANGED**
  - Update client to version v1.0.88
  ([PR #125](https://github.com/cycloidio/cycloid-cli/pull/125))
  - cy-wrapper: Now use Cycloid URL instead of Github to avoid API request limits
  ([PR #126](https://github.com/cycloidio/cycloid-cli/pull/126))
  - cy-wrapper: provide debug mode using CY_DEBUG env var
  ([PR #127](https://github.com/cycloidio/cycloid-cli/pull/127))
  - cy-wrapper: CY_WAIT_NETWORK now default to false
  ([PR #127](https://github.com/cycloidio/cycloid-cli/pull/127))
  - cy-wrapper: CY_VERSION can be enforced to a specific version
  ([PR #127](https://github.com/cycloidio/cycloid-cli/pull/127))
  - cy-wrapper: Reduce the default number of retries and provide new CY_DOWNLOAD_RETRIES variable
  ([PR #127](https://github.com/cycloidio/cycloid-cli/pull/127))
  - cy-wrapper: Error handling when CY_RELEASES_URL is not reachable or incorrect
  ([PR #127](https://github.com/cycloidio/cycloid-cli/pull/127))

## [v1.0.86] _2022-01-04_
- **CHANGED**
  - Update client to version v1.0.86
  ([PR #122](https://github.com/cycloidio/cycloid-cli/pull/122))

## [v1.0.85] _2021-12-08_
- **CHANGED**
  - Update client to version v1.0.85
  ([PR #120](https://github.com/cycloidio/cycloid-cli/pull/120))

## [v1.0.84] _2021-11-10_
- **CHANGED**
  - Update client to version 1.0.84
    ([PR #115](https://github.com/cycloidio/cycloid-cli/pull/115))
- **ADDED**
  - Adding new pipeline synced command
    ([PR #111](https://github.com/cycloidio/cycloid-cli/pull/111))
  - CY wrapper, add more retry on and wait for network
    ([PR #110](https://github.com/cycloidio/cycloid-cli/pull/110))

## [v1.0.82] _2021-10-04_
- **ADDED**
  - Adding cycloid dev readme and local e2e make targets
    ([PR #112](https://github.com/cycloidio/cycloid-cli/pull/105))

- **CHANGED**
  - Update client to version 1.0.82
    ([PR #105](https://github.com/cycloidio/cycloid-cli/pull/105))
  - Changed makefile to allow better use of targets for the cli automatic bump
    ([PR #112](https://github.com/cycloidio/cycloid-cli/pull/105))  
    
- **FIXED**
  - Fix error in help of delete-env command
    ([PR #105](https://github.com/cycloidio/cycloid-cli/pull/105))

## [v1.0.76] _2021-07-08_
- **CHANGED**
  - Update client to version 1.0.76
    ([PR #101](https://github.com/cycloidio/cycloid-cli/pull/101))
  - Update client to version 1.0.72
    ([PR #99](https://github.com/cycloidio/cycloid-cli/pull/99))

- **BREAKING**
  - Switching to API KEY login only
   ([PR #79](https://github.com/cycloidio/cycloid-cli/pull/79))

## [v1.0.64] _2021-03-23_
- **CHANGED**
  - Update to version 1.0.64
   ([PR #92](https://github.com/cycloidio/cycloid-cli/pull/92))

- **BREAKING**
  - Remove create api key command
   ([PR #92](https://github.com/cycloidio/cycloid-cli/pull/92))

- **FIXED**
  - Fix wrong list function execution in `members list-invites`
  ([PR #94](https://github.com/cycloidio/cycloid-cli/pull/94))

## [v1.0.61] _2021-03-15_
- **ADDED**
  - Adding better error details display
   ([PR #76](https://github.com/cycloidio/cycloid-cli/pull/76))
  - Adding organizations `list-children` command
   ([PR #76](https://github.com/cycloidio/cycloid-cli/pull/76))

- **CHANGED**
  - Update to version 1.0.61
   ([PR #84](https://github.com/cycloidio/cycloid-cli/pull/84))
  - wrapper: change behavior to look for the closest lower version
   ([PR #85](https://github.com/cycloidio/cycloid-cli/pull/85))

- **REMOVED**
  - Remove deprecated KPIs `list-avaiable` command
   ([PR #84](https://github.com/cycloidio/cycloid-cli/pull/84))

## [v1.0.58] _2021-02-05_
- **ADDED**
  - Adding new kpis command
   ([PR #81](https://github.com/cycloidio/cycloid-cli/pull/81))
  - Adding new list-invites command
   ([PR #72](https://github.com/cycloidio/cycloid-cli/pull/72))
  - `gen-doc` subcommand
   ([PR #61](https://github.com/cycloidio/cycloid-cli/pull/61))
  - `--insecure` flag to allow TLS verification skipping
   ([Issue #70](https://github.com/cycloidio/cycloid-cli/issues/70))
  - Adding all missing creds type into the CLI (GCP, AWS, ...)
   ([PR #74](https://github.com/cycloidio/cycloid-cli/pull/74))

## [v1.0.53] _2020-12-01_
- **CHANGED**
  - Update to version 1.0.53
   ([PR #62](https://github.com/cycloidio/cycloid-cli/pull/62))
  - Wrapper now fallback to RC version before trying the dev one
   ([PR #62](https://github.com/cycloidio/cycloid-cli/pull/62))
  - Add pipeline list command to list all pipeline in an organization
   ([PR #62](https://github.com/cycloidio/cycloid-cli/pull/62))

## [v1.0.51] _2020-11-12_
- **CHANGED**
  - Update to version 1.0.51
   ([PR #52](https://github.com/cycloidio/cycloid-cli/pull/52))

## [v1.0.50] _2020-11-04_
- **ADDED**
  - Add organization create/delete
   ([PR #51](https://github.com/cycloidio/cycloid-cli/pull/51))
  - `api-keys` commands
   ([PR #57](https://github.com/cycloidio/cycloid-cli/pull/57))

- **CHANGED**
  - `login` method to allow login using API key
   ([PR #57](https://github.com/cycloidio/cycloid-cli/pull/57))

## [v1.0.49] _2020-11-02_
- **ADDED**
  - Add validate-form command
   ([PR #35](https://github.com/cycloidio/cycloid-cli/pull/35))
  - Bump CLI version
   ([PR #35](https://github.com/cycloidio/cycloid-cli/pull/35))
  - printer /helpers for each command
   ([PR #25](https://github.com/cycloidio/cycloid-cli/pull/25))
  - `login list` subcommand
   ([PR #24](https://github.com/cycloidio/cycloid-cli/pull/24))
  - support for child org login
   ([PR #37](https://github.com/cycloidio/cycloid-cli/pull/37))
  - status endpoint implementation
   ([PR #42](https://github.com/cycloidio/cycloid-cli/pull/42))
  - bash/zsh auto-complete
   ([PR #47](https://github.com/cycloidio/cycloid-cli/pull/47))

## [v1.0.47] _2020-09-21_
- **ADDED**
  - Printer with support for `json`, `yaml` and `table` format
  ([PR #4](https://github.com/cycloidio/cycloid-cli/pull/4))
  - First iteration of login command 
  ([PR #9](https://github.com/cycloidio/cycloid-cli/pull/9))

- **CHANGED**
  - Second iteration of login command : allow multiple orgs login
  :warning: the signature of common.ClientCredentials has changed
  ([PR #15](https://github.com/cycloidio/cycloid-cli/pull/15))

- **DEPRECATED**

- **REMOVED**

- **FIXED**

- **SECURITY**

## [1.0.46] _2020-09-20_
- **ADDED**
  - First changelog template
  ([PR #0](https://github.com/cycloidio/cycloid-cli/pull/0))

- **CHANGED**

- **DEPRECATED**

- **REMOVED**

- **FIXED**

- **SECURITY**
