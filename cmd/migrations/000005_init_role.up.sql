INSERT INTO roles (role_id,role) VALUES (1,'admin');
INSERT INTO roles (role_id,role) VALUES (2,'hack-portal-operator');
INSERT INTO roles (role_id,role) VALUES (3,'guest');

INSERT INTO rbac_policies (policy_id,p_type, v0, v1, v2, v3) VALUES  (1,'p', 1, '*', '*', 'allow');