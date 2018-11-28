CREATE TABLE `user` (
`id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
`nickname` varchar(50) NULL COMMENT '昵称',
PRIMARY KEY (`id`) 
);
CREATE TABLE `user_oauth` (
`id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
`user_id` int(11) UNSIGNED NULL COMMENT '用户信息所对应的id',
`oauth_type` tinyint(3) UNSIGNED NULL COMMENT '1QQ,2微信',
`oauth_access_token` varchar(100) NULL,
`oauth_expires` int(11) UNSIGNED NULL COMMENT '认证过期时间',
PRIMARY KEY (`id`) 
)
COMMENT = '第三方认证';
CREATE TABLE `user_login_oauth` (
`id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
`user_id` int(11) UNSIGNED NULL,
`username` varchar(50) NULL COMMENT '用户名',
`password` varchar(50) NULL COMMENT '密码',
`type` tinyint(3) UNSIGNED NULL COMMENT '1手机号2账号3邮箱等自定义',
PRIMARY KEY (`id`) 
);

ALTER TABLE `user_oauth` ADD CONSTRAINT `fk_user_auths_users_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);
ALTER TABLE `user_login_oauth` ADD CONSTRAINT `fk_user_login_auths_users_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

