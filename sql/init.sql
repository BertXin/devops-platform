-- 创建数据库
CREATE DATABASE IF NOT EXISTS `devops_platform` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE `devops_platform`;

-- 1. 用户表
CREATE TABLE `user` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `uid` VARCHAR(255) DEFAULT NULL COMMENT '用户uid',
  `username` VARCHAR(255) NOT NULL COMMENT '用户名',
  `password` VARCHAR(255) NOT NULL COMMENT '用户密码',
  `phone` VARCHAR(20) DEFAULT NULL COMMENT '手机号码',
  `email` VARCHAR(255) DEFAULT NULL COMMENT '邮箱',
  `nickname` VARCHAR(255) DEFAULT NULL COMMENT '用户昵称',
  `avatar` VARCHAR(500) DEFAULT 'https://www.dnsjia.com/luban/img/head.png' COMMENT '用户头像',
  `status` TINYINT(1) DEFAULT 1 COMMENT '用户状态(1:正常 0:禁用)',
  `mfa_secret` TEXT DEFAULT NULL COMMENT 'mfa密钥',
  `role_id` BIGINT DEFAULT NULL COMMENT '角色id外键',
  `dept_id` BIGINT DEFAULT NULL COMMENT '部门id外键',
  `title` VARCHAR(255) DEFAULT NULL COMMENT '职位',
  `create_by` VARCHAR(50) DEFAULT NULL COMMENT '创建来源,ldap/local/dingtalk',
  `password_updated` DATETIME DEFAULT NULL COMMENT '密码更新时间',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by_id` BIGINT DEFAULT 0 COMMENT '创建人ID',
  `created_by_name` VARCHAR(255) DEFAULT '系统' COMMENT '创建人姓名',
  `last_modified_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  `last_modified_by_id` BIGINT DEFAULT 0 COMMENT '最后修改人ID',
  `last_modified_by_name` VARCHAR(255) DEFAULT '系统' COMMENT '最后修改人姓名',
  `updated_at` DATETIME DEFAULT NULL COMMENT '最后登录时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- 2. 角色表
CREATE TABLE `role` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `name` VARCHAR(128) NOT NULL COMMENT '角色名称',
  `code` VARCHAR(64) NOT NULL COMMENT '角色唯一标识符',
  `description` VARCHAR(255) DEFAULT NULL COMMENT '角色描述',
  `status` TINYINT DEFAULT 1 COMMENT '状态 1:启用 0:禁用',
  `sort_order` INT DEFAULT 0 COMMENT '排序',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by_id` BIGINT DEFAULT 0 COMMENT '创建人ID',
  `created_by_name` VARCHAR(255) DEFAULT '系统' COMMENT '创建人姓名',
  `last_modified_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  `last_modified_by_id` BIGINT DEFAULT 0 COMMENT '最后修改人ID',
  `last_modified_by_name` VARCHAR(255) DEFAULT '系统' COMMENT '最后修改人姓名',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色表';

