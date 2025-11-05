# focusgopher

CLI tool to eliminate online distractions and boost focus. Customize your blacklist, focus on tasks, and reclaim your time.

## Installation

```bash
go install github.com/erenworld/focusgopher
```

## Usage

focusgopher needs sudo to modify `/etc/hosts` file. It won't affect your existing configuration, the changes made my focusgopher are separated by `# focusgopher` comment.

```bash
sudo focusgopher
```

## Supported platforms

- macOS
- Linux
- Windows

## Run tests

```bash
go test -v -race ./...
```

# Some other system entries

Your `/etc/hosts/` will look like this.
```
127.0.0.1 localhost

#focusgopher:start
#focusgopher:on
127.0.0.1 facebook.com
127.0.0.1 twitter.com

# 127.0.0.1 reddit.com

#focusgopher:end
```