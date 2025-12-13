<p align="center">
<img src="https://raw.githubusercontent.com/panjf2000/logos/master/ants/logo.png" />
<b>A goroutine pool for Go</b>
<br/><br/>
<a title="Build Status" target="_blank" href="https://github.com/panjf2000/ants/actions?query=workflow%3ATests"><img src="https://img.shields.io/github/actions/workflow/status/panjf2000/ants/test.yml?branch=master&style=flat-square&logo=github-actions" /></a>
<a title="Codecov" target="_blank" href="https://codecov.io/gh/panjf2000/ants"><img src="https://img.shields.io/codecov/c/github/panjf2000/ants?style=flat-square&logo=codecov" /></a>
<a title="Release" target="_blank" href="https://github.com/panjf2000/ants/releases"><img src="https://img.shields.io/github/v/release/panjf2000/ants.svg?color=161823&style=flat-square&logo=smartthings" /></a>
<a title="Tag" target="_blank" href="https://github.com/panjf2000/ants/tags"><img src="https://img.shields.io/github/v/tag/panjf2000/ants?color=%23ff8936&logo=fitbit&style=flat-square" /></a>
<br/>
<a title="Minimum Go Version" target="_blank" href="https://github.com/panjf2000/gnet"><img src="https://img.shields.io/badge/go-%3E%3D1.18-30dff3?style=flat-square&logo=go" /></a>
<a title="Go Report Card" target="_blank" href="https://goreportcard.com/report/github.com/panjf2000/ants"><img src="https://goreportcard.com/badge/github.com/panjf2000/ants?style=flat-square" /></a>
<a title="Doc for ants" target="_blank" href="https://pkg.go.dev/github.com/panjf2000/ants/v2?tab=doc"><img src="https://img.shields.io/badge/go.dev-doc-007d9c?style=flat-square&logo=read-the-docs" /></a>
<a title="Mentioned in Awesome Go" target="_blank" href="https://github.com/avelino/awesome-go#goroutines"><img src="https://awesome.re/mentioned-badge-flat.svg" /></a>
</p>

English | [中文](README_ZH.md)

## 📖 Introduction

Library `ants` implements a goroutine pool with fixed capacity, managing and recycling a massive number of goroutines, allowing developers to limit the number of goroutines in your concurrent programs.

## Attribution

This project is derived from:
<https://github.com/panjf2000/ants>

Original Author: Andy Pan
License: MIT

This fork includes significant modifications and refactoring.

## 🚀 Features

- Managing and recycling a massive number of goroutines automatically
- Purging overdue goroutines periodically
- Abundant APIs: submitting tasks, getting the number of running goroutines, tuning the capacity of the pool dynamically, releasing the pool, rebooting the pool, etc.
- Handle panic gracefully to prevent programs from crash
- Efficient in memory usage and it may even achieve ***higher performance*** than unlimited goroutines in Go
- Nonblocking mechanism
- Preallocated memory (ring buffer, optional)

## 💡 How `ants` works

### Flow Diagram

<p align="center">
<img width="1011" alt="ants-flowchart-en" src="https://user-images.githubusercontent.com/7496278/66396509-7b42e700-ea0c-11e9-8612-b71a4b734683.png">
</p>

### Activity Diagrams

![](https://raw.githubusercontent.com/panjf2000/illustrations/master/go/ants-pool-1.png)

![](https://raw.githubusercontent.com/panjf2000/illustrations/master/go/ants-pool-2.png)

![](https://raw.githubusercontent.com/panjf2000/illustrations/master/go/ants-pool-3.png)

![](https://raw.githubusercontent.com/panjf2000/illustrations/master/go/ants-pool-4.png)

## 🧰 How to install

```powershell
go get -u github.com/phamnam2003/ants
```

## 🛠 How to use

Check out [the examples](https://pkg.go.dev/github.com/panjf2000/ants/v2#pkg-examples) for basic usage.

### Functional options for pool

`ants.Options`contains all optional configurations of the ants pool, which allows you to customize the goroutine pool by invoking option functions to set up each configuration in `NewPool`/`NewPoolWithFunc`/`NewPoolWithFuncGeneric` method.

Check out [ants.Options](https://pkg.go.dev/github.com/panjf2000/ants/v2#Options) and [ants.Option](https://pkg.go.dev/github.com/panjf2000/ants/v2#Option) for more details.

### Customize pool capacity

`ants` supports customizing the capacity of the pool. You can call the `NewPool` method to instantiate a `Pool` with a given capacity, as follows:

``` go
p, _ := ants.NewPool(10000)
```

### Submit tasks

Tasks can be submitted by calling `ants.Submit`

```go
ants.Submit(func(){})
```

### Tune pool capacity at runtime

You can tune the capacity of `ants` pool at runtime with `ants.Tune`:

``` go
pool.Tune(1000) // Tune its capacity to 1000
pool.Tune(100000) // Tune its capacity to 100000
```

Don't worry about the contention problems in this case, the method here is thread-safe (or should be called goroutine-safe).

### Pre-malloc goroutine queue in pool

`ants` allows you to pre-allocate the memory of the goroutine queue in the pool, which may get a performance enhancement under some special certain circumstances such as the scenario that requires a pool with ultra-large capacity, meanwhile, each task in goroutine lasts for a long time, in this case, pre-mallocing will reduce a lot of memory allocation in goroutine queue.

```go
// ants will pre-malloc the whole capacity of pool when calling ants.NewPool.
p, _ := ants.NewPool(100000, ants.WithPreAlloc(true))
```

### Release pool

```go
pool.Release()
```

or

```go
pool.ReleaseTimeout(time.Second * 3)
```

### Reboot pool

```go
// A pool that has been released can be still used after calling the Reboot().
pool.Reboot()
```

## ⚙️ About sequence

All tasks submitted to `ants` pool will not be guaranteed to be addressed in order, because those tasks scatter among a series of concurrent workers, thus those tasks would be executed concurrently.

## 👏 Contributors

Please read our [Contributing Guidelines](CONTRIBUTING.md) before opening a PR and thank you to all the developers who already made contributions to `ants`!

<a href="https://github.com/panjf2000/ants/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=panjf2000/ants" />
</a>

## 📄 License

The source code in `ants` is available under the [MIT License](/LICENSE).
