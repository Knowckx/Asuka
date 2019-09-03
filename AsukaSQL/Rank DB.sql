
-- 排行榜实例化
CREATE TABLE IF NOT EXISTS `t_Rank_Snap` (
		`ID` INT NOT NULL auto_increment PRIMARY KEY COMMENT '唯一ID',
		`BrokerID` int NOT NULL COMMENT '经纪商ID',
		`Account` varchar(64)  NOT NULL COMMENT 'mt4账户',

		`NickName` varchar(64)  NOT NULL COMMENT '昵称',
		`UserID` int NOT NULL COMMENT 'FM用户ID',
		`AccountIndex` int NOT NULL COMMENT '账户index',

		`ROI` double NOT NULL COMMENT '收益率',
		`Weeks` double NOT NULL COMMENT '交易周期',
		`DealAmount` int NOT NULL COMMENT '交易笔数',
		`Equity` double NOT NULL COMMENT '净值',
		`FollowProfits` double NOT NULL COMMENT '跟随者的收益金额',
		`MaxDD` double NOT NULL COMMENT '最大回撤',
		`GradeScore` varchar(64)  NOT NULL COMMENT '账户评级',
		`Subscribers` int NOT NULL COMMENT '策略市场 订阅人数',
		`SubPrice` double NOT NULL COMMENT '策略市场 订阅价格',
		`FollowMoney` double  NOT NULL COMMENT '实盘跟随总额',
		`CreateTime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		INDEX `idx_mt4account`(`BrokerID`, `Account`) USING BTREE,
		INDEX `idx_createtime`(`CreateTime`) USING BTREE
) CHARACTER SET = utf8mb4 COMMENT = '排行榜数据快照';