-- 3. 权限表
CREATE TABLE `permission` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '权限ID',
  `parent_id` BIGINT DEFAULT 0 COMMENT '父权限ID',
  `name` VARCHAR(128) NOT NULL COMMENT '权限名称',
  `type` VARCHAR(20) NOT NULL COMMENT '权限类型: menu, api, button',
  `path` VARCHAR(200) DEFAULT NULL COMMENT '路径',
  `method` VARCHAR(20) DEFAULT NULL COMMENT 'HTTP方法',
  `icon` VARCHAR(128) DEFAULT NULL COMMENT '图标',
  `component` VARCHAR(128) DEFAULT NULL COMMENT '组件路径',
  `permission` VARCHAR(128) DEFAULT NULL COMMENT '权限标识',
  `status` TINYINT DEFAULT 1 COMMENT '状态 1:启用 0:禁用',
  `hidden` TINYINT(1) DEFAULT 0 COMMENT '是否隐藏',
  `sort_order` INT DEFAULT 0 COMMENT '排序',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by_id` BIGINT DEFAULT 0 COMMENT '创建人ID',
  `created_by_name` VARCHAR(255) DEFAULT '系统' COMMENT '创建人姓名',
  `last_modified_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  `last_modified_by_id` BIGINT DEFAULT 0 COMMENT '最后修改人ID',
  `last_modified_by_name` VARCHAR(255) DEFAULT '系统' COMMENT '最后修改人姓名',
  PRIMARY KEY (`id`),
  KEY `idx_parent_id` (`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='权限表';

-- 4. 角色权限关联表
CREATE TABLE `role_permission` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `role_id` BIGINT NOT NULL COMMENT '角色ID',
  `permission_id` BIGINT NOT NULL COMMENT '权限ID',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_role_perm` (`role_id`, `permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色权限关联表';

-- 5. 用户角色关联表
CREATE TABLE `user_role` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` BIGINT NOT NULL COMMENT '用户ID',
  `role_id` BIGINT NOT NULL COMMENT '角色ID',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_role` (`user_id`, `role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户角色关联表';

-- 6. Casbin规则表
CREATE TABLE `casbin_rule` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `p_type` VARCHAR(100) DEFAULT NULL COMMENT '策略类型',
  `v0` VARCHAR(100) DEFAULT NULL COMMENT '角色或用户',
  `v1` VARCHAR(100) DEFAULT NULL COMMENT '资源',
  `v2` VARCHAR(100) DEFAULT NULL COMMENT '操作',
  `v3` VARCHAR(100) DEFAULT NULL COMMENT '域',
  `v4` VARCHAR(100) DEFAULT NULL COMMENT '策略规则',
  `v5` VARCHAR(100) DEFAULT NULL COMMENT '扩展字段',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Casbin规则表';

-- 7. 部门表
CREATE TABLE `department` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '部门ID',
  `parent_id` BIGINT DEFAULT 0 COMMENT '父部门ID',
  `name` VARCHAR(128) NOT NULL COMMENT '部门名称',
  `code` VARCHAR(64) NOT NULL COMMENT '部门编码',
  `description` VARCHAR(255) DEFAULT NULL COMMENT '描述',
  `status` TINYINT DEFAULT 1 COMMENT '状态 1:启用 0:禁用',
  `sort` INT DEFAULT 0 COMMENT '排序',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by_id` BIGINT DEFAULT 0 COMMENT '创建人ID',
  `created_by_name` VARCHAR(255) DEFAULT '系统' COMMENT '创建人姓名',
  `last_modified_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  `last_modified_by_id` BIGINT DEFAULT 0 COMMENT '最后修改人ID',
  `last_modified_by_name` VARCHAR(255) DEFAULT '系统' COMMENT '最后修改人姓名',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_code` (`code`),
  KEY `idx_parent_id` (`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='部门表';

-- 8. 用户部门关联表
CREATE TABLE `user_department` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` BIGINT NOT NULL COMMENT '用户ID',
  `department_id` BIGINT NOT NULL COMMENT '部门ID',
  `type` INT DEFAULT 0 COMMENT '关系类型 0:普通成员 1:负责人',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_dept` (`user_id`, `department_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户部门关联表';

-- 9. 应用表
CREATE TABLE `app` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '应用ID',
  `name` VARCHAR(100) NOT NULL COMMENT '应用名称',
  `description` VARCHAR(500) DEFAULT NULL COMMENT '应用描述',
  `creator` BIGINT NOT NULL COMMENT '创建者ID',
  `status` VARCHAR(20) NOT NULL DEFAULT 'active' COMMENT '应用状态',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by_id` BIGINT DEFAULT 0 COMMENT '创建人ID',
  `created_by_name` VARCHAR(255) DEFAULT '系统' COMMENT '创建人姓名',
  `last_modified_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  `last_modified_by_id` BIGINT DEFAULT 0 COMMENT '最后修改人ID',
  `last_modified_by_name` VARCHAR(255) DEFAULT '系统' COMMENT '最后修改人姓名',
  `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_name` (`name`),
  KEY `idx_creator` (`creator`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='应用表';

-- 10. 应用分组表
CREATE TABLE `app_group` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '应用分组ID',
  `name` VARCHAR(100) NOT NULL COMMENT '分组名称',
  `description` VARCHAR(500) DEFAULT NULL COMMENT '分组描述',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by_id` BIGINT DEFAULT 0 COMMENT '创建人ID',
  `created_by_name` VARCHAR(255) DEFAULT '系统' COMMENT '创建人姓名',
  `last_modified_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  `last_modified_by_id` BIGINT DEFAULT 0 COMMENT '最后修改人ID',
  `last_modified_by_name` VARCHAR(255) DEFAULT '系统' COMMENT '最后修改人姓名',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='应用分组表';

-- 11. 应用分组关联表
CREATE TABLE `relation_app_group_app` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `group_id` BIGINT NOT NULL COMMENT '分组ID',
  `app_id` BIGINT NOT NULL COMMENT '应用ID',
  PRIMARY KEY (`id`),
  KEY `idx_group_id` (`group_id`),
  KEY `idx_app_id` (`app_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='应用分组关联表';

-- 12. 应用环境表
CREATE TABLE `app_env` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '环境ID',
  `name` VARCHAR(100) NOT NULL COMMENT '环境名称',
  `cluster_id` BIGINT NOT NULL COMMENT '集群ID',
  `namespace` VARCHAR(100) NOT NULL COMMENT '命名空间',
  `description` VARCHAR(500) DEFAULT NULL COMMENT '环境描述',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by_id` BIGINT DEFAULT 0 COMMENT '创建人ID',
  `created_by_name` VARCHAR(255) DEFAULT '系统' COMMENT '创建人姓名',
  `last_modified_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  `last_modified_by_id` BIGINT DEFAULT 0 COMMENT '最后修改人ID',
  `last_modified_by_name` VARCHAR(255) DEFAULT '系统' COMMENT '最后修改人姓名',
  PRIMARY KEY (`id`),
  KEY `idx_cluster_id` (`cluster_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='应用环境表';

-- 13. 应用HPA配置表
CREATE TABLE `app_hpa` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT 'HPA配置ID',
  `app_id` BIGINT NOT NULL COMMENT '应用ID',
  `min_replicas` INT NOT NULL DEFAULT 1 COMMENT '最小副本数',
  `max_replicas` INT NOT NULL DEFAULT 10 COMMENT '最大副本数',
  `target_cpu` INT NOT NULL DEFAULT 80 COMMENT '目标CPU使用率',
  `target_memory` INT DEFAULT 0 COMMENT '目标内存使用率',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by_id` BIGINT DEFAULT 0 COMMENT '创建人ID',
  `created_by_name` VARCHAR(255) DEFAULT '系统' COMMENT '创建人姓名',
  `last_modified_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  `last_modified_by_id` BIGINT DEFAULT 0 COMMENT '最后修改人ID',
  `last_modified_by_name` VARCHAR(255) DEFAULT '系统' COMMENT '最后修改人姓名',
  PRIMARY KEY (`id`),
  KEY `idx_app_id` (`app_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='应用HPA配置表';

-- 14. 镜像仓库表
CREATE TABLE `image_registry` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '镜像仓库ID',
  `name` VARCHAR(100) NOT NULL COMMENT '仓库名称',
  `url` VARCHAR(200) NOT NULL COMMENT '仓库地址',
  `username` VARCHAR(100) DEFAULT NULL COMMENT '用户名',
  `password` VARCHAR(100) DEFAULT NULL COMMENT '密码',
  `email` VARCHAR(200) DEFAULT NULL COMMENT '邮箱',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by_id` BIGINT DEFAULT 0 COMMENT '创建人ID',
  `created_by_name` VARCHAR(255) DEFAULT '系统' COMMENT '创建人姓名',
  `last_modified_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  `last_modified_by_id` BIGINT DEFAULT 0 COMMENT '最后修改人ID',
  `last_modified_by_name` VARCHAR(255) DEFAULT '系统' COMMENT '最后修改人姓名',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='镜像仓库表';

-- 15. 应用镜像仓库关联表
CREATE TABLE `app_image_registry` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `app_id` BIGINT NOT NULL COMMENT '应用ID',
  `registry_id` BIGINT NOT NULL COMMENT '镜像仓库ID',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by_id` BIGINT DEFAULT 0 COMMENT '创建人ID',
  `created_by_name` VARCHAR(255) DEFAULT '系统' COMMENT '创建人姓名',
  `last_modified_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  `last_modified_by_id` BIGINT DEFAULT 0 COMMENT '最后修改人ID',
  `last_modified_by_name` VARCHAR(255) DEFAULT '系统' COMMENT '最后修改人姓名',
  PRIMARY KEY (`id`),
  KEY `idx_app_id` (`app_id`),
  KEY `idx_registry_id` (`registry_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='应用镜像仓库关联表';

-- 16. 发布计划表
CREATE TABLE `app_release_plan` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '发布计划ID',
  `app_id` BIGINT NOT NULL COMMENT '应用ID',
  `env_id` BIGINT NOT NULL COMMENT '环境ID',
  `version` VARCHAR(50) NOT NULL COMMENT '版本号',
  `strategy` VARCHAR(50) NOT NULL DEFAULT 'rolling' COMMENT '发布策略',
  `status` VARCHAR(20) NOT NULL DEFAULT 'pending' COMMENT '发布状态',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by_id` BIGINT DEFAULT 0 COMMENT '创建人ID',
  `created_by_name` VARCHAR(255) DEFAULT '系统' COMMENT '创建人姓名',
  `last_modified_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  `last_modified_by_id` BIGINT DEFAULT 0 COMMENT '最后修改人ID',
  `last_modified_by_name` VARCHAR(255) DEFAULT '系统' COMMENT '最后修改人姓名',
  PRIMARY KEY (`id`),
  KEY `idx_app_id` (`app_id`),
  KEY `idx_env_id` (`env_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='发布计划表';

-- 17. 部署历史表
CREATE TABLE `deploy_history` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '部署ID',
  `app_id` BIGINT NOT NULL COMMENT '应用ID',
  `env_id` BIGINT NOT NULL COMMENT '环境ID',
  `version` VARCHAR(50) NOT NULL COMMENT '版本号',
  `status` VARCHAR(20) NOT NULL DEFAULT 'pending' COMMENT '部署状态',
  `start_time` DATETIME NOT NULL COMMENT '开始时间',
  `end_time` DATETIME DEFAULT NULL COMMENT '结束时间',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by_id` BIGINT DEFAULT 0 COMMENT '创建人ID',
  `created_by_name` VARCHAR(255) DEFAULT '系统' COMMENT '创建人姓名',
  `last_modified_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  `last_modified_by_id` BIGINT DEFAULT 0 COMMENT '最后修改人ID',
  `last_modified_by_name` VARCHAR(255) DEFAULT '系统' COMMENT '最后修改人姓名',
  PRIMARY KEY (`id`),
  KEY `idx_app_id` (`app_id`),
  KEY `idx_env_id` (`env_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='部署历史表';

-- 18. 部署步骤表
CREATE TABLE `deploy_step` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '步骤ID',
  `deploy_id` BIGINT NOT NULL COMMENT '部署ID',
  `name` VARCHAR(100) NOT NULL COMMENT '步骤名称',
  `status` VARCHAR(20) NOT NULL DEFAULT 'pending' COMMENT '步骤状态',
  `message` VARCHAR(1000) DEFAULT NULL COMMENT '步骤消息',
  `start_time` DATETIME NOT NULL COMMENT '开始时间',
  `end_time` DATETIME DEFAULT NULL COMMENT '结束时间',
  `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `created_by_id` BIGINT DEFAULT 0 COMMENT '创建人ID',
  `created_by_name` VARCHAR(255) DEFAULT '系统' COMMENT '创建人姓名',
  `last_modified_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  `last_modified_by_id` BIGINT DEFAULT 0 COMMENT '最后修改人ID',
  `last_modified_by_name` VARCHAR(255) DEFAULT '系统' COMMENT '最后修改人姓名',
  PRIMARY KEY (`id`),
  KEY `idx_deploy_id` (`deploy_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='部署步骤表';

-- 19. 登录日志表
CREATE TABLE `login_log` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '日志ID',
  `user_id` BIGINT NOT NULL COMMENT '用户ID',
  `username` VARCHAR(255) NOT NULL COMMENT '用户名',
  `ip` VARCHAR(45) DEFAULT NULL COMMENT '登录IP',
  `user_agent` VARCHAR(500) DEFAULT NULL COMMENT '用户代理',
  `login_type` VARCHAR(50) DEFAULT NULL COMMENT '登录类型',
  `status` INT NOT NULL DEFAULT 1 COMMENT '状态 1成功 0失败',
  `message` VARCHAR(255) DEFAULT NULL COMMENT '消息',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_username` (`username`),
  KEY `idx_created_at` (`created_at`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='登录日志表';