SET NAMES utf8;  --建立连接的时候设置字符集

-- ----------------------------
-- 标准创建表 SQL模版
-- ----------------------------
CREATE TABLE IF NOT EXISTS `t_customrank_config` (
		`ID` INT NOT NULL auto_increment PRIMARY KEY COMMENT '排行榜唯一ID',
		`RankName` VARCHAR ( 64 ) NOT NULL COMMENT '排行榜名称',
		`RankText` VARCHAR ( 128 ) COMMENT '排行榜描述',
		`RankIndex` INT NOT NULL COMMENT '排行榜的位置序号',
		`ListType` INT NOT NULL COMMENT '上榜方式',
		`ManualCfg` INT NULL COMMENT '外键 手动上榜配置',
		`CondCfg` INT NULL COMMENT '外键 条件上榜配置',
		`HideConfig` INT NOT NULL COMMENT '字段隐藏配置',
		`CreateTime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		`UpdateTime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		UNIQUE `idx_unique_index` ( `RankIndex` )
	) CHARACTER SET = utf8mb4 
  COMMENT = '细分排行榜配置';

