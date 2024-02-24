#!/usr/bin/env node

const meow = require("meow");
const chalk = require("chalk");
const { exec } = require("child_process");
const path = require("path");

function banner() {
  console.log(chalk.yellow("\nGo Build Runner"));
  console.log(
    chalk.yellow(
      "-------------------------------------------------------------"
    )
  );
  console.log(
    chalk.cyan("Linggawasistha Djohari <linggawasistha.djohari@outlook.com>\n")
  );

  console.log(chalk.yellow(`Current working directory:`) + `${process.cwd()}`);
}

function handleExecResult(error, stdout, stderr) {
  if (error) {
    console.error(chalk.red(`Execution Error: ${error}`));
    return;
  }
  if (stderr.trim()) {
    console.error(chalk.yellow(`Execution Stderr: ${stderr}`));
  }
  if (stdout.trim()) {
    console.log(chalk.green("Execution Output:"), chalk.blue(stdout));
  } else {
    console.log(chalk.yellow("Note: The command produced no output."));
  }
}

const cli = meow(
  `
    Usage
      $ node gobuild-node.js (--build | --run)  [--platform <platform>] [--arch <arch>] [--cgo] [--version <version>] [--out <output path>] <main package path>

    Options
      --platform, -p  Target platform (linux, win, mac)
      --arch, -a      Target architecture (e.g., amd64, arm64)
      --cgo           Enable CGO (default: disabled)
      --version, -v   Application version (default: 1.0.0)
      --out, -o       Output binary path (default: ./bin/<main package>)
      --help          Show help text

    Examples
      $ node gobuild-node.js --build --platform linux --arch amd64 --cgo --version 1.0.0 --out bin/trident-server ./server/trident/trident-server
`,
  {
    flags: {
      platform: {
        type: "string",
        alias: "p",
        isRequired: false,
        isMultiple: false,
        default: "linux",
        validate: (platform) =>
          ["linux", "win", "mac"].includes(platform) ||
          "Platform must be one of linux, win, mac",
      },
      arch: {
        type: "string",
        alias: "a",
        isRequired: false,
        default: "amd64",
      },
      cgo: {
        type: "boolean",
        default: false,
      },
      version: {
        type: "string",
        alias: "v",
        default: "1.0.0",
      },
      out: {
        type: "string",
        alias: "o",
        isRequired: false, // We will compute the default if not provided
      },
      build: {
        type: "boolean",
        default: false,
        isRequired: (flags) => !(flags.build && flags.run),
      },
      run: {
        type: "boolean",
        default: false,
        isRequired: (flags) => !(flags.build && flags.run),
      },
    },
    autoHelp: true,
  }
);

try {
  // Compute the default output path if not provided

  banner();

  if (cli.flags.build && cli.flags.run) {
    console.error(
      chalk.red("Error: --build and --run cannot be used together.")
    );
    process.exit(1);
  }

  const defaultOutPath =
    cli.flags.out || `./bin/${path.basename(cli.input[0])}`;

  if (cli.flags.build) {
    // Build go app
    const buildCommand = `CGO_ENABLED=${cli.flags.cgo ? "1" : "0"} GOOS=${
      cli.flags.platform
    } GOARCH=${cli.flags.arch} go build -ldflags="-s -w -X 'main.Version=${
      cli.flags.version
    }' -X 'main.BuildTime=$(date)'" -o ${defaultOutPath} ${cli.input[0]}`;
    console.log(
      chalk.blue("Building with command:"),
      chalk.yellow(buildCommand)
    );

    exec(buildCommand, (error, stdout, stderr) => {
      handleExecResult(error, stdout, stderr);
      console.log(chalk.green("Build successful!"), chalk.blue(stdout));
    });
  } else if (cli.flags.run) {
    // Run go
    const runCommand = `go run ${cli.input[0]}`;
    console.log(chalk.blue("Running with command:"), chalk.yellow(runCommand));
    exec(runCommand, (error, stdout, stderr) =>
      handleExecResult(error, stdout, stderr)
    );
  }
} catch (error) {
  console.error(chalk.red(error.message));
}
