SET NAMES utf8;  --建立连接的时候设置字符集

-- ----------------------------
-- 标准创建表 SQL模版
-- ----------------------------
CREATE TABLE IF NOT EXISTS `t_broker` 
( 
  `id`       INT NOT NULL auto_increment PRIMARY KEY COMMENT '经纪商名称',            -- 经纪商隔夜利息记录表 
  `brokerid` INT NOT NULL,                                       COMMENT '经纪商ID',
  `symbol`   VARCHAR (128) NOT NULL,                             COMMENT '交易品种',
  `swaptype` INT NOT NULL,                                       COMMENT '隔夜利息计算模式',
  `swaplong` DOUBLE NOT NULL,                                    COMMENT '买单利息',
  `swapshort` DOUBLE NOT NULL,                                   COMMENT '卖单利息',
  `createtime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,     COMMENT '记录日期'
  `UpdateTime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  
  INDEX `idx_symbol_search` ( `brokerid`, `symbol` ) using btree,-- 索引设置 
  INDEX `idx_time_search` ( `createtime` ) USING btree 
) 
CHARACTER SET = utf8mb4 COLLATE = utf8_general_ci comment='表的注释';

/*
COLLATE = utf8_general_ci   
  ci是case insensitive的缩写，意思是大小写不敏感
  utf8_general_ci本身是一种排序方式
*/

/*
时间可为空值
  `EndTime`        TIMESTAMP(3) NULL DEFAULT NULL
*/

