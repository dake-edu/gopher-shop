# Chapter 01: Why Go?

Go was designed to make large software projects easier to build and maintain. Its tools include a compiler, formatter, tests, and a standard library for networking. See the [Go FAQ](https://go.dev/doc/faq) for the language designers’ account.

Compilation does not guarantee that a particular application is faster than one written in another language. Measure the same workload, inputs, environment, and resource limits before comparing performance. Build time also depends on project size and cached work.

Our goal is concrete: read a small program, predict its output, change it, and explain the result. Goroutines will later let us coordinate concurrent work; they still need memory, synchronization, and limits.

## First checkpoint

Describe a book using a title, a price with a currency, and a publication state. Which of these would change after editing the title? Which would change after a discount? We will turn these questions into data and tests.

This is the earlier course. The Russian book and independently executable chapter checkpoints are in [book/](https://github.com/dake-edu/gopher-shop/tree/main/book). The two chapter sequences are different.
