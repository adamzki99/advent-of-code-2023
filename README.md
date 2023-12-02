# Advent of Code 2023

In this year's edition of Advent of Code (AoC), I will continue with the lessons I learned from the [2022 edition](https://github.com/adamzki99/advent-of-code-2022).

Recently, I finished reading "The Clean Coder" by Robert C. Martin, where he speaks highly of test-driven development. So, I will try to achieve as high of a test coverage as possible for this rendition of Advent of Code.
This will be in conjunction with a "proper" development environment. With experience developing within the [Home Assistant](https://www.home-assistant.io/) project, I got familiar with the concept of dev-containers with VS-Code.
So, a dev container will also be included in this repository.

In short, I hope to get familiar with the following:

- Test-driven development
- Development containers
- GO

## Other thoughts

With the wide accessibility of "coding knowledge" through ChatGPT, forums, YouTube, etc. I have found it very relaxing and giving to go back to the documentation and books to find solutions to the issues I encounter. 

Since there is a speed element to the AoC-challenge, I will also allow myself to use ChatGPT/Bard to get some "quick fixes". This will mainly/solely be used to understand error messages spat out while compiling/running the code.

In short, the knowledge sources for this edition will be:

1. [An Introduction to Programming in Go](https://www.golang-book.com/books/intro)
2. ChatGPT/Bard

## Lessons learned

### Day 2

Today, I am closing in to the holy 100% test coverage!

```bash

go test -cover
PASS
coverage: 92.1% of statements
ok      github.com/adamzki99/advent-of-code-2023/day_2  0.003s

```

### Day 1

Today, I only achieved a 54.5% test coverage, but using TDD is fun! It makes one really think about the problem before trying to "throw" a solution at it.

```bash

go test -cover

PASS
coverage: 54.5% of statements
ok      day_1/main      0.004s

```

I have also extended the knowledge sources to include the [official documentation](https://go.dev/).
