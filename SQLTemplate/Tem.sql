SET NAMES utf8;  --建立连接的时候设置字符集

-- ----------------------------
-- 标准创建表 SQL模版
-- ----------------------------
CREATE TABLE IF NOT EXISTS `t_broker_overnightrate_record`  (
  `ID` int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY, -- 经纪商隔夜利息记录表
  `BrokerID` int(11) NOT NULL, --经纪商ID
  `Symbol` varchar(64) NOT NULL, --交易品种
  `RecodeTime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP , --记录日期   

  `OrderID` int(11) NOT NULL,
  `Swaptype` varchar(64)   NOT NULL, -- 隔夜利息计算模式
  `Swaplong` varchar(64)   NOT NULL, -- 买单利息
  `Swapshort` varchar(64)   NOT NULL -- 卖单利息
  INDEX `idx_order_search`(`BrokerID`, `Account`, `OrderID`) USING BTREE --索引设置
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci;  


/*
COLLATE = utf8_general_ci   
  ci是case insensitive的缩写，意思是大小写不敏感
  utf8_general_ci本身是一种排序方式
*/


/*
更新时间的字段  不为空，默认值是现在的时间，自动更新本条时间
  `UpdateTime`   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  utf8_general_ci本身是一种排序方式

时间可为空值
  `EndTime`        TIMESTAMP(3) NULL DEFAULT NULL
*/

