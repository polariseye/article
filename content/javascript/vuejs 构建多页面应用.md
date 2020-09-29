vuejs构建多页面应用
------------------------------
# 前言
最近打算开发一个工具网站，使它能够不断积累开发所需工具，以方便个人使用。网上查阅资料后，最终打算使用vue来做前端的主要开发框架。原因如下：

1. vue官方文档与相关工具链极其完整，语法简单，入手快
2. 操作domain极其方便 
3. 应用范围广泛，可以使用vue[开发微信小程序](http://mpvue.com)，[手机应用](https://nativescript-vue.org)，[桌面程序 electron-vue](https://github.com/SimulatedGREG/electron-vue)
4. 与webpack集成，代码修改保存后，浏览器会立即刷新，方便好用

# 安装与初始化

1. 安装vue的命令行工具,具体如下
```
vue cli 3使用命令行: npm install -g @vue/cli
vue cli2 2使用命令行安装: npm install -g vue-cli
```
建议使用vue cli 3版本,目测3版本的命令行更好用,多页的模板修改也比较简单   
2. 初始化项目:`vue create testSite`,相关参数默认即可
3. 运行项目:`npm run serve`,然后即可在浏览器里查看网页


# 修改模板为多页面应用

单页面应用在很多场景都无法使用。比如面向搜索时，单页面将会很麻烦。所以建议基于多页面的vue进行开发，具体修改方法参见[官方文档](https://cli.vuejs.org/zh/guide/html-and-static-assets.html#构建一个多页应用)

在项目根目录新建文件:`vue.config.js`,内容可以直接复制以下内容。
```
module.exports = {
    publicPath:"./", // 部署应用包时的基本 URL。用法和 webpack 本身的 output.publicPath 一致，但是 Vue CLI 在一些其他地方也需要用到这个值，所以请始终使用 publicPath 而不要直接修改 webpack 的 output.publicPath。
    outputDir:"dist", // 当运行 vue-cli-service build 时生成的生产环境构建文件的目录,请始终使用 outputDir 而不要修改 webpack 的 output.path
    indexPath:"index.html", // 指定生成的 index.html 的输出路径 (相对于 outputDir)。也可以是一个绝对路径
    filenameHashing:true, // 默认情况下，生成的静态资源在它们的文件名中包含了 hash 以便更好的控制缓存。然而，这也要求 index 的 HTML 是被 Vue CLI 自动生成的。如果你无法使用 Vue CLI 生成的 index HTML，你可以通过将这个选项设为 false 来关闭文件名哈希
    pages: {
      index: {
        // page 的入口
        entry: "src/pages/index/main.js",
        // 模板来源
        template: "src/pages/index/index.html",
        // 在 dist/index.html 的输出
        filename: "index.html",
        // 当使用 title 选项时，
        // template 中的 title 标签需要是 <title><%= htmlWebpackPlugin.options.title %></title>
        title: "今天是个好天气"
      },
      hello: {
        // page 的入口
        entry: "src/pages/helloPage/main.js",
        // 模板来源
        template: "src/pages/helloPage/index.html",
        // 在 dist/index.html 的输出
        filename: "helloPage.html",
        // 当使用 title 选项时，
        // template 中的 title 标签需要是 <title><%= htmlWebpackPlugin.options.title %></title>
        title: "今天是个好天气"
      }
    }
  };
```
说明:
  * 如果需要新增加一个页面，则直接在pages节点下，新增一个节点就好


#参考资料

* [vue cli3超详细创建多页面配置](https://www.cnblogs.com/guiltyWay/p/10320653.html)
* [vue cli 官方教程](https://cli.vuejs.org/zh/guide/html-and-static-assets.html#构建一个多页应用)
* [vue cli配置参考](https://cli.vuejs.org/zh/config/#pages)