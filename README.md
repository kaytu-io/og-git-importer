# MyGitApp

**MyGitApp** is a versatile command-line tool designed to simplify fetching and diffing Git repositories across multiple platforms, including GitHub, Azure DevOps, GitLab, and other private Git servers. Whether you're working with public or private repositories, MyGitApp provides seamless integration using various authentication methods such as Personal Access Tokens (PATs), SSH Keys/GPG, and Username/Password combinations.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Configuration](#configuration)
  - [Authentication Methods](#authentication-methods)
    - [1. Personal Access Token (PAT)](#1-personal-access-token-pat)
    - [2. SSH Keys/GPG](#2-ssh-keysgpg)
    - [3. Username/Password](#3-usernamepassword)
- [Usage](#usage)
  - [Fetch Command](#fetch-command)
    - [Examples](#examples)
      - [1. Fetching from GitHub](#1-fetching-from-github)
      - [2. Fetching from Azure DevOps](#2-fetching-from-azure-devops)
      - [3. Fetching from GitLab](#3-fetching-from-gitlab)
      - [4. Fetching from a Private Git Server](#4-fetching-from-a-private-git-server)
  - [Diff Command](#diff-command)
    - [Examples](#examples-1)
      - [1. Diffing Between Two Commits in a Local Repository](#1-diffing-between-two-commits-in-a-local-repository)
      - [2. Diffing Between Two Commits in a Remote GitHub Repository](#2-diffing-between-two-commits-in-a-remote-github-repository)
      - [3. Diffing Between Two Commits in a Remote Azure DevOps Repository](#3-diffing-between-two-commits-in-a-remote-azure-devops-repository)
      - [4. Diffing Between Two Commits in a Remote GitLab Repository](#4-diffing-between-two-commits-in-a-remote-gitlab-repository)
- [Troubleshooting](#troubleshooting)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Multi-Platform Support**: Seamlessly interact with GitHub, Azure DevOps, GitLab, and other private Git servers.
- **Flexible Authentication**: Supports Personal Access Tokens (PATs), SSH Keys/GPG, and Username/Password combinations.
- **Diff Functionality**: Compare two commits within a repository and output the differences in a structured JSON format.
- **Clean Output**: Logs are directed to `stderr`, ensuring that JSON outputs remain clean and valid for further processing.

## Installation

Ensure you have [Go](https://golang.org/dl/) installed (version 1.16 or later).

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/yourusername/mygitapp.git
   cd mygitapp

   cd mygitapp
MyGitApp
MyGitApp is a versatile command-line tool designed to simplify fetching and diffing Git repositories across multiple platforms, including GitHub, Azure DevOps, GitLab, and other private Git servers. Whether you're working with public or private repositories, MyGitApp provides seamless integration using various authentication methods such as Personal Access Tokens (PATs), SSH Keys/GPG, and Username/Password combinations.

Table of Contents
Features
Installation
Configuration
Authentication Methods
1. Personal Access Token (PAT)
2. SSH Keys/GPG
3. Username/Password
Usage
Fetch Command
Examples
1. Fetching from GitHub
2. Fetching from Azure DevOps
3. Fetching from GitLab
4. Fetching from a Private Git Server
Diff Command
Examples
1. Diffing Between Two Commits in a Local Repository
2. Diffing Between Two Commits in a Remote GitHub Repository
3. Diffing Between Two Commits in a Remote Azure DevOps Repository
4. Diffing Between Two Commits in a Remote GitLab Repository
Troubleshooting
Contributing
License
Features
Multi-Platform Support: Seamlessly interact with GitHub, Azure DevOps, GitLab, and other private Git servers.
Flexible Authentication: Supports Personal Access Tokens (PATs), SSH Keys/GPG, and Username/Password combinations.
Diff Functionality: Compare two commits within a repository and output the differences in a structured JSON format.
Clean Output: Logs are directed to stderr, ensuring that JSON outputs remain clean and valid for further processing.
Installation
Ensure you have Go installed (version 1.16 or later).

Clone the Repository:

bash
Copy code
git clone https://github.com/yourusername/mygitapp.git
cd mygitapp
Build the Application:

bash
Copy code
go build -o mygitapp
This will generate an executable named mygitapp in the current directory.

Configuration
MyGitApp utilizes environment variables to handle authentication securely. Depending on the authentication method you choose, set the appropriate environment variables as described below.

Authentication Methods
1. Personal Access Token (PAT)
Supported Platforms: GitHub, Azure DevOps, GitLab

Usage:

Set Environment Variables:

bash
Copy code
export GIT_USERNAME=your_username_or_token_username
export GIT_PASSWORD=your_personal_access_token
Notes:

For GitHub, your PAT replaces the password. The username can be your GitHub username.
For Azure DevOps and GitLab, the PAT is used similarly, replacing the password. Ensure your PAT has the necessary scopes/permissions (e.g., read_repository).
2. SSH Keys/GPG
Supported Platforms: GitHub, Azure DevOps, GitLab, and other SSH-enabled Git servers.

Setup:

Generate SSH Keys (if not already done):

bash
Copy code
ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
Follow the prompts to save the key (default location is ~/.ssh/id_rsa) and set a passphrase.

Add SSH Key to Git Platform:

GitHub: Adding a new SSH key to your GitHub account
Azure DevOps: Add an SSH key to your Azure DevOps profile
GitLab: Adding an SSH key to your GitLab account
Ensure known_hosts is Updated:

The application automatically handles host key verification. Ensure your ~/.ssh/known_hosts includes the SSH host keys for your Git platforms.

3. Username/Password
Supported Platforms: Primarily for private Git servers that do not support PATs or SSH.

Usage:

Set Environment Variables:

bash
Copy code
export GIT_USERNAME=your_username
export GIT_PASSWORD=your_password
Notes:

This method is less secure and generally not recommended for platforms that support PATs or SSH.
Ensure that the Git server you're interacting with allows Username/Password authentication.
Usage
MyGitApp provides two primary commands: fetch and diff.

Fetch Command
The fetch command clones a Git repository from a specified URI to a local directory. It supports cloning from GitHub, Azure DevOps, GitLab, and other private Git servers using various authentication methods.

Syntax:

bash
Copy code
./mygitapp fetch <git URI> [targetDir]
<git URI>: The URL of the Git repository to clone.
[targetDir]: (Optional) The directory where the repository will be cloned. If not specified, it defaults to the repository name or a timestamped directory.
Examples
1. Fetching from GitHub
Using HTTPS with PAT:

bash
Copy code
export GIT_USERNAME=your_github_username
export GIT_PASSWORD=your_github_pat

./mygitapp fetch "https://github.com/kaytu-io/managed-platform-config.git" "managed-platform-config"
Using SSH:

bash
Copy code
./mygitapp fetch "git@github.com:kaytu-io/managed-platform-config.git" "managed-platform-config"
Fetching a Specific Branch into a Custom Directory:

bash
Copy code
./mygitapp fetch "https://github.com/kaytu-io/managed-platform-config/tree/development" "custom-folder"
2. Fetching from Azure DevOps
Using HTTPS with PAT:

bash
Copy code
export GIT_USERNAME=your_azure_devops_username
export GIT_PASSWORD=your_azure_devops_pat

./mygitapp fetch "https://dev.azure.com/acme-org/acme-project/_git/acme-repo?version=GBmain" "acme-repo"
Using SSH:

bash
Copy code
./mygitapp fetch "git@ssh.dev.azure.com:v3/acme-org/acme-project/acme-repo" "acme-repo"
3. Fetching from GitLab
Using HTTPS with PAT:

bash
Copy code
export GIT_USERNAME=your_gitlab_username
export GIT_PASSWORD=your_gitlab_pat

./mygitapp fetch "https://gitlab.com/acme-group/acme-project.git?ref=main" "acme-project"
Using SSH:

bash
Copy code
./mygitapp fetch "git@gitlab.com:acme-group/acme-project.git" "acme-project"
4. Fetching from a Private Git Server
Using HTTPS with Username/Password:

bash
Copy code
export GIT_USERNAME=your_username
export GIT_PASSWORD=your_password

./mygitapp fetch "https://private.gitserver.com/acme-project.git" "acme-project"
Using SSH:

bash
Copy code
./mygitapp fetch "git@private.gitserver.com:acme-project.git" "acme-project"
Diff Command
The diff command compares two commits within a Git repository and outputs the differences in a structured JSON format. It supports both local and remote repositories.

Syntax:

bash
Copy code
./mygitapp diff <repository path or remote URI> <first_commit_SHA> [second_commit_SHA]
<repository path or remote URI>: The path to a local repository or the remote URI of the repository you want to analyze.
<first_commit_SHA>: The SHA-1 hash of the first commit.
[second_commit_SHA]: (Optional) The SHA-1 hash of the second commit. If omitted, the latest commit will be used.
Examples
1. Diffing Between Two Commits in a Local Repository
bash
Copy code
./mygitapp diff "./managed-platform-config" 0afcb0c7d8407948f3381a3bd24d0e44adbf28ff 1562507995016574770549bdc173e8258c9ea6fa > diff_output.json
Explanation:
Compares commit 0afcb0c7d8407948f3381a3bd24d0e44adbf28ff with 1562507995016574770549bdc173e8258c9ea6fa in the local repository located at ./managed-platform-config.
Outputs the differences to diff_output.json.
Logs are directed to stderr.
2. Diffing Between Two Commits in a Remote GitHub Repository
bash
Copy code
./mygitapp diff "https://github.com/kaytu-io/managed-platform-config.git" 0afcb0c7d8407948f3381a3bd24d0e44adbf28ff 1562507995016574770549bdc173e8258c9ea6fa > github_diff.json
Explanation:
Clones the GitHub repository.
Compares the specified commits.
Outputs the differences to github_diff.json.
Logs are directed to stderr.
3. Diffing Between Two Commits in a Remote Azure DevOps Repository
bash
Copy code
./mygitapp diff "https://dev.azure.com/acme-org/acme-project/_git/acme-repo" 0afcb0c7d8407948f3381a3bd24d0e44adbf28ff 1562507995016574770549bdc173e8258c9ea6fa > azure_diff.json
Explanation:
Clones the Azure DevOps repository.
Compares the specified commits.
Outputs the differences to azure_diff.json.
Logs are directed to stderr.
4. Diffing Between Two Commits in a Remote GitLab Repository
bash
Copy code
./mygitapp diff "https://gitlab.com/acme-group/acme-project.git" 0afcb0c7d8407948f3381a3bd24d0e44adbf28ff 1562507995016574770549bdc173e8258c9ea6fa > gitlab_diff.json
Explanation:
Clones the GitLab repository.
Compares the specified commits.
Outputs the differences to gitlab_diff.json.
Logs are directed to stderr.
Authentication Considerations for Diff Command
When diffing remote repositories that require authentication:

Ensure Environment Variables:

Set GIT_USERNAME and GIT_PASSWORD for HTTPS cloning of private repositories.
Ensure SSH keys are correctly configured for SSH cloning.
Example:

bash
Copy code
export GIT_USERNAME=your_username
export GIT_PASSWORD=your_pat_or_password

./mygitapp diff "https://gitlab.com/acme-group/acme-project.git" abc123 def456 > gitlab_diff.json
Troubleshooting
Authentication Errors:

Verify that your environment variables (GIT_USERNAME and GIT_PASSWORD) are correctly set.
Ensure that your SSH keys are correctly configured and added to your Git platform.
Check that your Personal Access Tokens have the necessary scopes/permissions.
Invalid Commit SHAs:

Ensure that the commit SHAs you provided exist in the repository.
Verify that both commits are part of the same branch or repository history.
Network Issues:

Ensure that your network allows outbound connections to the Git platforms.
Check for firewall or proxy settings that might block Git operations.
Permission Issues:

Ensure you have the necessary permissions to clone and access the repositories.
Verify that the target directory is writable.
Dependency Issues:

Run go mod tidy to ensure all dependencies are correctly installed.
Ensure you have the latest versions of Go and dependencies.
Logs for Debugging:

Increase the log level to debug by setting LOG_LEVEL=debug to get more detailed logs.
bash
Copy code
export LOG_LEVEL=debug
./mygitapp fetch "https://gitlab.com/acme-group/acme-project.git?ref=main" "acme-project"
Contributing
Contributions are welcome! Please open issues or submit pull requests for any features, bugs, or enhancements.

Fork the Repository:

bash
Copy code
git clone https://github.com/yourusername/mygitapp.git
cd mygitapp
Create a New Branch:

bash
Copy code
git checkout -b feature/your-feature-name
Commit Your Changes:

bash
Copy code
git commit -m "Add your commit message"
Push to Your Fork:

bash
Copy code
git push origin feature/your-feature-name
Create a Pull Request.

License
This project is licensed under the MIT License.

Note: Replace yourusername, your_github_username, your_azure_devops_username, your_gitlab_username, your_gitlab_pat, and other placeholders with your actual information.