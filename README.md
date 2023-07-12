# Go Create Branch

Go Create Branch is a simple command-line tool written in Golang that creates a Git branch with a sanitized name leaving only letters and numbers as well as replacing all spaces with a dash.

## Installation

Make sure you have Golang installed on your machine.

```
go get github.com/username/go-create-branch
```

### Alternative installation

```
git clone git@github.com:nemes1s/go-create-branch.git

cd go-create-branch

go install .
```

## Usage

To use Go Create Branch, run the following command:

```
go-create-branch "branchname"
```

Make sure to replace "branchname" with the name you want to give to your Git branch.

## How it works

Go Create Branch reads the input string from the command line arguments, checks if the input string is empty and sanitizes it by removing all special characters and converting it to lowercase. It then checks if the Git branch already exists and if not, it creates and checks out a new Git branch with the sanitized name.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.