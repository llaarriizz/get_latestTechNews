# get_latestTechNews

A simple Go CLI tool that fetches and displays the latest tech news from [TechCrunch RSS feed](https://techcrunch.com/feed/).

## Usage

```bash
go run main.go
```

Or build and run:

```bash
go build -o get_latestTechNews
./get_latestTechNews
```

Running the program will also automatically add the `news` alias to your shell config (`~/.bashrc` or `~/.zshrc`) if it's not already there. After running once, you can reload your shell or start a new terminal and use:

```bash
news
```

## Output

Displays each article as:
```
$ Title
- Description
```

## Dependencies

- Go 1.27.1+
