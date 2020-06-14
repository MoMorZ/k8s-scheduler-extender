# 一个简单的K8S调度器

## 来源

云计算实践实验五

## 编译

`go build -o extender`

## 逻辑

### Predicate

或者说是过滤。

随机生成整数进行奇偶性判别，若生成的整数为奇数，则该节点被筛掉。

### Prioritize

根据节点的可用内存和CPU，得到的乘积进行优先级判别。