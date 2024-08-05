# shellama

This is a simple utility that allows you to communicate with Ollama running locally in your shell. This can be used for quick iteration while programming or trouble shooting and is a small binary that does... very little! Here are some examples of what you can do with `shellama`

### Prerequisites

Before running `shellama` you need `ollama` [installed](https://ollama.com/) on your machine and running. The default model is also set to `llama3.1`, this needs to be downloaded if you don't plan to overwrite the model. If you want to use a different model ensure that it is pulled locally using the following command.

```sh
ollama pull llama3.1
```

### Installing

`shellama` can be installed through `homebrew` or through GitHub releases!

```sh
brew tap rossedman/tap
brew install rossedman/tap/shellama
```

### Using

`shellama` is meant to be used as a simple shell utility that can speed up development cycles. I built this to prevent from having to look things up or even paste things into GUI forms for ChatGPT. Because this uses `ollama` it also runs everything locally allowing you lots of leverage over what gets run.

You can ask `shellama` for simple examples like below:

```sh
$ shellama "provide an example of a bash for loop"

for i in {1..5}; do
  echo "Loop iteration: $i"
done
```

You can also `cat` files into `shellama` to be analyzed.

```sh
$ shellama "what does this file do?" $(cat main.go)

This is a Go (Golang) source file that builds and runs the `shellama` command. The main function joins two strings (`version` and `commit`) with "+" in between using `strings.Join()`, then passes them to `cmd.Execute()` to run the shellama command.
```

### Using Profiles

Profiles are a way to use different combinations of models and prompts so you can have different "actors" that can be called on. Below is the example of a profile file placed at `$HOME/.shellama.yaml`

```yaml
profiles:
  - name: default
    model: llama3.1
    prompt: |
      you are a simple programming assistant who only returns code
  - name: advanced
    model: llama3.1
    prompt: |
      you are a programming assistant who can also write advanced code and explain it in depth
  - name: editor
    model: llama3.1
    prompt: |
      you are a copy editor who can help with writing and editing text
```

This would allow you to run commands and switch profiles for different purposes

```sh
$ shellama --profile advanced "provide an example of a bash for loop"
```
