

-- 赛季表
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
		INDEX `idx_unique_mt4account`(`BrokerID`, `Account`) USING BTREE,
		INDEX `idx_unique_createtime`(`CreateTime`) USING BTREE
) CHARACTER SET = utf8mb4 COMMENT = '交易大赛 参赛人员每日状态记录表';


-- 交易大赛 榜单定位器
CREATE TABLE IF NOT EXISTS `t_ltl_rankmap` (
		`ID` INT NOT NULL auto_increment PRIMARY KEY COMMENT '唯一ID',
		`Season` TINYINT NOT NULL COMMENT '大赛赛季序号',
		`RankType` varchar(64)  NOT NULL COMMENT '榜单类型，例如：轻、重、经纪商',
		`TimeIndex` varchar(64) NOT NULL COMMENT '对应大赛的时间周期',
		`StartAt` TIMESTAMP NOT NULL COMMENT '统计周期的起点',
		`EndAt` TIMESTAMP NOT NULL COMMENT '统计周期的终点',
		-- `CalcTime` TIMESTAMP NOT NULL COMMENT '榜单实际生成日期',
		INDEX `idx_unique_rank`(`RankType`, `TimeIndex`) USING BTREE
) CHARACTER SET = utf8mb4 COMMENT = '交易大赛 榜单定位器';

-- 交易大赛 系统生成的榜单
CREATE TABLE IF NOT EXISTS `t_ltl_rankdetail` (
		`ID` INT NOT NULL auto_increment PRIMARY KEY COMMENT '唯一ID',
		`RankmapID` INT NOT NULL COMMENT '榜单定位器ID',

		`BrokerID` int NOT NULL COMMENT '经纪商ID',
		`Account` varchar(64)  NOT NULL COMMENT 'mt4账户',

		`RankIndex` double NOT NULL COMMENT '排名序号',
		`RankScore` double NOT NULL COMMENT '总分',

		`EquityScore` double NOT NULL COMMENT '净值得分',
		`EquityIndex` int NOT NULL COMMENT '净值排名',
		`Equity` double NOT NULL COMMENT '净值',

		`MaxDDScore` double NOT NULL COMMENT '最大回撤得分',
		`MaxDDIndex` int NOT NULL COMMENT '最大回撤排名',
		`MaxDD` double NOT NULL COMMENT '最大回撤',

		`ProfitScore` double NOT NULL COMMENT '利润得分',
		`ProfitIndex` int NOT NULL COMMENT '利润排名',
		`Profit` double NOT NULL COMMENT '利润',

		`ROIScore` double NOT NULL COMMENT '收益率得分',
		`ROIIndex` int NOT NULL COMMENT '收益率排名',
		`ROI` double NOT NULL COMMENT '收益率',
		
		`CreateTime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '计算时间',

		INDEX `idx_unique_mt4account`(`BrokerID`, `Account`) USING BTREE,
		INDEX `idx_unique_createtime`(`CreateTime`) USING BTREE,
		INDEX `idx_unique_rankmap`(`RankmapID`) USING BTREE
) CHARACTER SET = utf8mb4 COMMENT = '交易大赛 榜单详细排名信息';


-- 交易大赛 前端展示的榜单
CREATE TABLE IF NOT EXISTS `t_ltl_rankdisplay` (
		`ID` INT NOT NULL auto_increment PRIMARY KEY COMMENT '唯一ID',
		`RankmapID` INT NOT NULL COMMENT '榜单定位器ID',

		`BrokerID` int NOT NULL COMMENT '经纪商ID',
		`Account` varchar(64)  NOT NULL COMMENT 'mt4账户',

		`RankIndex` double NOT NULL COMMENT '排名序号',
		`Score` double NOT NULL COMMENT '总分',
		`EquityScore` double NOT NULL COMMENT '净值得分',
		`MaxDDScore` double NOT NULL COMMENT '最大回撤得分',
		`ProfitScore` double NOT NULL COMMENT '利润得分',
		`ROIScore` double NOT NULL COMMENT '收益率得分',

		`GradeScore` varchar(64)  NOT NULL COMMENT '账户评级',
		`FoProfits` double NOT NULL COMMENT '跟随者收益',
		`CreateTime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间',

		INDEX `idx_unique_mt4account`(`BrokerID`, `Account`) USING BTREE,
		INDEX `idx_unique_createtime`(`CreateTime`) USING BTREE,
		INDEX `idx_unique_rankmap`(`RankmapID`) USING BTREE
) CHARACTER SET = utf8mb4 COMMENT = '交易大赛 前端展示的榜单';


-- 交易大赛  黑名单
CREATE TABLE IF NOT EXISTS `t_rank_banlist` (
		`ID` INT NOT NULL auto_increment PRIMARY KEY COMMENT '唯一ID',
		`Season` TINYINT NOT NULL COMMENT '赛季序号',
		`BrokerID` int NOT NULL COMMENT '经纪商ID',
		`Account` varchar(64)  NOT NULL COMMENT 'mt4账户',
		`OpType` TINYINT NOT NULL COMMENT '1 加入 -1 移除  0 目前状态为banned',
		`Note` varchar(64) NOT NULL COMMENT '拉黑原因',
		`CreateTime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '操作时间'
) CHARACTER SET = utf8mb4 COMMENT = '交易大赛  黑名单';



