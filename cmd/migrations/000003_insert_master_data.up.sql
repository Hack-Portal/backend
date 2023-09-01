INSERT INTO locates (name) VALUES 
('北海道'), 
('青森県'), 
('岩手県'), 
('宮城県'), 
('秋田県'), 
('山形県'), 
('福島県'), 
('茨城県'), 
('栃木県'), 
('群馬県'), 
('埼玉県'), 
('千葉県'), 
('東京都'), 
('神奈川県'), 
('新潟県'), 
('富山県'), 
('石川県'), 
('福井県'), 
('山梨県'), 
('長野県'), 
('岐阜県'), 
('静岡県'), 
('愛知県'), 
('三重県'), 
('滋賀県'), 
('京都府'), 
('大阪府'), 
('兵庫県'), 
('奈良県'), 
('和歌山県'), 
('鳥取県'), 
('島根県'), 
('岡山県'), 
('広島県'), 
('山口県'), 
('徳島県'), 
('香川県'), 
('愛媛県'), 
('高知県'), 
('福岡県'), 
('佐賀県'), 
('長崎県'), 
('熊本県'), 
('大分県'), 
('宮崎県'), 
('鹿児島県'), 
('沖縄県');
-- テクノロジータグ（tech_tags）のデータを挿入
INSERT INTO tech_tags (language,icon)
VALUES  ('Python','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/tech_tag_python.png?alt=media&token=49c12d37-1de4-4f45-a09b-a7941b88a056'), 
        ('JavaScript','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/tech_tag_js.png?alt=media&token=0d9383a9-9543-4065-a3b8-542705ec5dd2'), 
        ('Java','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/tech_tag_java.png?alt=media&token=eb5b8782-8b37-4b74-a7f8-e27a3bec46dd'), 
        ('Go','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/tech_tag_go.png?alt=media&token=f353e349-6ef8-4475-ae17-ad55a2a583a1'), 
        ('C','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/tech_tag_c.png?alt=media&token=1ace1611-202c-46a4-9591-77795f07457d'), 
        ('Csharp','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/tech_tag_csharp.png?alt=media&token=8af18c6f-03c3-4e28-adfe-ce020020e9fe'), 
        ('Cpp','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/tech_tag_cpp.png?alt=media&token=556d045f-072a-4dc0-8594-9e1872492035'), 
        ('kotlin','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/tech_tag_kotlin.png?alt=media&token=0a8d5d5f-21a8-4812-a50d-e1de99fdf110'), 
        ('PHP','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/tech_tag_php.png?alt=media&token=1f38c3ad-dd5b-4dbb-9122-7ce4f4e9c348'), 
        ('Rust','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/tech_tag_rust.png?alt=media&token=1f3e0982-a3ff-41c6-981f-1c97bfcd1f90'), 
        ('Ruby','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/tech_tag_ruby.png?alt=media&token=7611b57f-f525-44e0-bd8d-7e3bc2deedb7'),
        ('R','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/tech_tag_r.png?alt=media&token=8e7be7d6-68a7-43dd-924b-e16fc72462b7'),
        ('DataBase','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/tech_tag_database.png?alt=media&token=0a1babad-5684-40ff-9922-953b57836fc0'),
        ('Cloud','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/tech_tag_cloud.png?alt=media&token=6ed766e8-bb04-42d4-b2c2-c7418f75a067');
