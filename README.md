# Ai-tools
可视化操作日常工作场景，自动执行重复工作流程
[ui](https://github.com/yusys-cloud/ai-tools-ui)
## Features
* [自动流程步骤](base/flow)
  * 命令行执行失败自动重试
  * 文本解析到HTTP接口
* 低代码SSH工具
    * 可视化自动输入Linux命令
    * 可视化拖拽流程策略
* 根据使用频率推荐自动化执行
* [对指定文件夹进行文本内容搜索](base/search)
  * 显示搜索内容上下相关行
  * 对搜索内容进行替换,替换前自动备份文件夹
  * 自定义正则对匹配到的文本进行处理
* [REST-APIs](./docs/rest-api.md)
## Quick Start

``` 
	serv := Server{}
	conf.LoadJsonConfigFile(confName, &serv)
```

## Todo
- 自动记录执行重复工作步骤
  - -i docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=admin -d mysql:5.7
- Vue 工程创建
- Istio deploy
- Wasm
- Elastisearch集群计算公式:大小，节点，索引由多少分片
- 技术栈清单 lang type name
- 表单设计器,json表单保存，运行
- [命令行增加后台执行](https://github.com/alist-org/alist/blob/main/cmd/start.go)

## Technologies
- gin
- nutsdb

## Links 
- [声明式API](https://skyao.io/learning-cloudnative/declarative)
- https://github.com/gostor/awesome-go-storage
- https://github.com/xujiajun/nutsdb/blob/master/README-CN.md
- https://github.com/peterbourgon/diskv
- https://github.com/dgraph-io/badger
- [日志](github.com/sirupsen/logrus)
- [Generate a modern Web project](https://github.com/Shpota/goxygen)
- [drag-and-drop component](https://github.com/SortableJS/Vue.Draggable)
- [form generator and parser](https://github.com/JakHuang/form-generator)