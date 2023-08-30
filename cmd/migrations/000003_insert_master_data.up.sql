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
INSERT INTO frameworks (tech_tag_id, framework)
VALUES ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Python'), 'Django'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Python'), 'Flask'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Python'), 'FastAPI');
-- JavaScriptのフレームワーク
INSERT INTO frameworks (tech_tag_id, framework)
VALUES ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'JavaScript'), 'React.js'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'JavaScript'), 'Vue.js'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'JavaScript'), 'Three.js'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'JavaScript'), 'Next.js'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'JavaScript'), 'Node.js'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'JavaScript'), 'Angular');
-- Javaのフレームワーク
INSERT INTO frameworks (tech_tag_id, framework)
VALUES ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Java'), 'JavaPlayFramework'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Java'), 'Spring'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Java'), 'ApacheStruts'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Java'), 'JSF'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Java'), 'Wicket');
-- Goのフレームワーク
INSERT INTO frameworks (tech_tag_id, framework)
VALUES ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Go'), 'Gin'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Go'), 'Beego'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Go'), 'Revel'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Go'), 'Echo');
-- Cのフレームワーク
INSERT INTO frameworks (tech_tag_id, framework)
VALUES ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'C'), '.NET'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'C'), 'ASP.NET'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'C'), 'ASP.NET MVC');
-- C#のフレームワーク
INSERT INTO frameworks (tech_tag_id, framework)
VALUES ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Csharp'), '.NET'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Csharp'), 'ASP.NET'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Csharp'), 'ASP.NET MVC');
-- C++のフレームワーク
INSERT INTO frameworks (tech_tag_id, framework)
VALUES ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Cpp'), 'Qt');
-- Kotlinのフレームワーク
INSERT INTO frameworks (tech_tag_id, framework)
VALUES ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'kotlin'), 'Spring'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'kotlin'), 'Ktor');
-- PHPのフレームワーク
INSERT INTO frameworks (tech_tag_id, framework)
VALUES ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'PHP'), 'Laravel'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'PHP'), 'Symfony'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'PHP'), 'CodeIgniter');
-- Rustのフレームワーク
INSERT INTO frameworks (tech_tag_id, framework)
VALUES ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Rust'), 'Rocket'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Rust'), 'Actix-web'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Rust'), 'Tide');
-- Rubyのフレームワーク
INSERT INTO frameworks (tech_tag_id, framework)
VALUES ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Ruby'), 'Ruby on Rails'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Ruby'), 'Sinatra'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Ruby'), 'Hanami'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Ruby'), 'Padrino');
-- Rのフレームワーク
INSERT INTO frameworks (tech_tag_id, framework)
VALUES ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'R'), 'Mojolicious'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'R'), 'Dancer');
-- データベースのフレームワーク
INSERT INTO frameworks (tech_tag_id, framework)
VALUES ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'DataBase'), 'MySQL'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'DataBase'), 'PostgreSQL'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'DataBase'), 'MongoDB'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'DataBase'), 'Oracle'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'DataBase'), 'Couchbase'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'DataBase'), 'SQLServer'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'DataBase'), 'Redis'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'DataBase'), 'AlibabaCloud'),
       ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'DataBase'), 'OracleCloud');
-- クラウド枠
INSERT INTO frameworks (tech_tag_id, framework)
VALUES  ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Cloud'),'AWS'),
        ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Cloud'),'Microsoft Azure'),
        ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Cloud'),'GCP'),
        ((SELECT "tech_tag_id" FROM tech_tags WHERE "language" = 'Cloud'),'IBM Cloud');


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