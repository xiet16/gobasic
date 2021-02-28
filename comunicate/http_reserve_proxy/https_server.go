package http_reserve_proxy

/*
https 是http协议的安全版本，http 协议是明文的，不安全的， https 使用了ssl/tls 协议进行了加密处理
http 和https 使用的连接方式也不同，默认端口也不同

http 在三次握手后，直接请求和响应
https 需要进行ssl握手，ca 机构
其实就是交换公钥密钥的过程，然后使用对方的密钥就行加密数据
*/