-- Pythonのフレームワーク
INSERT INTO frameworks (tech_tag_id, framework,icon)
VALUES ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Python'), 'Django','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/framewroks_django.png?alt=media&token=ee3d3d5c-4688-47b3-8f3a-fc92540b4265'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Python'), 'Flask','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/framewroks_flask.png?alt=media&token=5e3a25bd-d693-471a-82e3-b14cc0d666ce'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Python'), 'FastAPI','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/framewroks_fastapi.png?alt=media&token=eff5c2c0-e379-4494-b7ec-8ac9cdd3abcc');
-- JavaScriptのフレームワーク
INSERT INTO frameworks (tech_tag_id, framework,icon)
VALUES ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'JavaScript'), 'React.js','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/framewroks_react.png?alt=media&token=5d677a7a-225f-4e89-880f-3ac40fb85aaf'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'JavaScript'), 'Vue.js','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/framewroks_vue.png?alt=media&token=fbc9c33f-bbaa-4f2a-bb12-03f32e7acc70'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'JavaScript'), 'Three.js','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/framewroks_three.png?alt=media&token=f3a3e4d1-ee9a-4c40-b5a0-bddca824afb4'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'JavaScript'), 'Next.js','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/framewroks_next.png?alt=media&token=bd457894-4f00-4c17-8f28-b7f480c38e2f'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'JavaScript'), 'Node.js','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/framewroks_node.png?alt=media&token=6e772370-90b1-459e-be41-9d4e64f8231b'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'JavaScript'), 'Angular','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/framewroks_Anguler.png?alt=media&token=064d7c8d-86ac-43ad-a189-355a54589c39');
-- Javaのフレームワーク
INSERT INTO frameworks (tech_tag_id, framework,icon)
VALUES ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Java'), 'JavaPlayFramework','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/framewroks_play.png?alt=media&token=9b62abac-8051-4873-b62c-884594695796'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Java'), 'Spring','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/framewroks_spring.png?alt=media&token=dca9eef6-6960-4051-b6d9-9e241cc267c6'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Java'), 'ApacheStruts','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/framewroks_struts.png?alt=media&token=8e271d9a-d3f3-4ed4-aa6a-62e0913aa0ca'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Java'), 'JSF','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/framewroks_jsf.png?alt=media&token=52a4046f-08f5-4b3d-afb6-42ea0dd5ea04'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Java'), 'Wicket','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/framewroks_wichet.png?alt=media&token=f8cf244c-09e3-41b0-b147-33ad3de2e801');
-- Goのフレームワーク
INSERT INTO frameworks (tech_tag_id, framework,icon)
VALUES ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Go'), 'Gin','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/frameworks_gin.png?alt=media&token=944c9579-879b-47e8-8f5e-60355657c7f4'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Go'), 'Beego','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/framewroks_Beego.png?alt=media&token=208c8bbe-b2ae-48ec-87b9-95c859e60ce7'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Go'), 'Revel','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/framewroks_go.png?alt=media&token=25dba5a7-1610-446e-a581-67757b7335bd'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Go'), 'Echo','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/framewroks_echo.png?alt=media&token=9543a7af-2252-433d-b7d9-d4f991254377');
-- Cのフレームワーク
INSERT INTO frameworks (tech_tag_id, framework,icon)
VALUES ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'C'), '.NET','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/framewroks_net.png?alt=media&token=da703ee4-6e17-48bb-a712-cc8e50ef0f5d'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'C'), 'ASP.NET','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/framewroks_aspnet.png?alt=media&token=e2b70d2b-616f-40dd-a4d8-1685bd5b4022'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'C'), 'ASP.NET MVC','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/framewroks_aspnet.png?alt=media&token=e2b70d2b-616f-40dd-a4d8-1685bd5b4022');
-- C#のフレームワーク
INSERT INTO frameworks (tech_tag_id, framework,icon)
VALUES ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Csharp'), '.NET','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/framewroks_net.png?alt=media&token=da703ee4-6e17-48bb-a712-cc8e50ef0f5d'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Csharp'), 'ASP.NET','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/framewroks_aspnet.png?alt=media&token=e2b70d2b-616f-40dd-a4d8-1685bd5b4022'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Csharp'), 'ASP.NET MVC','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/framewroks_aspnet.png?alt=media&token=e2b70d2b-616f-40dd-a4d8-1685bd5b4022');
-- C++のフレームワーク
INSERT INTO frameworks (tech_tag_id, framework,icon)
VALUES ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Cpp'), 'Qt','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/framewroks_qt.png?alt=media&token=ebc8ea56-b0e6-4313-8163-ac2e585d1adc');
-- Kotlinのフレームワーク
INSERT INTO frameworks (tech_tag_id, framework,icon)
VALUES ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'kotlin'), 'Spring','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/framewroks_spring.png?alt=media&token=dca9eef6-6960-4051-b6d9-9e241cc267c6'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'kotlin'), 'Ktor','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/framewroks_ktor.png?alt=media&token=20ce9153-16be-4517-88be-a7fe06f8d9ae');
-- PHPのフレームワーク
INSERT INTO frameworks (tech_tag_id, framework,icon)
VALUES ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'PHP'), 'Laravel','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/framewroks_laravel.png?alt=media&token=2ad52138-fe09-4193-97b8-23513b96198f'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'PHP'), 'Symfony','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/framewroks_symfony.png?alt=media&token=e7782cdd-ff93-4943-8470-e38b932155c6'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'PHP'), 'CodeIgniter','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/framewroks_codeIgniter.png?alt=media&token=75ead3dc-0c17-47d9-b4db-b3db92019c81');
-- Rustのフレームワーク
INSERT INTO frameworks (tech_tag_id, framework,icon)
VALUES ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Rust'), 'Rocket','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/framewroks_rocket.png?alt=media&token=ecba7706-4cc1-4bcb-8de0-02d4cac0433e'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Rust'), 'Actix-web','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/framewroks_rust.png?alt=media&token=f2809528-2515-4e93-835a-34b21e264aa1'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Rust'), 'Tide','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/framewroks_Tide.png?alt=media&token=3aec2937-01f8-4a93-8ab6-e1fdcf583dcc');
-- Rubyのフレームワーク
INSERT INTO frameworks (tech_tag_id, framework,icon)
VALUES ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Ruby'), 'Ruby on Rails','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/frameworks_rails.png?alt=media&token=c32da1d0-025b-44ea-b5d4-a329ff02322d'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Ruby'), 'Sinatra','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/frameworks_sinatra.png?alt=media&token=0cf13432-dee6-4a37-9d80-3390bcf143e4'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Ruby'), 'Hanami','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/frameworks_hanami.png?alt=media&token=efd8ec7a-4ea0-4a9a-a244-5e980f1c67da'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Ruby'), 'Padrino','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/frameworks_padrino.png?alt=media&token=f0b9e1e0-c56e-4471-8d63-e5fee587c910');
-- Rのフレームワーク
INSERT INTO frameworks (tech_tag_id, framework,icon)
VALUES ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'R'), 'Mojolicious','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/frameworks_mojolicious.png?alt=media&token=36480918-2bec-4bf3-a491-fc71b31763d6'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'R'), 'Dancer','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/frameworks_dancer.png?alt=media&token=7760c2ac-fbc6-4234-b64f-844934091794');
-- データベースのフレームワーク
INSERT INTO frameworks (tech_tag_id, framework,icon)
VALUES ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'DataBase'), 'MySQL','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/frameworks_mysql.png?alt=media&token=2b8c71f0-0516-4334-a0e9-b39a035dea4a'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'DataBase'), 'PostgreSQL','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/frameworks_postgres.png?alt=media&token=56d4a71d-b6d6-4112-b61f-3786b3114628'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'DataBase'), 'MongoDB','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/frameworks_mongodb.png?alt=media&token=034d2f46-6cdd-4144-8f2c-825104b76c29'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'DataBase'), 'Oracle','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/frameworks_oracle.png?alt=media&token=d24b7ae9-dd81-49a3-bd1d-ef97ce1e500f'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'DataBase'), 'Couchbase','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/frameworks_couchbase.png?alt=media&token=ea1c2136-132c-4d3c-9f57-9e71f871e489'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'DataBase'), 'SQLServer','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/frameworks_sqlserver.png?alt=media&token=28e94217-128e-48ca-a81e-94e4a2a9e9fc'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'DataBase'), 'Redis','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/frameworks_redis.png?alt=media&token=5b0b5ddf-9a17-43eb-8f6e-f9dc83efd1f4'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'DataBase'), 'AlibabaCloud','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/frameworks_alibaba.png?alt=media&token=633aff8d-f791-4bdc-bcd9-815c5ac4c971'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'DataBase'), 'OracleCloud','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/frameworks_oracle.png?alt=media&token=d24b7ae9-dd81-49a3-bd1d-ef97ce1e500f');
-- クラウド枠
INSERT INTO frameworks (tech_tag_id, framework,icon)
VALUES  ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Cloud'),'AWS','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/frameworks_aws.png?alt=media&token=185b4b40-9a0a-4c9c-98eb-8cf49550cbf4'),
        ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Cloud'),'Microsoft Azure','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/frameworks_azure.png?alt=media&token=019c4a72-b82b-4ed9-a54e-1a9a6a069ab0'),
        ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Cloud'),'GCP','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/frameworks_gcp.png?alt=media&token=c0bdada2-9279-4ca4-ba88-08cb8b7649db'),
        ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Cloud'),'IBM Cloud','https://firebasestorage.googleapis.com/v0/b/hackthon-geek-v6.appspot.com/o/frameworks_ibmcoud.png?alt=media&token=d91782db-aee0-43c8-82ac-ddafe819e7d9');


-- ステータス追加
INSERT INTO "status_tags" ("status") VALUES
('オンライン'),
('オフライン'),
('初心者歓迎'),
('急募');

INSERT INTO roles (role_id,role) VALUES 
(1,'フロント'), 
(2,'バック'), 
(3,'モバイル'), 
(4,'XR'), 
(5,'インフラ'),
(6,'電子工学'),
(7,'初心者'),
(8,'デザイン'),
(9,'モデラー'),
(10,'マネジメント'),
(11,'デザイン');