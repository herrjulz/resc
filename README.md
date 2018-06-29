# ReSc (Remote Script)

`resc` is a CLI tool that executes remote _bash_ scripts located on Github on your local machine. This allows you to:

- **Share your scripts easily**
- **Run your scripts from everywhere** (as long as it is a mac or linux machine 😉)

## Install

### OS X

```
$ wget -O /usr/local/bin/resc https://github.com/JulzDiverse/resc/releases/download/v0.1.0/resc-darwin-amd64 && chmod +x /usr/local/bin/resc
```

OR

**Homebrew**

```
$ brew tap julzdiverse/tools
$ brew install resc
```

### Linux

```
$ wget -O /usr/bin/resc https://github.com/JulzDiverse/resc/releases/download/v0.1.0/resc-linux-amd64 && chmod +x /usr/bin/resc
```

## Hello, World! 

You can use `resc` to run `resc` scripts located in any Github repository. For example, let's run a `Hello, World!` script:

```bash
$ resc run hello-world --repo JulzDiverse/remote-scripts
```

This runs the `hello-world` script located in the `JulzDiverse/remote-scripts` repository.

## `resc` scripts

`resc` scripts requires one or more top level directories inside a GitHub repository that contain a `run.sh` script and a `README.md` file. In case of the `hello-world` script the directory looks like this:

```
.remote-scripts
└── hello-world 
   ├── run.sh
   └── README.md
``` 

- The `directory name` (here `hello-world`) indicates the script
- The `run.sh` is the bash script that is run by `resc`
- The `README.md` is a Markdown file that provides information for a script (eg description, usage). The `README.md` is processed by the `resc` CLI and should only contain the following markdown syntax:
  - H1 (#)
  - H2 (##)
  - Bold (\*\*text\*\*)
  - Italic (\_text\_)

## Usage

### 🏃  Run `resc` scripts (`run`) 
Running a `resc` script is nothing more than:

```
$ resc run <script-name> --repo <github-user|github-org>/<github-repo>
```

or if you have set a repository, it's even simpler:

```
$ resc run <script-name>
```

You can also provide parameters to a script using `--args|-a` option. Try it:

```
$ resc run hello-world -r JulzDiverse/remote-scripts -a your-name
```

### 🧐 Set base `resc` script repository (`set`) 

You can set a base `resc` script repository, s.t you are not required to specify the repository of a script everytime you execute the `resc run`.

```
$ resc set <github-user|github-org>/<github-repo>
```

### 📖  Get some script info (`man`) 

If you want to know what a script does before you run it, you can check the provided README by calling `man`:

```
$ resc man <script-name> 
```

### 🖨 Print a script (`print`) 

If you want to see what a script exactly does or you want to save it to your local machine, you can use the `print` command:

```
$ resc print <script-name>
```

to save a script, pipe it to a file:

```
$ resc print <script-name> > script.sh
```

## 💻  Development 

```
$ go get github.com/JulzDiverse/resc (or git clone repository)
$ dep ensure
```
