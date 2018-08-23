SET NAMES utf8;  --建立连接的时候设置字符集

-- ----------------------------
-- 标准创建表 SQL模版
-- ----------------------------
CREATE TABLE IF NOT EXISTS `t_broker_overnightrate_record` 
( 
  `id`       INT NOT NULL auto_increment PRIMARY KEY,            -- 经纪商隔夜利息记录表 
  `brokerid` INT NOT NULL,                                       -- 经纪商ID 
  `symbol`   VARCHAR ( 64 ) NOT NULL,                            -- 交易品种 
  `swaptype` INT NOT NULL,                                       -- 隔夜利息计算模式 
  `swaplong` DOUBLE NOT NULL,                                    -- 买单利息 
  `swapshort` DOUBLE NOT NULL,                                   -- 卖单利息 
  `createtime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,     -- 记录日期 
  
  INDEX `idx_symbol_search` ( `brokerid`, `symbol` ) using btree,-- 索引设置 
  INDEX `idx_time_search` ( `createtime` ) USING btree 
) 
engine = innodb CHARACTER SET = utf8 COLLATE = utf8_general_ci;

 
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

