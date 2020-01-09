
-- 解绑快照
CREATE TABLE IF NOT EXISTS `t_Unbind_Snap` (
		`ID` INT NOT NULL auto_increment PRIMARY KEY COMMENT '唯一ID',
		`BrokerID` int NOT NULL COMMENT '经纪商ID',
		`Account` varchar(64)  NOT NULL COMMENT 'mt4账户',
		`UserID` int NOT NULL COMMENT 'FM用户ID',
		`AccountIndex` int NOT NULL COMMENT '账户index',
        `NickName` varchar(64)  NOT NULL COMMENT '昵称',
		`UserType` tinyint NOT NULL COMMENT '1：交易员  2：跟随者',
		`AccountType` tinyint NOT NULL COMMENT '0：Demo账户 1：真实账户 2：sam 账户 3：mam 账户',
		`StartTime` TIMESTAMP NOT NULL,
		`EndTime` TIMESTAMP NOT NULL,
		`GSInfo` text  NOT NULL COMMENT '账户评级',
		`DataOverview` text  NOT NULL COMMENT '数据概览',
		`ExtraData` text  NOT NULL COMMENT '未包含在数据概览里额外字段值',
		`CreateTime` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		INDEX `idx_fmuser`(`UserID`) USING BTREE,
		INDEX `idx_createtime`(`CreateTime`) USING BTREE,
		UNIQUE `idx_unique_Snap` ( `StartTime`,`BrokerID`,`Account`)
) CHARACTER SET = utf8mb4 COMMENT = '解绑快照';

