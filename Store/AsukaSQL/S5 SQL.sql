
-- 赛季表
drop table IF EXISTS t_ltl_season;
CREATE TABLE IF NOT EXISTS `t_ltl_season` (
		`ID` INT NOT NULL auto_increment PRIMARY KEY COMMENT '唯一ID',
		`Name` VARCHAR ( 64 ) NOT NULL COMMENT '名称',
		`Season` INT NOT NULL COMMENT '赛季编号',
		`Text` Text COMMENT '大赛描述',
		`Admin` VARCHAR ( 64 ) NOT NULL COMMENT '负责人',
		`SignupStart` TIMESTAMP NOT NULL COMMENT '报名开始时间',
		`SignupEnd` TIMESTAMP NOT NULL COMMENT '报名结束时间',
		`StartAt` TIMESTAMP NOT NULL COMMENT '比赛开始时间',
		`EndAt` TIMESTAMP NOT NULL COMMENT '比赛结束时间',
		`UserLimit` TINYINT NOT NULL COMMENT '每个用户参赛次数',
		`AccountLimit` TINYINT NOT NULL COMMENT '每个账户参赛次数',

		`CreateTime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		`UpdateTime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) CHARACTER SET = utf8mb4 COMMENT = '交易大赛赛季表';

-- 插入一条默认数据
INSERT INTO `datastatistic`.`t_ltl_season`(`ID`, `Name`, `Season`, `Text`, `Admin`, `SignupStart`, `SignupEnd`, `StartAt`, `EndAt`, `UserLimit`, `AccountLimit`, `CreateTime`, `UpdateTime`) VALUES (1, 'S5交易大赛', 5, 'Followme《我是交易员》职业联赛S5赛季', '梁容敏、郑程', '2019-08-29 10:55:04', '2019-10-30 00:00:00', '2019-09-01 00:00:00', '2019-11-30 23:59:59', 2, 1, '2019-08-15 14:20:00', '2019-08-29 10:55:04');



drop table IF EXISTS t_ltl_rankmap;
-- 交易大赛 榜单定位器
CREATE TABLE IF NOT EXISTS `t_ltl_rankmap` (
		`ID` INT NOT NULL auto_increment PRIMARY KEY COMMENT '唯一ID',
		`Season` TINYINT NOT NULL COMMENT '大赛赛季序号',
		`RankType` varchar(64)  NOT NULL COMMENT '榜单类型，例如：轻、重、经纪商',
		`TimeIndex` varchar(64) NOT NULL COMMENT '对应大赛的时间周期',
		`StartAt` TIMESTAMP NOT NULL COMMENT '统计周期的起点',
		`EndAt` TIMESTAMP NOT NULL COMMENT '统计周期的终点',
		-- `CalcTime` TIMESTAMP NOT NULL COMMENT '榜单实际生成日期',
		INDEX `idx_rank`(`RankType`, `TimeIndex`) USING BTREE
) CHARACTER SET = utf8mb4 COMMENT = '交易大赛 榜单定位器';

-- 交易大赛 参赛人员每日状态记录表
CREATE TABLE IF NOT EXISTS `t_ltl_daily` (
		`ID` INT NOT NULL auto_increment PRIMARY KEY COMMENT '唯一ID',
		`Season` TINYINT NOT NULL COMMENT '赛季序号',
		`BrokerID` int NOT NULL COMMENT '经纪商ID',
		`Account` varchar(64)  NOT NULL COMMENT 'mt4账户',

		`GradeScore` varchar(64)  NOT NULL COMMENT '账户评级',
		`Equity` double NOT NULL COMMENT '净值',
		`Profit` double NOT NULL COMMENT '利润金额',
		`ROI` double NOT NULL COMMENT '收益率',
		`MaxDD` double NOT NULL COMMENT '最大回撤',
		
		`CreateTime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		INDEX `idx_mt4account`(`BrokerID`, `Account`) USING BTREE,
		INDEX `idx_createtime`(`CreateTime`) USING BTREE
) CHARACTER SET = utf8mb4 COMMENT = '交易大赛 参赛人员每日状态记录表';

-- 交易大赛 系统生成的榜单
CREATE TABLE IF NOT EXISTS `t_ltl_rankdetail` (
		`ID` INT NOT NULL auto_increment PRIMARY KEY COMMENT '唯一ID',
		`RankmapID` INT NOT NULL COMMENT '榜单定位器ID',

		`BrokerID` int NOT NULL COMMENT '经纪商ID',
		`Account` varchar(64)  NOT NULL COMMENT 'mt4账户',

		`RankIndex` int NOT NULL COMMENT '排名序号',
		`RankScore` double NOT NULL COMMENT '总分',

		`EquityIndex` int NOT NULL COMMENT '净值排名',
		`EquityScore` double NOT NULL COMMENT '净值得分',
		`Equity` double NOT NULL COMMENT '净值',

		`MaxDDIndex` int NOT NULL COMMENT '最大回撤排名',
		`MaxDDScore` double NOT NULL COMMENT '最大回撤得分',
		`MaxDD` double NOT NULL COMMENT '最大回撤',

		`ProfitIndex` int NOT NULL COMMENT '利润排名',
		`ProfitScore` double NOT NULL COMMENT '利润得分',
		`Profit` double NOT NULL COMMENT '利润',

		`ROIIndex` int NOT NULL COMMENT '收益率排名',
		`ROIScore` double NOT NULL COMMENT '收益率得分',
		`ROI` double NOT NULL COMMENT '收益率',
		
		`CreateTime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '计算时间',

		INDEX `idx_mt4account`(`BrokerID`, `Account`) USING BTREE,
		INDEX `idx_createtime`(`CreateTime`) USING BTREE,
		INDEX `idx_rankmap`(`RankmapID`) USING BTREE
) CHARACTER SET = utf8mb4 COMMENT = '交易大赛 榜单详细排名信息';


