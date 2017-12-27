# mqantleaf
leaf迁移到mqant基础架构

可以在不改leaf客户端的情况下将服务端迁移到mqant

## 开发时重点关注


server/gateleaf/routemap.go

## 网关

leaf tcp网关与websocket网关粘包协议不同

tcp 有粘包
websocket 无粘包

导致tcp跟websocket无法用相同的粘包协议，因此分成两个网关



