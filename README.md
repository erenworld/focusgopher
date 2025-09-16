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