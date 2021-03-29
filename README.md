# Ai-tools
可视化操作日常工作场景，自动执行重复代码流程
[ui](https://github.com/yusys-cloud/ai-tools-ui)
## Features
* 自动记录重复工作步骤
* 根据使用频率推荐自动化执行
* 可视化拖拽配置策略
* 故障模拟策略管理
    * 主机列表
    * chaos故障 (故障名称 [限制参数列表])
    * 主机故障策略配置矩阵 (nodes 网络延迟[10秒 port8080])

## Quick Start
- CRUD
``` 
curl -d 'v={"id":"cpu", "name":"CPU实验2"}' -X POST http://localhost:9999/api/kv/chaos/cpu
curl http://localhost:9999/api/kv/chaos
curl http://localhost:9999/api/kv/chaos/cpu
```
## Todo
- Vue 工程创建
- Istio deploy
- Wasm
- Elastisearch集群计算公式:大小，节点，索引由多少分片
- 技术栈清单 lang type name
- 表单设计器,json表单保存，运行
- 数据记录到ai-tools.dat

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