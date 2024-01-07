INSERT INTO roles (role_id,role) VALUES (1,'admin');
INSERT INTO roles (role_id,role) VALUES (2,'hack-portal-operator');
INSERT INTO roles (role_id,role) VALUES (3,'guest');

INSERT INTO rbac_policies (policy_id,p_type, v0, v1, v2, v3) VALUES (1,'p', 3, '/v1/hackathons', 'GET', 'allow');
INSERT INTO rbac_policies (policy_id,p_type, v0, v1, v2, v3) VALUES (2,'p', 3, '/v1/status_tags', 'GET', 'allow');
INSERT INTO rbac_policies (policy_id,p_type, v0, v1, v2, v3) VALUES (3,'p', 2, '/v1/hackathons', 'GET', 'allow');
INSERT INTO rbac_policies (policy_id,p_type, v0, v1, v2, v3) VALUES (4,'p', 2, '/v1/hackathons', 'POST', 'allow');
INSERT INTO rbac_policies (policy_id,p_type, v0, v1, v2, v3) VALUES (5,'p', 2, '/v1/hackathons/*', 'PUT', 'allow');
INSERT INTO rbac_policies (policy_id,p_type, v0, v1, v2, v3) VALUES (6,'p', 2, '/v1/status_tags', 'GET', 'allow');
INSERT INTO rbac_policies (policy_id,p_type, v0, v1, v2, v3) VALUES (7,'p', 2, '/v1/status_tags', 'POST', 'allow');
INSERT INTO rbac_policies (policy_id,p_type, v0, v1, v2, v3) VALUES (8,'p', 2, '/v1/status_tags/*', 'PUT', 'allow');
INSERT INTO rbac_policies (policy_id,p_type, v0, v1, v2, v3) VALUES  (9,'p', 1, '*', '*', 'allow');