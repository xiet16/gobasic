package tcp_reserve_proxy

/*
tcp 反向代理
开启一个tcp监听，
然后如何将请求转到对应的下游服务器
如何将请求tcp 和代理tcp 关联起来:
1 构建tcp服务器并实现回调
2 构建tcp代理服务器
3 整合tcp 服务器和代理服务器

构建tcp 服务器
1 监听服务
2 获取构建新连接对象并设置超时时间以及keepalive
3 设置方法退出时连接关闭
4 调用回调接口
*/
