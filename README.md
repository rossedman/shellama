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

### Using With Different Models

`shellama` can be used with any models that are available through `ollama`, just pass the `--model` flag in (and make sure you already have it downloaded with `ollama`).

```
# run with a different model
$ shellama --model phi3 "provide an example of a bash for loop"

for file in *.txt; do echo "Processing $file"; done
```

### Using With Custom Models

Below is an example of how to run with a custom model using a `Modelfile`. Given the below `Modelfile` you can create a simple coding assistant.

```
FROM llama3.1
PARAMETER temperature 1
SYSTEM You are a coding assistant, you only return code when asked questions.
```

```sh
$ ollama create coding-assistant -f Modelfile
$ shellama -f coding-assistant "provide an example of a bash for loop"

for i in 1 2 3; do echo $i; done
```