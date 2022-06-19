

配置maven

配置jdk





# github搜索技巧



开源代码，电子书，开源项目

## （1）按照文件搜索

```
android in:file
```

## (2)按照路径检索

```
andrioid in:path
```

## (4)按照文件大小

```
android size:>100
```

## (5)按照后缀名检索

```
android extention:css
```

## (6)按照是否被fork过

```
android fork:true
```



1. 搜索仅包含 fork过的项目，在搜索条件中加入fork:true或fork:only

示例（搜autotest）：autotest fork:true

2. 搜索项目名称中包含关键字的 in:name

示例(搜tensorflow)：tensorflow in:name

3.搜索描述中包含关键字的 in:description

示例（搜python）：python in:description

4.搜索readme中包含关键字的 in:readme



6.搜索组织或机构名包含ORGNAME的 org:ORGNAME

示例（搜组织名包含github的项目）：org:github 

7.按项目大小搜索 size:n

示例1（搜项目大小超过1MB的）：size:1000

示例2(搜项目大小超过30MB的)：size:>=30000

示例3（搜项目大小小于50KB的）size:<50

示例4（搜项目大小介于50KB到120KB的） size:50..120

8.根据fork数量搜索 fork:n(方法同7)

9.根据star数量搜索 star:n(方法同7)

10.根据编程语言搜索language:LANGUAGE

示例（搜索javascript编写的包含rails的）：rails language:javascript 

11.根据路径搜索 in:path

示例（搜路径中包含tensorflow的）：tensorflow in:path


language:Golang stars:>3000

搜 “CollectionView OR UICollectionView OR collection” 而不是 “collectionView”

示例（搜用户名为defunkt并且forks>100的项目）：user:defunkt forks:>100 

tom location:"San Francisco, CA"	android location:beijing

搜索不包含”cat”的所有仓库：NOT cat

搜索名为”jquery”并库大小在1024至4089KB之间的所有仓库：jquery size:1024..4089

搜索用户名为fengbingchun并且仓库在2019年1月1日后有更新的所有仓库：user:fengbingchun pushed:>2019-01-01

language:javascript location:china 

Awesome windows



**(问题搜索)Issue search** 

A. 搜索用户名为fengbingchun并issue中含有”opencv”字段的所有issues：opencv user:fengbingchun

B. 搜索issue是open状态并且issue中含有”fengbingchun”字段的所有issues：fengbingchun is:open 

C. 搜素issue中comments数大于4次且含有”fengbingchun”字段的所有issues：fengbingchun comments:>4 

D. 搜索issue创建者是fengbingchun的所有issues：author:fengbingchun

E. 搜索issue在2019年2月15日后创建的且含有”opencv”字段的所有issues：opencv created:>2019-03-15

<https://github.com/topics> 

<https://github.com/search/advanced> 

<https://github.com/trending> 

# 《netty实战》





















