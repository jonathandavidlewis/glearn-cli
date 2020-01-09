# glearn-cli

This is the command line interface for previewing and publishing Learn curriculum.

## Installation with Homebrew

Install
```
brew tap gSchool/learn
brew install learn
learn set --api_token={your token from https://learn-2.galvanize.com/api_token}
```

Update
```
brew upgrade learn
```

Uninstall
```
brew uninstall learn
```

## Getting Started

```
mkdir test-content && cd test-content
learn new
```
Follow directions in `test-content/README.md`

## Alternatives to Homebrew

Use curl on Mac
```
curl -L $(curl -s https://api.github.com/repos/gSchool/glearn-cli/releases/latest | grep -o "http.*Darwin_x86_64.tar.gz") | tar -xzf - -C /usr/local/bin
```

Use curl on Linux
```
curl -L $(curl -s https://api.github.com/repos/gSchool/glearn-cli/releases/latest | grep -o "http.*Linux_x86_64.tar.gz") | tar -xzf - -C /usr/local/bin
```

Download binaries for all platforms directly from
https://github.com/gSchool/glearn-cli/releases

After using any of these options, set your API token with
`learn set --api_token={your token from https://learn-2.galvanize.com/api_token}` or setting `api_token` directly in `~/.glearn-config.yaml`.

## Example Usage

See a list of commands
```
learn help
```

Preview a single file
```
learn preview my_file.md
```

Preview an entire directory:
```
learn preview my_curriculum_directory
```

Publishing an entire repo
* add/commit/push to github
```
learn publish
```

## Development
Build
```
go build -o glearn-cli main.go
```

Run
```
./learn [commands...] [flags...]
```

Or for quicker iterations:
```
go run main.go [commands...] [flags...]
```

### Specifying Learn App URL

By default, the CLI tool will use Learn's base url `https://learn-2.galvanize.com`. This value can be changed by exporting the environment variable `LEARN_BASE_URL` to specify the desired address. This is convenient for testing stage/PR environments.

## Releases

Create a github token with `repo` access. This gives you the ability to push releases and their binaries and allows glearn-cli write commits when necessary.

Create a new semantic version tag (ex. 0.1.0)
```
git tag -a v0.1.0 -m "Some new release commit"
```

Push new tag
```
git push origin v0.1.0
```

To release run:
```
GITHUB_TOKEN=your_githhub_token goreleaser release --rm-dist
```
