
# Tv shows / Anime CLI

This is a CLI tool to manage your TV shows and Anime. It allows you to list all the tv shows and anime, and also watch specific tv shows and anime.

## Usage

The CLI tool accepts the following flags:

--tvshows or -t : List all the tv shows
--anime or -a : List all the anime
You can also use the command tv or t followed by the name of the tv show or anime you want to watch.

For example, to watch the tv show "Suits", you can run:

```console
./main tv Suits
```

or, you can input it as an alias

```console
./main t s
```

and it's going to generate an output

```console
I'm sure you gonna enjoy watching:  Suits
```

## Installation

To use this CLI tool, you need to have Go installed on your machine.

- Clone this repository

```console
git clone https://github.com/LinJTF/tv_shows_cli.git
```

- Build the binary:

```console
go build main.go
```

- Run the binary:

```console
./main help
```

or

```console
./main h
```

## Author

Felipe Lin

felipejtlin@gmail.com

## Version
0.0.1





