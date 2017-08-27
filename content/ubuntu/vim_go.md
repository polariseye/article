---
title: "vim搭建go打开发环境"
date: 2017-08-14 09:58:40
draft: true
---

vim搭建go打开发环境
------------------
安装go打vim插件，具体步骤如下：

1. 安装vim插件管理工具
2. 安装go的vim插件
3. 安装相关具体插件

## 安装vim插件管理工具
网上有很多vim插件管理工具，由于第一个接触到的是Vundle，所以我选则安装Vundle

```
git clone https://github.com/gmarik/Vundle.vim.git ~/.vim/bundle/Vundle.vim
```

完成后，调整vim配置：修改~/.vim/.vimrc (如果不存在此文件，则创建)，具体增加以下内容：

```
" -------------  
" Vundle  
" https://github.com/gmarik/Vundle.vim  
" -------------  
  
set nocompatible              " be iMproved, required  
filetype off                  " required  
  
" set the runtime path to include Vundle and initialize  
set rtp+=~/.vim/bundle/Vundle.vim  
call vundle#begin()  
" alternatively, pass a path where Vundle should install plugins  
"call vundle#begin('~/some/path/here')  
  
" let Vundle manage Vundle, required  
Plugin 'gmarik/Vundle.vim'  
  
" The following are examples of different formats supported.  
" Keep Plugin commands between vundle#begin/end.  
" plugin on GitHub repo  
""Plugin 'tpope/vim-fugitive'  
" plugin from http://vim-scripts.org/vim/scripts.html  
""Plugin 'L9'  
" Git plugin not hosted on GitHub  
""Plugin 'git://git.wincent.com/command-t.git'  
" git repos on your local machine (i.e. when working on your own plugin)  
""Plugin 'file:///home/gmarik/path/to/plugin'  
" The sparkup vim script is in a subdirectory of this repo called vim.  
" Pass the path to set the runtimepath properly.  
""Plugin 'rstacruz/sparkup', {'rtp': 'vim/'}  
" Avoid a name conflict with L9  
""Plugin 'user/L9', {'name': 'newL9'}  
  
" Install Vim-go  
Plugin 'fatih/vim-go'  
  
" All of your Plugins must be added before the following line  
call vundle#end()            " required  
filetype plugin indent on    " required  
" To ignore plugin indent changes, instead use:  
"filetype plugin on  
"  
" Brief help  
" :PluginList       - lists configured plugins  
" :PluginInstall    - installs plugins; append `!` to update or just :PluginUpdate  
" :PluginSearch foo - searches for foo; append `!` to refresh local cache  
" :PluginClean      - confirms removal of unused plugins; append `!` to auto-approve removal  
"  
" see :h vundle for more details or wiki for FAQ  
" Put your non-Plugin stuff after this line 
```

## 安装go打vim插件
执行以下命令
```
git clone https://github.com/fatih/vim-go.git ~/.vim/bundle/vim-go
```

然后确保```~/.vim/.vimrc``` 文件包含```Plugin 'fatih/vim-go'```

## 配置具体的相关工具

 打开vim，在命令模式下，输入：```GoInstallBinaries```,等待自动安装完成后，即可

说明：

  如果相关go包下载失败，则可以在[此下载](https://www.golangtc.com/download/package)

# 问题小计

* 环境搭建完成后，使用go build进行编译时，提示，多余的尾部字符

  原因：

  是因为在linux中，每行打行尾部以\n结尾，windows中以\r\n结尾，而在mac中以\r结尾，这导致在linux中进行了错误打识别

  解决办法:

  :set fileformat=unix
