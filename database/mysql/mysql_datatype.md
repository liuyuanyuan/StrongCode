# MySQL数据类型





## JSON

```mysql
-- tb_stu_segment definition
CREATE TABLE `tb_stu_segment` (
  `user_id` bigint(20) DEFAULT NULL COMMENT '用户ID',
  `course_id` int(11) DEFAULT NULL COMMENT '课程id',
  `chapter_id` int(11) DEFAULT NULL COMMENT '章节id',
  `lesson_id` int(11) DEFAULT NULL COMMENT '课节id',
  `segement_id` int(11) DEFAULT NULL COMMENT '环节id',
  `interact_count` tinyint(4) DEFAULT NULL COMMENT '本环节-互动题总数（固定值）',
  `finish_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '本环节-本次学完时间',
  `star_count` tinyint(4) DEFAULT NULL COMMENT '本环节-本次获得星星总数',
  `answer` json DEFAULT NULL COMMENT '本环节-本次答题路径 Map<String, List<String>>',
  `express_count` tinyint(4) DEFAULT NULL COMMENT '本环节-语音跟读题-答题次数',
  `sid` bigint(20) NOT NULL AUTO_INCREMENT,
  `goods_plan_id` int(11) DEFAULT NULL COMMENT '商品期数id',
  PRIMARY KEY (`sid`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='学员-环节-学完记录';


-- insert test data
INSERT INTO tb_stu_segment (user_id,course_type,course_id,chapter_id,lesson_id,segement_id,interact_count,finish_time,star_count,answer,express_count,goods_plan_id,`tss.answer->'$[0]'`,`tss.answer->'$[1]'`) VALUES 
(1263814329964630018,NULL,95,172,912,2,3,'2020-07-06 15:25:39.0',6,'[{"id": 1, "star": 2, "type": 201, "count": 2, "answer": ["A", "B"]}, {"id": 2, "star": 2, "type": 203, "count": 2, "answer": ["stu-record/xxbgd.mp3"]}, {"id": 3, "star": 2, "type": 204, "count": 2, "answer": ["stu-record/yuan.png"]}]',2,20,'{"id": 1, "star": 2, "type": 201, "count": 2, "answer": ["A", "B"]}','{"id": 2, "star": 2, "type": 203, "count": 2, "answer": ["stu-record/xxbgd.mp3"]}')
;


-- select json array
 select user_id, segement_id , answer->'$[0].count' from tb_stu_segment where answer->'$[0].id'=1;
 
 
```