drop table IF EXISTS t_ltl_rankdisplay;
-- 交易大赛 前端展示的榜单
CREATE TABLE IF NOT EXISTS `t_ltl_rankdisplay` (
		`ID` INT NOT NULL auto_increment PRIMARY KEY COMMENT '唯一ID',
		`RankmapID` INT NOT NULL COMMENT '榜单定位器ID',

		`BrokerID` int NOT NULL COMMENT '经纪商ID',
		`Account` varchar(64)  NOT NULL COMMENT 'mt4账户',

		`RankIndex` int NOT NULL COMMENT '排名序号',
		`RankScore` double NOT NULL COMMENT '总分',
		`EquityScore` double NOT NULL COMMENT '净值得分',
		`MaxDDScore` double NOT NULL COMMENT '最大回撤得分',
		`ProfitScore` double NOT NULL COMMENT '利润得分',
		`ROIScore` double NOT NULL COMMENT '收益率得分',

		`GradeScore` varchar(64)  NOT NULL COMMENT '账户评级',
		`FoProfits` double NOT NULL COMMENT '跟随者收益',
		`CreateTime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间',

		INDEX `idx_mt4account`(`BrokerID`, `Account`) USING BTREE,
		INDEX `idx_createtime`(`CreateTime`) USING BTREE,
		INDEX `idx_rankmap`(`RankmapID`) USING BTREE
) CHARACTER SET = utf8mb4 COMMENT = '交易大赛 前端展示的榜单';


-- 交易大赛  黑名单
CREATE TABLE IF NOT EXISTS `t_ltl_banlist` (
		`ID` INT NOT NULL auto_increment PRIMARY KEY COMMENT '唯一ID',
		`Season` TINYINT NOT NULL COMMENT '赛季序号',
		`BrokerID` int NOT NULL COMMENT '经纪商ID',
		`Account` varchar(64)  NOT NULL COMMENT 'mt4账户',
		`OpType` TINYINT NOT NULL COMMENT '1 加入 -1 移除  0 目前状态为banned',
		`Note` varchar(64) NOT NULL COMMENT '拉黑原因',
		`CreateTime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '操作时间'
) CHARACTER SET = utf8mb4 COMMENT = '交易大赛  黑名单';


-- 报名账户的历史净值
DROP TABLE IF EXISTS `t_ltl_s4_hisequtiy`;
CREATE TABLE `t_ltl_EquityHistroy`  (
  `ID` INT NOT NULL auto_increment PRIMARY KEY COMMENT '唯一ID',
  `Season` TINYINT NOT NULL COMMENT '赛季序号',
  `BrokerID` INT NOT NULL COMMENT '经纪商ID',
  `Account` VARCHAR ( 64 ) NOT NULL COMMENT 'mt4账户',
  `Equity` DOUBLE NOT NULL COMMENT '净值',
  `CreateTime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	INDEX `idx_mt4account` (`BrokerID`, `Account`),
	INDEX `idx_createtime` (`CreateTime`)
) CHARACTER  SET = utf8mb4   COMMENT = '交易大赛 报名账户的历史净值';



-- S4大赛之前的已经失效的旧表数据
DROP TABLE IF EXISTS `t_ltl_day`;
DROP TABLE IF EXISTS `t_ltl_week`;
DROP TABLE IF EXISTS `t_ltl_month`;
