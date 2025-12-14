# Contributing to nuke-port

Thank you for your interest in contributing to nuke-port! We welcome contributions from everyone.

## How to Contribute

### Reporting Bugs

- Open an issue on GitHub.
- Describe the bug clearly.
- Include OS version and `lsof`/`netstat` output if relevant.

### Suggesting Enhancements

- Open a feature request issue.
- Explain why the feature would be useful.

### Pull Requests

1.  **Fork** the repository to your GitHub account.
2.  **Clone** your fork locally:
    ```bash
    git clone https://github.com/geekaara/nuke-port.git
    ```
3.  **Create a branch** for your feature or fix:
    ```bash
    git checkout -b feature/my-new-feature
    ```
4.  **Make your changes**. Ensure code follows the existing style.
5.  **Test your changes**.
    ```bash
    go build -o nuke
    ./nuke 8080 # Test locally
    ```
6.  **Commit** your changes:
    ```bash
    git commit -m "Add some feature"
    ```
7.  **Push** to your fork:
    ```bash
    git push origin feature/my-new-feature
    ```
8.  **Open a Pull Request** on the main repository.

## Development

- **Language**: Go 1.x
- **Dependencies**: None (Standard Library only)
- **Cross-Platform**: Please test or consider Mac, Linux, and Windows build tags.

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
