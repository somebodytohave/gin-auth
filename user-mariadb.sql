CREATE TABLE `user` (
`id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
`nickname` varchar(50) NULL COMMENT '昵称',
`created_at` datetime NULL ON UPDATE CURRENT_TIMESTAMP,
`updated_at` datetime NULL ON UPDATE CURRENT_TIMESTAMP,
`deleted_at` datetime NULL ON UPDATE CURRENT_TIMESTAMP,
`status` tinyint UNSIGNED NULL DEFAULT 1,
PRIMARY KEY (`id`) 
);
CREATE TABLE `user_oauth` (
`id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
`user_id` int(11) UNSIGNED NULL COMMENT '用户信息所对应的id',
`oauth_type` tinyint(3) UNSIGNED NOT NULL COMMENT '1QQ,2微信,3.Github',
`oauth_id` varchar(100) NOT NULL COMMENT '第三方唯一标识符',
`oauth_access_token` varchar(100) NULL,
`oauth_expires` int(11) UNSIGNED NULL COMMENT '认证过期时间',
`status` tinyint UNSIGNED NULL DEFAULT 1,
PRIMARY KEY (`id`) ,
UNIQUE INDEX `index_oauth_id_unique` (`oauth_id` ASC) USING BTREE COMMENT '第三方唯一标识id索引' 
)
COMMENT = '第三方认证';
CREATE TABLE `user_login` (
`id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
`user_id` int(11) UNSIGNED NULL,
`login_name` varchar(30) NULL COMMENT '用户名',
`login_phone` char(11) NULL COMMENT '登录手机号',
`login_email` varchar(30) NULL COMMENT '登录邮箱',
`password` varchar(60) NULL COMMENT '密码',
`status` tinyint(3) UNSIGNED NULL DEFAULT 1 COMMENT '登录状态',
PRIMARY KEY (`id`) ,
UNIQUE INDEX `index_user_name_unique` (`login_name` ASC, `login_phone` ASC, `login_email` ASC) USING BTREE COMMENT '用户名唯一' 
);

ALTER TABLE `user_oauth` ADD CONSTRAINT `fk_user_auths_users_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);
ALTER TABLE `user_login` ADD CONSTRAINT `fk_user_login_auths_users_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

