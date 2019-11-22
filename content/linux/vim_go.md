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
Plugin 'scrooloose/nerdtree'                                                             
Plugin 'Valloric/YouCompleteMe'  
  
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

" set up Vim-go
let mapleader=";"                                                     
map <leader>n :NERDTreeToggle<CR>                                                         
let g:go_fmt_command = "goimports"                                                       
let g:go_highlight_functions = 1                                                         
let g:go_highlight_methods = 1                                                           
let g:go_highlight_structs = 1                                                          
let g:ycm_add_preview_to_completeopt = 0                                                 
let g:ycm_min_num_of_chars_for_completion = 1                                             
let g:ycm_auto_trigger = 1                                                               
set completeopt-=preview 
```

## 安装go打vim插件
执行以下命令
```
git clone https://github.com/fatih/vim-go.git ~/.vim/bundle/vim-go
```

然后确保```~/.vim/.vimrc``` 文件包含```Plugin 'fatih/vim-go'```

## 配置具体的相关工具

 打开vim，在命令模式下，输入：```PluginInstall```,等待自动安装完成后，即可

说明：

  如果相关go包下载失败，则可以在[此下载](https://www.golangtc.com/download/package)

单独安装工具`YoucompleteMe`
````
yum install vim g++ python build-essential cmake python3-dev
cd ~/.vim/bundle/YouCompleteMe/
./install.py --go-completer

" 临时把GOPATH设置到插件目录
export GOPATH=~/.vim/bundle/YouCompleteMe/third_party/ycmd/third_party/go/
" 安装gocode
go get github.com/mdempsky/gocode
cd github.com/mdempsky/gocode
go build

" 安装godef
go get github.com/rogpeppe/godef
cd ../../../github.com/rogpeppe/godef
go build
````

# 问题小计

* **环境搭建完成后，使用go build进行编译时，提示，多余的尾部字符**

  **原因**：

  是因为在linux中，每行打行尾部以\n结尾，windows中以\r\n结尾，而在mac中以\r结尾，这导致在linux中进行了错误打识别

  **解决办法**:

  :set fileformat=unix

* **打开vim，提示`YouCompleteMe unavailable: requires Vim compiled with Python (2.7.1+ or 3.5.1+) support.`**

  **原因**：

  是因为`YouCompleteMe`需要vim支持python插件。具体可以通过`vim --version|grep python`查看支持情况

  **解决办法**:
 
  使用[vim源码](https://github.com/vim/vim.git)自己安装

  ````
  git clone https://github.com/vim/vim.git

  " 卸载旧版本的vim (由于sudo命令会依赖包vim-minimal，所以源码安装完成后，需要再重新安装sudo命令:yum install sudo)
  yum remove vim* -y

  " 切换到源码目录
  cd vim
  
  " 编译配置具体参数见 ./configure --help
  ./configure --prefix=/usr/local/vim8 --enable-pythoninterp=yes --enable-python3interp=yes --enable-rubyinterp=yes --enable-multibyte=yes   
  
  " 编译 
  make

  " 应用编译结果到--prefix指定的位置
  sudo make install
  ````

* **通过源码安装后，使用`vim --version|grep python`发现vim仍然不支持pytho或python3**

  **原因**：
  
  通过分析执行命令在vim中执行`./configure`时的输出的错误信息`checking if compile and link flags for Python are sane... disabled` 可以看出是vim编译需要安装python的开发版本

  **解决办法**:
  ````
  " 安装python3的开发版本
  yum install python3-dev

  " 安装python2的开发版本
  yum install python-devel
  ````

* **打开vim提示`The ycmd server SHUT DOWN (restart with ...low the instructions in the docu`**

  **原因**：

  是因为`YouCompleteMe`插件没有被完整安装导致

  **解决办法**:

  ````
  " 安装YouCompleteMe需要的其他工具
  yum install vim g++ python build-essential cmake python3-dev

  cd .vim/bundle/YouCompleteMe/

  " 执行实际安装
  ./install.py --go-completer
  ````

* **vim中回退键使用无效**

  **原因**：

  这是因为安装vim插件管理工具`ycm`导致的

  **解决办法**:
 
  在`~/.vimrc`中添加一行配置即可解决
  ````
  set backspace=indent,eol,start
  ````

# 参考资料
* [CentOS7解决YouCompleteMe对Python的依赖](https://blog.csdn.net/uu203/article/details/82621523)
* [vim打造成golang的IDE](https://blog.csdn.net/qq_35976351/article/details/88992931)
* [Centos7 安装vim8.0 并对python2 和 python3的支持](https://blog.csdn.net/lj_trestg/article/details/79754268)