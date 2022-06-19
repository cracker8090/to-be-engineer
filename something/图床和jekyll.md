# 图床

收费图床

1. 七牛云

2. 又拍云

3. 腾讯云

4. 阿里云

免费图床

1. 微博图床

2. SM.MS图床

3. 路过图床(imgchr)


imgur
聚合图床
Flickr

坚果云

开源图床

[https://imgchr.com](https://imgchr.com/) 

[https://www.superbed.cn/](https://baiyue.one/%e9%83%a8%e8%90%bd%e8%b7%b3%e8%bd%ac.html?url=https://www.superbed.cn/) 聚合图床

博客写作套件**Typora + PicGo + Snipaste**，Typora写文档，Snipaste一键截图，PicGo一键上传图片返回链接，上传 github



[https://typlog.com](https://typlog.com/) 

[plan](https://mednoter.com/plan-for-2020.html) 

# jekyll

<https://knightyun.github.io/2018/04/01/github-pages-blog#5> 

主题网站也是jekyll的，还有一个类似的工具叫“hexo”

| Tables             | Jekyll                                                       | Hexo                                                         |
| ------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| 依赖               | ruby gem （gem依赖也许会带来不兼容问题`弱点`）               | nodejs                                                       |
|                    | `brew install ruby`                                          | `brew install node`                                          |
| 安装               | `gem install jekyll`                                         | `npm install hexo-cli -g`                                    |
| 生成静态站点的速度 | 随着网站内容增加越来越慢 `最大弱点`                          | 相当快 `优点`                                                |
| 与git Pages关系    | 背后运行引擎，支持html/md格式（把原文上传github， 可以直接生成博客，也可以用在线编辑器处理，但只能用[Github-safe plugins](https://help.github.com/articles/adding-jekyll-plugins-to-a-github-pages-site/)）`优点` | 无直接关系（本地生成 html 再上传），部署简单：deploy to Github pages or any other host with one deploy command |
| 格式               | html/md                                                      | md                                                           |
| 是否需本地环境     | 不需                                                         | 需要                                                         |
| 编辑               | 不支持在titles或者YAML使用变量，许多插件过时了`弱点`；不buid-in支持livereload；不buid-in支持post pagination as of Jekyll 3 | 大量可用的[免费插件](https://hexo.io/plugins/) `优点`        |
| 模板               | copy Jekyll创始人的[示例库](https://github.com/mojombo/tpw)，以及其他用Jekyll搭建的[blog](https://github.com/jekyll/jekyll/wiki/Sites) | 大量可用的[免费开源主题](https://hexo.io/themes/) 中国社区很活跃`优点` |
| 迁移               | 依靠[Jekyll importers](https://import.jekyllrb.com/docs/home/)可从其他平台迁移博客（例如WordPress）`优点` | https://hexo.io/zh-cn/docs/migration                         |
| 开源               | 免费开源                                                     | 免费开源                                                     |
| category 分类      | 需要自己写标签语言遍历然后在创建各个分类的主页，再设置页面css，或者用ruby写插件 | 文章前使用“category: 分类名”，自动右边生成，包括分类主页，默认样式 |
| 教程难度           | 相对更难                                                     | 更简洁                                                       |
| 上手难度           | 相对更难                                                     | 更简单                                                       |
