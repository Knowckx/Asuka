SET NAMES utf8;  --建立连接的时候设置字符集


-- ----------------------------
-- 标准创建表 SQL模版
-- ----------------------------
CREATE TABLE IF NOT EXISTS `t_customrank_config` (
		`ID` INT NOT NULL auto_increment PRIMARY KEY COMMENT '主键ID',
		`RankName` VARCHAR ( 64 ) NOT NULL COMMENT '排行榜名称',
		`Text` text COMMENT '描述',
		`UserLimit` TINYINT NOT NULL COMMENT '每个用户参赛次数',
		`ManualCfg` INT NULL COMMENT '外键 手动上榜配置',
		`CondCfg` INT NULL COMMENT '外键 条件上榜配置',
		`HideConfig` INT NOT NULL COMMENT '字段隐藏配置',
		`MaxRetracementMax` FLOAT NULL COMMENT '回撤最大值',
		`EquityMin` DOUBLE NULL COMMENT '净值最小值',
		`StartAt` TIMESTAMP NOT NULL COMMENT '比赛开始时间',
		`EndAt` TIMESTAMP NOT NULL COMMENT '比赛结束时间',
		`CreateTime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		`UpdateTime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		UNIQUE `idx_unique_index` ( `RankIndex` ),
		UNIQUE `idx_unique_MT4Account` ( `BrokerID`,`Account` ),
		INDEX `idx_time_search` ( `createtime` )  -- 括号前不要有逗号
) CHARACTER SET = utf8mb4 COMMENT = '细分排行榜配置';





-- 哪个库
using copytrading;

------- 表结构修改修改
-- 加个字段
Alter Table T_table ADD `FilterGradeScore` TEXT  not null COMMENT '等级评分细则'

------- 数据修改
-- 改数据
Update t_customrank_config set RankText = ""  where RankIndex = 3;

加默认数据 
Insert into tableA { (field1,field2…)}  values (value1 , value2…)
	value1 可以为 null ,让DB自己去修正为默认值

