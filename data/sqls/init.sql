CREATE DATABASE IF NOT EXISTS casbin;

USE casbin;

CREATE TABLE `casbin_rule` (
    id bigint unsigned AUTO_INCREMENT,
    ptype varchar(100),
    v0 varchar(100),
    v1 varchar(100),
    v2 varchar(100),
    v3 varchar(100),
    v4 varchar(100),
    v5 varchar(100),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	
    PRIMARY KEY (`id`)
);
CREATE UNIQUE INDEX idx_casbin_rule ON casbin_rule (ptype,v0,v1,v2,v3,v4,v5);