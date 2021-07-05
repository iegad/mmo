CREATE TABLE IF NOT EXISTS `DB_BASIC`.`T_USER_INFO`(
  F_AID         BIGINT       NOT NULL AUTO_INCREMENT COMMENT '主键, 自增ID',
  F_EMAIL       VARCHAR(50)  NULL                    COMMENT '邮箱',
  F_PHONE_NUM   VARCHAR(15)  NULL                    COMMENT '手机号',
  F_GENDER      TINYINT      NOT NULL DEFAULT 0      COMMENT '性别',
  F_NICKNAME    VARCHAR(8)   NOT NULL DEFAULT ''     COMMENT '昵称',
  F_AVATOR      VARCHAR(500) NOT NULL DEFAULT ''     COMMENT '头像',
  F_CREATE_TIME BIGINT       NOT NULL DEFAULT 0      COMMENT '创建时间',
  F_LAST_UPDATE BIGINT       NOT NULL DEFAULT 0      COMMENT '最后更新时间',
  F_VER         TEXT         NOT NULL                COMMENT '版本',
  F_VER_CODE    INT UNSIGNED NOT NULL DEFAULT 0      COMMENT '版本号',

  PRIMARY KEY (F_AID),
  UNIQUE KEY (F_EMAIL),
  UNIQUE KEY (F_PHONE_NUM),
  KEY (F_GENDER),
  KEY (F_NICKNAME),
  KEY (F_CREATE_TIME),
  KEY (F_LAST_UPDATE),
  KEY (F_VER_CODE)
)COMMENT = '用户信息表' ENGINE = INNODB CHARSET = utf8mb4;