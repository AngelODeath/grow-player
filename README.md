# GrowPlayer

My little golang video player. Built on fyne.

Uses toml configuration file

<img alt="Go" src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white"/>
<img alt="Shell Script" src="https://img.shields.io/badge/shell_script-%23121011.svg?style=for-the-badge&logo=gnu-bash&logoColor=white"/>
<img alt="GitHub Actions" src="https://img.shields.io/badge/githubactions-%232671E5.svg?style=for-the-badge&logo=githubactions&logoColor=white"/>
<img alt="Notion" src="https://img.shields.io/badge/Notion-%23000000.svg?style=for-the-badge&logo=notion&logoColor=white"/>

# Documentation

## Writing a Go App or Library:

Use go to initialize the project modules: 

The following examples show how to set up a blank project properly.

<ins>Command:</ins>

> **go mod**

<ins>Example:</ins>

* To create a new application, initialize a module in an empty directory:
  ```shell
  cd ~/Dev
  mkdir my-example-app
  cd my-example-app
  go mod init my-example-app
  touch main.go
  ```

* To create a new library, initialize a submodule inside the application directory:
  ```shell
  cd ~/Dev/my-example-app
  mkdir libs
  cd libs
  go mod init libs
  touch helpers.go
  ```


## Adding packages:

Use go to add project dependencies (requirements):

The following examples show how to add both public packages and local modules.

### Add a public package residing on github:

<ins>Command:</ins>

  > **go get**

<ins>Package:</ins>

  > github.com/pelletier/go-toml/v2

<ins>Repository:</ins>

  * [go-toml v2](https://github.com/pelletier/go-toml/tree/v2)

<ins>Example:</ins>

  * Add package github.com/pelletier/go-toml/v2 to module dependencies:

    > go get github.com/pelletier/go-toml/v2
    
### Add a local module residing in the libs directory:

This is best done manually by editing the files and adding the following lines:

<ins>Example:</ins>

  * Add these two lines to the file: **go.mod**</ins>

    > replace example.com/yourname/library => ./libs
    >
    > require example.com/yourname/library v0.0.0

  * Add an import line to the file: **main.go**</ins>

    > import libalias "example.com/yourname/library"


## My commands:

Create variable to hold binary filename:
```shell
target_bin=`realpath --relative-to=.. .`
```

build (full):
```shell
go build -o $target_bin -v -x -a .
```

build (quick):
```shell
go build -o $target_bin -v -x .
```

Check binary:
```shell
echo Filename: $target_bin
echo Full path: `realpath ./$target_bin`
echo File size: `du -h ./$target_bin | awk '{print $1}'`
file ./$target_bin
```

Launch:
```shell
./$target_bin
```

Check GOPATH:
```shell
echo Current path: `pwd`
echo GOPATH: $GOPATH
```



## Links and Tutorials:

<ins>Intro to modules on go - part-1:</ins>

<a href="https://dev.to/prassee/intro-to-modules-on-go-part-1-1k77">
<img alt="Dev.to blog" src="https://img.shields.io/badge/dev.to-0A0A0A?style=for-the-badge&logo=dev.to&logoColor=white" >
</a>


<ins>Using Modules and Packages in Go</ins>

<a href="https://levelup.gitconnected.com/using-modules-and-packages-in-go-36a418960556">
levelup.gitconnected.com/using-modules-and-packages-in-go
</a>


## Support me if you like what I do:

<a href="https://ko-fi.com/angelodeath">
<img alt="Github-sponsors" src="https://img.shields.io/badge/sponsor-30363D?style=for-the-badge&logo=GitHub-Sponsors&logoColor=#EA4AAA" />
</a>

## Contact

You can reach me on email:

<a href="mailto:bothagabri@gmail.com?subject=[GitHub]%20GrowPlayer"><img src="https://img.shields.io/badge/gmail-%23DD0031.svg?&style=for-the-badge&logo=gmail&logoColor=white"/></a>
<a href="mailto:angelodeath@outlook.com?subject=[GitHub]%20GrowPlayer"><img alt="Outlook" src="https://img.shields.io/badge/Microsoft_Outlook-0078D4?style=for-the-badge&logo=microsoft-outlook&logoColor=white" /></a>

Or catch me on youtube:

<a href="https://www.youtube.com/channel/UCiXEPksGsnjHrdkJM_BGEKA">
<img alt="YouTube" src="https://img.shields.io/badge/Gabri Botha-%23FF0000.svg?style=for-the-badge&logo=YouTube&logoColor=white"/>
</a>


## License

[![Licence](https://img.shields.io/github/license/Ileriayo/markdown-badges?style=for-the-badge)](./LICENSE)
<br/>
[MIT](./LICENSE)
