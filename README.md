## Cross Platform OS Information Package

![go-distroyer](https://raw.githubusercontent.com/zekiahmetbayar/zekiahmetbayar.github.io/master/assets/distroyer.png)

The go-distroyer package provides multi-platform support, making it easy to perform distribution checks in code.

### Installation

```
go get -u "github.com/zekiahmetbayar/go-distroyer"
```

### Usage

```go

package main

import "github.com/zekiahmetbayar/go-distroyer"

func main() {
    // Returns the client's processor architecture
    // ex. x64, arm64...
    arch, err := distroyer.Arch()
    if err != nil {
        log.Fatalf("error when getting arch: %s", err.Error())
    }
    fmt.Printf("Arch: %s\n", arch)

    // Returns the codename of the client's distribution
    // ex. jammy, focal...
    codename, err := distroyer.Codename()
    if err != nil {
        log.Fatalf("error when getting codename: %s", err.Error())
    }
    fmt.Printf("Codename: %s\n", codename)

    // Returns the id like of the client's distribution
    // ex. debian, fedora, rhel...
    like, err := distroyer.Like()
    if err != nil {
        log.Fatalf("error when getting platform: %s", err.Error())
    }
    fmt.Printf("Like: %s\n", like)

    // Returns the pretty name of the client's distribution
    // ex. Ubuntu 20.04 LTS (Focal Fossa)
    pretty, err := distroyer.PrettyName()
    if err != nil {
        log.Fatalf("error when getting pretty name: %s", err.Error())
    }
    fmt.Printf("Pretty Name: %s\n", pretty)

    // Returns the platform of the client's distribution
    // ex. ubuntu, centos, ol...
    platform, err := distroyer.ID()
    if err != nil {
        log.Fatalf("error when getting platform: %s", err.Error())
    }
    fmt.Printf("ID: %s\n", platform)

    // Returns the version of the client's distribution
    // ex. 22.04, 8.6, 6.3...
    version, err := distroyer.ID()
    if err != nil {
        log.Fatalf("error when getting version: %s", err.Error())
    }
    fmt.Printf("Version: %s\n", version)
}

```

### Tested distribution and operating systems

OS/Distro | Arch | Version | Codename | Like | P.Name 
--- | --- | --- | --- |--- |--- 
Ubuntu 18 / 20 / 22 | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: 
Oracle Server 8 | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: 
Centos 7 / 8 | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: 
Pardus 19 / 21 | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: 
Windows 7 / 10 | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark: 