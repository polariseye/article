vue.js学习笔记
-------------------
# 原因
nodejs+electron开发桌面应用是一个不错的选择。但存在一些绕不过的问题。
1. JavaScript语法太随意。纯粹的js的node代码没有api提示。api使用着实不便。此点可以通过使用typescript解决
2. 基于electron的桌面开发严重依赖于nodejs，导致所有代码重用基本变得不可能。有现成的electron-vue框架代码，从而为代码重用做好了基础。

基于以上两个原因，	决定使用vue做为前后端分离的框架


# vue 概要了解
## 兼容性

Vue 不支持 IE8 及以下版本，因为 Vue 使用了 IE8 无法模拟的 ECMAScript 5 特性。但它支持所有兼容 ECMAScript 5 的浏览器。

## 安装
**npm install vue**

## 工具
Vue 提供了一个[官方的 CLI](https://github.com/vuejs/vue-cli)，为单页面应用 (SPA) 快速搭建繁杂的脚手架。它为现代前端工作流提供了 batteries-included 的构建设置。只需要几分钟的时间就可以运行起来并带有热重载、保存时 lint 校验，以及生产环境可用的构建版本。更多详情可查阅[Vue CLI 的文档](https://cli.vuejs.org/)。

# 参考资料
* [vue.js教程](https://cn.vuejs.org/v2/guide/installation.html)